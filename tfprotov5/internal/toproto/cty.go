package toproto

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/vmihailenco/msgpack"
)

type errWithPath struct {
	err  error
	path []string
}

func (e errWithPath) Error() string {
	return fmt.Errorf("error marshaling %s: %w", strings.Join(e.path, "."), e.err).Error()
}

type unknownType struct{}

var unknownVal = unknownType{}

func (u unknownType) MarshalMsgpack() ([]byte, error) {
	return []byte{0xd4, 0, 0}, nil
}

func Cty(in tftypes.RawValue) (tfplugin5.DynamicValue, error) {
	// always populate msgpack, as per
	// https://github.com/hashicorp/terraform/blob/doc-provider-value-wire-protocol/docs/plugin-protocol/object-wire-format.md
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	err := marshal(in.Value, in.Type, nil, enc)
	if err != nil {
		return tfplugin5.DynamicValue{}, err
	}
	return tfplugin5.DynamicValue{
		Msgpack: buf.Bytes(),
	}, nil
}

func marshal(val interface{}, typ tftypes.Type, path []string, enc *msgpack.Encoder) error {
	if typ.Is(tftypes.DynamicPseudoType) {
		dst, ok := val.(tftypes.RawValue)
		if !ok {
			return unexpectedValueTypeError(path, tftypes.RawValue{}, val, typ)
		}
		typeJSON, err := dst.Type.MarshalJSON()
		if err != nil {
			return pathError(path, "error generating JSON for type %s: %w", dst.Type, err)
		}
		err = enc.EncodeArrayLen(2)
		if err != nil {
			return pathError(path, "error encoding array length:  %w", err)
		}
		err = enc.EncodeBytes(typeJSON)
		if err != nil {
			return pathError(path, "error encoding JSON type info: %w", err)
		}
		err = marshal(dst.Value, dst.Type, path, enc)
		if err != nil {
			return pathError(path, "error marshaling DynamicPseudoType value: %w", err)
		}
		return nil

	}
	if val == tftypes.UnknownValue {
		err := enc.Encode(unknownVal)
		if err != nil {
			return pathError(path, "error encoding UnknownValue: %w", err)
		}
		return nil
	}
	if val == nil {
		err := enc.EncodeNil()
		if err != nil {
			return pathError(path, "error encoding null value: %w", err)
		}
		return nil
	}
	switch {
	case typ.Is(tftypes.String):
		s, ok := val.(string)
		if !ok {
			u, ok := val.([]uint8)
			if !ok {
				return unexpectedValueTypeError(path, s, val, typ)
			}
			b := make([]byte, 0, len(u))
			for _, i := range u {
				b = append(b, i)
			}
			s = string(b)
		}
		err := enc.EncodeString(s)
		if err != nil {
			return pathError(path, "error encoding string value: %w", err)
		}
		return nil
	case typ.Is(tftypes.Number):
		// TODO: can we accept other types besides *big.Float?
		// at the very least it would be nice to have built-in handling
		// for Go's int/float type variations
		n, ok := val.(*big.Float)
		if !ok {
			u, ok := val.(uint16)
			if !ok {
				return unexpectedValueTypeError(path, n, val, typ)
			}
			err := enc.EncodeUint16(u)
			if err != nil {
				return pathError(path, "error encoding uint16 value: %w", err)
			}
			return nil
		}
		if iv, acc := n.Int64(); acc == big.Exact {
			err := enc.EncodeInt(iv)
			if err != nil {
				return pathError(path, "error encoding int value: %w", err)
			}
		} else if fv, acc := n.Float64(); acc == big.Exact {
			err := enc.EncodeFloat64(fv)
			if err != nil {
				return pathError(path, "error encoding float value: %w", err)
			}
		} else {
			err := enc.EncodeString(n.Text('f', -1))
			if err != nil {
				return pathError(path, "error encoding number string value: %w", err)
			}
		}
		return nil
	case typ.Is(tftypes.Bool):
		b, ok := val.(bool)
		if !ok {
			return unexpectedValueTypeError(path, b, val, typ)
		}
		err := enc.EncodeBool(b)
		if err != nil {
			return pathError(path, "error encoding bool value: %w", err)
		}
		return nil
	case typ.Is(tftypes.List{}):
		// TODO: can we make it so you don't need to pass in an []interface{}
		// maybe we can at least special-case the builtin primitives and tftypes.RawValue?
		l, ok := val.([]interface{})
		if !ok {
			return unexpectedValueTypeError(path, l, val, typ)
		}
		err := enc.EncodeArrayLen(len(l))
		if err != nil {
			return pathError(path, "error encoding list length: %w", err)
		}
		for pos, i := range l {
			elemPath := newPath(path, pos)
			err := marshal(i, typ.(tftypes.List).ElementType, elemPath, enc)
			if err != nil {
				return pathError(path, "error encoding list element: %w", err)
			}
		}
		return nil
	case typ.Is(tftypes.Set{}):
		// TODO: can we make it so you don't need to pass in an []interface{}
		// maybe we can at least special-case the builtin primitives and tftypes.RawValue?
		s, ok := val.([]interface{})
		if !ok {
			return unexpectedValueTypeError(path, s, val, typ)
		}
		err := enc.EncodeArrayLen(len(s))
		if err != nil {
			return pathError(path, "error encoding set length: %w", err)
		}
		for _, i := range s {
			elemPath := newPath(path, i)
			err := marshal(i, typ.(tftypes.Set).ElementType, elemPath, enc)
			if err != nil {
				return pathError(path, "error encoding set element: %w", err)
			}
		}
		return nil
	case typ.Is(tftypes.Map{}):
		// TODO: can we make it so you don't need to pass in a map[string]interface{}
		// maybe we can at least special-case the built-in primitives and tftypes.RawValue?
		m, ok := val.(map[string]interface{})
		if !ok {
			return unexpectedValueTypeError(path, m, val, typ)
		}
		err := enc.EncodeMapLen(len(m))
		if err != nil {
			return pathError(path, "error encoding map length: %w", err)
		}
		for k, v := range m {
			attrPath := newPath(path, k)
			err := marshal(k, tftypes.String, attrPath, enc)
			if err != nil {
				return pathError(path, "error encoding map key: %w", err)
			}
			err = marshal(v, typ.(tftypes.Map).AttributeType, attrPath, enc)
			if err != nil {
				return pathError(path, "error encoding map value: %v", err)
			}
		}
		return nil
	case typ.Is(tftypes.Tuple{}):
		// TODO: can we make it so you don't need to pass in an []interface{}
		// maybe we can at least special-case tftypes.RawValue?
		t, ok := val.([]interface{})
		if !ok {
			return unexpectedValueTypeError(path, t, val, typ)
		}
		types := typ.(tftypes.Tuple).ElementTypes
		err := enc.EncodeArrayLen(len(t))
		if err != nil {
			return pathError(path, "error encoding tuple length: %w", err)
		}
		for pos, v := range t {
			ty := types[pos]
			elementPath := newPath(path, pos)
			err := marshal(v, ty, elementPath, enc)
			if err != nil {
				return pathError(path, "error encoding tuple element: %w", err)
			}
		}
		return nil
	case typ.Is(tftypes.Object{}):
		// TODO: can we make it so you don't need to pass in a map[string]interface{}
		// maybe we can at least special-case tftypes.RawValue?
		o, ok := val.(map[string]interface{})
		if !ok {
			return unexpectedValueTypeError(path, o, val, typ)
		}
		types := typ.(tftypes.Object).AttributeTypes
		err := enc.EncodeMapLen(len(o))
		if err != nil {
			return pathError(path, "error encoding object length: %w", err)
		}
		for k, v := range o {
			ty := types[k]
			attrPath := newPath(path, k)
			err := marshal(k, tftypes.String, attrPath, enc)
			if err != nil {
				return pathError(path, "error encoding object key: %w", err)
			}
			err = marshal(v, ty, attrPath, enc)
			if err != nil {
				return pathError(path, "error encoding object value: %w", err)
			}
		}
		return nil
	}
	return fmt.Errorf("unknown type %s", typ)
}

func pathError(path []string, f string, args ...interface{}) error {
	return errWithPath{
		path: path,
		err:  fmt.Errorf(f, args...),
	}
}

func unexpectedValueTypeError(path []string, expected, got interface{}, typ tftypes.Type) error {
	return pathError(path, "unexpected value type %T, %s values must be of type %T", got, typ, expected)
}

func newPath(path []string, v interface{}) []string {
	n := make([]string, len(path)+1)
	for i, v := range path {
		n[i] = v
	}
	n[len(n)-1] = fmt.Sprintf("%v", v)
	return n
}

func CtyType(in tftypes.Type) ([]byte, error) {
	switch {
	case in.Is(tftypes.String), in.Is(tftypes.Bool), in.Is(tftypes.Number),
		in.Is(tftypes.List{}), in.Is(tftypes.Map{}),
		in.Is(tftypes.Set{}), in.Is(tftypes.Object{}),
		in.Is(tftypes.Tuple{}), in.Is(tftypes.DynamicPseudoType):
		return in.MarshalJSON()
	}
	return nil, fmt.Errorf("unknown type %s", in)
}
