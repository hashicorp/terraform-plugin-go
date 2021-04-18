package tftypes

import (
	"bytes"
	"math/big"

	"github.com/vmihailenco/msgpack"
	msgpackCodes "github.com/vmihailenco/msgpack/codes"
)

func ValueFromMsgPack(data []byte, typ Type) (Value, error) {
	r := bytes.NewReader(data)
	dec := msgpack.NewDecoder(r)
	return msgpackUnmarshal(dec, typ, AttributePath{})
}

func msgpackUnmarshal(dec *msgpack.Decoder, typ Type, path AttributePath) (Value, error) {
	peek, err := dec.PeekCode()
	if err != nil {
		return Value{}, path.NewErrorf("error peeking next byte: %w", err)
	}
	if msgpackCodes.IsExt(peek) {
		// as with go-cty, assume all extensions are unknown values
		err := dec.Skip()
		if err != nil {
			return Value{}, path.NewErrorf("error skipping extension byte: %w", err)
		}
		return NewValue(typ, UnknownValue), nil
	}
	if typ.Is(DynamicPseudoType) {
		return msgpackUnmarshalDynamic(dec, path)
	}
	if peek == msgpackCodes.Nil {
		err := dec.Skip()
		if err != nil {
			return Value{}, path.NewErrorf("error skipping nil byte: %w", err)
		}
		return NewValue(typ, nil), nil
	}

	switch {
	case typ.Is(String):
		rv, err := dec.DecodeString()
		if err != nil {
			return Value{}, path.NewErrorf("error decoding string: %w", err)
		}
		return NewValue(String, rv), nil
	case typ.Is(Number):
		peek, err := dec.PeekCode()
		if err != nil {
			return Value{}, path.NewErrorf("couldn't peek number: %w", err)
		}
		if msgpackCodes.IsFixedNum(peek) {
			rv, err := dec.DecodeInt64()
			if err != nil {
				return Value{}, path.NewErrorf("couldn't decode number as int64: %w", err)
			}
			return NewValue(Number, big.NewFloat(float64(rv))), nil
		}
		switch peek {
		case msgpackCodes.Int8, msgpackCodes.Int16, msgpackCodes.Int32, msgpackCodes.Int64:
			rv, err := dec.DecodeInt64()
			if err != nil {
				return Value{}, path.NewErrorf("couldn't decode number as int64: %w", err)
			}
			return NewValue(Number, big.NewFloat(float64(rv))), nil
		case msgpackCodes.Uint8, msgpackCodes.Uint16, msgpackCodes.Uint32, msgpackCodes.Uint64:
			rv, err := dec.DecodeUint64()
			if err != nil {
				return Value{}, path.NewErrorf("couldn't decode number as uint64: %w", err)
			}
			return NewValue(Number, big.NewFloat(float64(rv))), nil
		case msgpackCodes.Float, msgpackCodes.Double:
			rv, err := dec.DecodeFloat64()
			if err != nil {
				return Value{}, path.NewErrorf("couldn't decode number as float64: %w", err)
			}
			return NewValue(Number, big.NewFloat(float64(rv))), nil
		default:
			rv, err := dec.DecodeString()
			if err != nil {
				return Value{}, path.NewErrorf("couldn't decode number as string: %w", err)
			}
			// according to
			// https://github.com/hashicorp/go-cty/blob/85980079f637862fa8e43ddc82dd74315e2f4c85/cty/value_init.go#L49
			// Base 10, precision 512, and rounding to nearest even
			// is the standard way to handle numbers arriving as
			// strings.
			fv, _, err := big.ParseFloat(rv, 10, 512, big.ToNearestEven)
			if err != nil {
				return Value{}, path.NewErrorf("error parsing %q as number: %w", rv, err)
			}
			return NewValue(Number, fv), nil
		}
	case typ.Is(Bool):
		rv, err := dec.DecodeBool()
		if err != nil {
			return Value{}, path.NewErrorf("couldn't decode bool: %w", err)
		}
		return NewValue(Bool, rv), nil
	case typ.Is(List{}):
		return msgpackUnmarshalList(dec, typ.(List).ElementType, path)
	case typ.Is(Set{}):
		return msgpackUnmarshalSet(dec, typ.(Set).ElementType, path)
	case typ.Is(Map{}):
		return msgpackUnmarshalMap(dec, typ.(Map).AttributeType, path)
	case typ.Is(Tuple{}):
		return msgpackUnmarshalTuple(dec, typ.(Tuple).ElementTypes, path)
	case typ.Is(Object{}):
		return msgpackUnmarshalObject(dec, typ.(Object).AttributeTypes, path)
	}
	return Value{}, path.NewErrorf("unsupported type %s", typ.String())
}

func msgpackUnmarshalList(dec *msgpack.Decoder, typ Type, path AttributePath) (Value, error) {
	length, err := dec.DecodeArrayLen()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding list length: %w", err)
	}

	switch {
	case length < 0:
		return NewValue(List{
			ElementType: typ,
		}, nil), nil
	case length == 0:
		return NewValue(List{
			ElementType: typ,
		}, []Value{}), nil
	}

	vals := make([]Value, 0, length)
	for i := 0; i < length; i++ {
		path.WithElementKeyInt(int64(i))
		val, err := msgpackUnmarshal(dec, typ, path)
		if err != nil {
			return Value{}, err
		}
		vals = append(vals, val)
		path.WithoutLastStep()
	}

	elTyp := typ
	if elTyp.Is(DynamicPseudoType) {
		elTyp, err = TypeFromElements(vals)
		if err != nil {
			return Value{}, err
		}
	}

	return NewValue(List{
		ElementType: elTyp,
	}, vals), nil
}

func msgpackUnmarshalSet(dec *msgpack.Decoder, typ Type, path AttributePath) (Value, error) {
	length, err := dec.DecodeArrayLen()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding set length: %w", err)
	}

	switch {
	case length < 0:
		return NewValue(Set{
			ElementType: typ,
		}, nil), nil
	case length == 0:
		return NewValue(Set{
			ElementType: typ,
		}, []Value{}), nil
	}

	vals := make([]Value, 0, length)
	for i := 0; i < length; i++ {
		path.WithElementKeyInt(int64(i))
		val, err := msgpackUnmarshal(dec, typ, path)
		if err != nil {
			return Value{}, err
		}
		vals = append(vals, val)
		path.WithoutLastStep()
	}

	elTyp, err := TypeFromElements(vals)
	if err != nil {
		return Value{}, err
	}

	return NewValue(Set{
		ElementType: elTyp,
	}, vals), nil
}

func msgpackUnmarshalMap(dec *msgpack.Decoder, typ Type, path AttributePath) (Value, error) {
	length, err := dec.DecodeMapLen()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding map length: %w", err)
	}

	switch {
	case length < 0:
		return NewValue(Map{
			AttributeType: typ,
		}, nil), nil
	case length == 0:
		return NewValue(Map{
			AttributeType: typ,
		}, map[string]Value{}), nil
	}

	vals := make(map[string]Value, length)
	for i := 0; i < length; i++ {
		key, err := dec.DecodeString()
		if err != nil {
			return Value{}, path.NewErrorf("error decoding map key: %w", err)
		}
		path.WithElementKeyString(key)
		val, err := msgpackUnmarshal(dec, typ, path)
		if err != nil {
			return Value{}, err
		}
		vals[key] = val
		path.WithoutLastStep()
	}
	return NewValue(Map{
		AttributeType: typ,
	}, vals), nil
}

func msgpackUnmarshalTuple(dec *msgpack.Decoder, types []Type, path AttributePath) (Value, error) {
	length, err := dec.DecodeArrayLen()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding tuple length: %w", err)
	}

	switch {
	case length < 0:
		return NewValue(Tuple{
			ElementTypes: types,
		}, nil), nil
	case length == 0:
		return NewValue(Tuple{
			// no elements means no types
			ElementTypes: nil,
		}, []Value{}), nil
	case length != len(types):
		return Value{}, path.NewErrorf("error decoding tuple; expected %d items, got %d", len(types), length)
	}

	vals := make([]Value, 0, length)
	for i := 0; i < length; i++ {
		path.WithElementKeyInt(int64(i))
		typ := types[i]
		val, err := msgpackUnmarshal(dec, typ, path)
		if err != nil {
			return Value{}, err
		}
		vals = append(vals, val)
		path.WithoutLastStep()
	}

	return NewValue(Tuple{
		ElementTypes: types,
	}, vals), nil
}

func msgpackUnmarshalObject(dec *msgpack.Decoder, types map[string]Type, path AttributePath) (Value, error) {
	length, err := dec.DecodeMapLen()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding object length: %w", err)
	}

	switch {
	case length < 0:
		return NewValue(Object{
			AttributeTypes: types,
		}, nil), nil
	case length == 0:
		return NewValue(Object{
			// no attributes means no types
			AttributeTypes: map[string]Type{},
		}, map[string]Value{}), nil
	case length != len(types):
		return Value{}, path.NewErrorf("error decoding object; expected %d attributes, got %d", len(types), length)
	}

	vals := make(map[string]Value, length)
	for i := 0; i < length; i++ {
		key, err := dec.DecodeString()
		if err != nil {
			return Value{}, path.NewErrorf("error decoding object key: %w", err)
		}
		typ, exists := types[key]
		if !exists {
			return Value{}, path.NewErrorf("unknown attribute %q", key)
		}
		path.WithAttributeName(key)
		val, err := msgpackUnmarshal(dec, typ, path)
		if err != nil {
			return Value{}, err
		}
		vals[key] = val
		path.WithoutLastStep()
	}

	return NewValue(Object{
		AttributeTypes: types,
	}, vals), nil
}

func msgpackUnmarshalDynamic(dec *msgpack.Decoder, path AttributePath) (Value, error) {
	length, err := dec.DecodeArrayLen()
	if err != nil {
		return Value{}, path.NewErrorf("error checking length of DynamicPseudoType value: %w", err)
	}

	switch {
	case length == -1:
		return NewValue(DynamicPseudoType, nil), nil
	case length != 2:
		return Value{}, path.NewErrorf("expected %d elements in DynamicPseudoType array, got %d", 2, length)
	}

	typeJSON, err := dec.DecodeBytes()
	if err != nil {
		return Value{}, path.NewErrorf("error decoding bytes: %w", err)
	}
	typ, err := ParseJSONType(typeJSON) //nolint:staticcheck
	if err != nil {
		return Value{}, path.NewErrorf("error parsing type information: %w", err)
	}
	return msgpackUnmarshal(dec, typ, path)
}
