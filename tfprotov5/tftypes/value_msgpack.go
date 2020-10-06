package tftypes

import (
	"fmt"
	"math/big"
	"sort"

	"github.com/vmihailenco/msgpack"
)

func marshalMsgPack(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	if typ.Is(DynamicPseudoType) && !val.Is(DynamicPseudoType) {
		return marshalMsgPackDynamicPseudoType(val, typ, p, enc)

	}
	if !val.IsKnown() {
		err := enc.Encode(msgPackUnknownVal)
		if err != nil {
			return p.NewErrorf("error encoding UnknownValue: %w", err)
		}
		return nil
	}
	if val.IsNull() {
		err := enc.EncodeNil()
		if err != nil {
			return p.NewErrorf("error encoding null value: %w", err)
		}
		return nil
	}
	switch {
	case typ.Is(String):
		return marshalMsgPackString(val, typ, p, enc)
	case typ.Is(Number):
		return marshalMsgPackNumber(val, typ, p, enc)
	case typ.Is(Bool):
		return marshalMsgPackBool(val, typ, p, enc)
	case typ.Is(List{}):
		return marshalMsgPackList(val, typ, p, enc)
	case typ.Is(Set{}):
		return marshalMsgPackSet(val, typ, p, enc)
	case typ.Is(Map{}):
		return marshalMsgPackMap(val, typ, p, enc)
	case typ.Is(Tuple{}):
		return marshalMsgPackTuple(val, typ, p, enc)
	case typ.Is(Object{}):
		return marshalMsgPackObject(val, typ, p, enc)
	}
	return fmt.Errorf("unknown type %s", typ)
}

func marshalMsgPackDynamicPseudoType(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	dst, ok := val.value.(Value)
	if !ok {
		return unexpectedValueTypeError(p, Value{}, val, typ)
	}
	typeJSON, err := dst.typ.MarshalJSON()
	if err != nil {
		return p.NewErrorf("error generating JSON for type %s: %w", dst.typ, err)
	}
	err = enc.EncodeArrayLen(2)
	if err != nil {
		return p.NewErrorf("error encoding array length:  %w", err)
	}
	err = enc.EncodeBytes(typeJSON)
	if err != nil {
		return p.NewErrorf("error encoding JSON type info: %w", err)
	}
	err = marshalMsgPack(dst, dst.typ, p, enc)
	if err != nil {
		return p.NewErrorf("error marshaling DynamicPseudoType value: %w", err)
	}
	return nil
}

func marshalMsgPackString(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	s, ok := val.value.(string)
	if !ok {
		return unexpectedValueTypeError(p, s, val, typ)
	}
	err := enc.EncodeString(s)
	if err != nil {
		return p.NewErrorf("error encoding string value: %w", err)
	}
	return nil
}

func marshalMsgPackNumber(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	// TODO: can we accept other types besides *big.Float?
	// at the very least it would be nice to have built-in handling
	// for Go's int/float type variations
	//
	// once we build up the Value.As interface, we may be able to take
	// advantage of that...
	n, ok := val.value.(*big.Float)
	if !ok {
		return unexpectedValueTypeError(p, n, val, typ)
	}
	if iv, acc := n.Int64(); acc == big.Exact {
		err := enc.EncodeInt(iv)
		if err != nil {
			return p.NewErrorf("error encoding int value: %w", err)
		}
	} else if fv, acc := n.Float64(); acc == big.Exact {
		err := enc.EncodeFloat64(fv)
		if err != nil {
			return p.NewErrorf("error encoding float value: %w", err)
		}
	} else {
		err := enc.EncodeString(n.Text('f', -1))
		if err != nil {
			return p.NewErrorf("error encoding number string value: %w", err)
		}
	}
	return nil
}

func marshalMsgPackBool(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	b, ok := val.value.(bool)
	if !ok {
		return unexpectedValueTypeError(p, b, val, typ)
	}
	err := enc.EncodeBool(b)
	if err != nil {
		return p.NewErrorf("error encoding bool value: %w", err)
	}
	return nil
}

func marshalMsgPackList(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	l, ok := val.value.([]Value)
	if !ok {
		return unexpectedValueTypeError(p, l, val, typ)
	}
	err := enc.EncodeArrayLen(len(l))
	if err != nil {
		return p.NewErrorf("error encoding list length: %w", err)
	}
	p = append(p, nil)
	for pos, i := range l {
		p[len(p)-1] = pos
		err := marshalMsgPack(i, typ.(List).ElementType, p, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

func marshalMsgPackSet(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	s, ok := val.value.([]Value)
	if !ok {
		return unexpectedValueTypeError(p, s, val, typ)
	}
	err := enc.EncodeArrayLen(len(s))
	if err != nil {
		return p.NewErrorf("error encoding set length: %w", err)
	}
	p = append(p, nil)
	for _, i := range s {
		p[len(p)-1] = i
		err := marshalMsgPack(i, typ.(Set).ElementType, p, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

func marshalMsgPackMap(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	m, ok := val.value.(map[string]Value)
	if !ok {
		return unexpectedValueTypeError(p, m, val, typ)
	}
	err := enc.EncodeMapLen(len(m))
	if err != nil {
		return p.NewErrorf("error encoding map length: %w", err)
	}
	p = append(p, nil)
	for k, v := range m {
		p[len(p)-1] = k
		err := marshalMsgPack(NewValue(String, k), String, p, enc)
		if err != nil {
			return p.NewErrorf("error encoding map key: %w", err)
		}
		err = marshalMsgPack(v, typ.(Map).AttributeType, p, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

func marshalMsgPackTuple(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	t, ok := val.value.([]Value)
	if !ok {
		return unexpectedValueTypeError(p, t, val, typ)
	}
	types := typ.(Tuple).ElementTypes
	err := enc.EncodeArrayLen(len(types))
	if err != nil {
		return p.NewErrorf("error encoding tuple length: %w", err)
	}
	p = append(p, nil)
	for pos, v := range t {
		ty := types[pos]
		p[len(p)-1] = pos
		err := marshalMsgPack(v, ty, p, enc)
		if err != nil {
			return err
		}
	}
	return nil
}

func marshalMsgPackObject(val Value, typ Type, p Path, enc *msgpack.Encoder) error {
	o, ok := val.value.(map[string]Value)
	if !ok {
		return unexpectedValueTypeError(p, o, val, typ)
	}
	types := typ.(Object).AttributeTypes
	keys := make([]string, 0, len(types))
	for k := range types {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	err := enc.EncodeMapLen(len(keys))
	if err != nil {
		return p.NewErrorf("error encoding object length: %w", err)
	}
	p = append(p, nil)
	for _, k := range keys {
		ty := types[k]
		p[len(p)-1] = k
		v, ok := o[k]
		if !ok {
			return p.NewErrorf("no value set")
		}
		err := marshalMsgPack(NewValue(String, k), String, p, enc)
		if err != nil {
			return err
		}
		err = marshalMsgPack(v, ty, p, enc)
		if err != nil {
			return err
		}
	}
	return nil
}
