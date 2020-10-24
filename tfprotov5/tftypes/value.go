package tftypes

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/vmihailenco/msgpack"
)

type Unmarshaler interface {
	UnmarshalTerraform5Type(Value) error
}

type ErrUnhandledType string

func (e ErrUnhandledType) Error() string {
	return fmt.Sprintf("unhandled Terraform type %s", string(e))
}

type msgPackUnknownType struct{}

var msgPackUnknownVal = msgPackUnknownType{}

func (u msgPackUnknownType) MarshalMsgpack() ([]byte, error) {
	return []byte{0xd4, 0, 0}, nil
}

// Value represents a form of a Terraform type that can be parsed into a Go
// type.
type Value struct {
	typ   Type
	value interface{}
}

func NewValue(t Type, val interface{}) Value {
	return Value{
		typ:   t,
		value: val,
	}
}

func (val Value) As(dst interface{}) error {
	unmarshaler, ok := dst.(Unmarshaler)
	if ok {
		return unmarshaler.UnmarshalTerraform5Type(val)
	}
	if !val.IsKnown() {
		return fmt.Errorf("unmarshaling unknown values is not supported")
	}
	switch target := dst.(type) {
	case *string:
		if val.IsNull() {
			*target = ""
			return nil
		}
		v, ok := val.value.(string)
		if !ok {
			return fmt.Errorf("can't unmarshal %s into %T, expected string", val.typ, dst)
		}
		*target = v
		return nil
	case **string:
		if val.IsNull() {
			*target = nil
			return nil
		}
		return val.As(*target)
	case *big.Float:
		if val.IsNull() {
			target.Set(big.NewFloat(0))
			return nil
		}
		v, ok := val.value.(*big.Float)
		if !ok {
			return fmt.Errorf("can't unmarshal %s into %T, expected *big.Float", val.typ, dst)
		}
		target.Set(v)
		return nil
	case **big.Float:
		if val.IsNull() {
			*target = nil
			return nil
		}
		return val.As(*target)
	case *bool:
		if val.IsNull() {
			*target = false
			return nil
		}
		v, ok := val.value.(bool)
		if !ok {
			return fmt.Errorf("can't unmarshal %s into %T, expected boolean", val.typ, dst)
		}
		*target = v
		return nil
	case **bool:
		if val.IsNull() {
			*target = nil
			return nil
		}
		return val.As(*target)
	case *map[string]Value:
		if val.IsNull() {
			*target = map[string]Value{}
			return nil
		}
		v, ok := val.value.(map[string]Value)
		if !ok {
			return fmt.Errorf("can't unmarshal %s into %T, expected map[string]tftypes.Value", val.typ, dst)
		}
		*target = v
		return nil
	case **map[string]Value:
		if val.IsNull() {
			*target = nil
			return nil
		}
		return val.As(*target)
	case *[]Value:
		if val.IsNull() {
			*target = []Value{}
			return nil
		}
		v, ok := val.value.([]Value)
		if !ok {
			return fmt.Errorf("can't unmarshal %s into %T expected []tftypes.Value", val.typ, dst)
		}
		*target = v
		return nil
	case **[]Value:
		if val.IsNull() {
			*target = nil
			return nil
		}
		return val.As(*target)
	}
	return fmt.Errorf("can't unmarshal into %T, needs UnmarshalTerraform5Type method", dst)
}

func (val Value) Is(t Type) bool {
	if val.typ == nil || t == nil {
		return val.typ == nil && t == nil
	}
	return val.typ.Is(t)
}

func (val Value) IsKnown() bool {
	return val.value != UnknownValue
}

func (val Value) IsFullyKnown() bool {
	switch val.typ.(type) {
	case primitive:
		return val.IsKnown()
	case List, Set, Tuple:
		for _, v := range val.value.([]Value) {
			if !v.IsFullyKnown() {
				return false
			}
		}
		return true
	case Map, Object:
		for _, v := range val.value.(map[string]Value) {
			if !v.IsFullyKnown() {
				return false
			}
		}
		return true
	}
	panic(fmt.Sprintf("unknown type %T", val.typ))
}

func (val Value) IsNull() bool {
	return val.value == nil
}

func (val Value) MarshalMsgPack(t Type) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)

	err := marshalMsgPack(val, t, AttributePath{}, enc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unexpectedValueTypeError(p AttributePath, expected, got interface{}, typ Type) error {
	return p.NewErrorf("unexpected value type %T, %s values must be of type %T", got, typ, expected)
}
