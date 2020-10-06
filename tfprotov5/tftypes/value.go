package tftypes

import (
	"bytes"
	"fmt"

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

func (v Value) As(dst interface{}) error {
	unmarshaler, ok := dst.(Unmarshaler)
	if !ok {
		// TODO: let's do unmarshaling to builtin types out of the box
		return fmt.Errorf("can't unmarshal into %T, needs UnmarshalTerraform5Type method", dst)
	}
	return unmarshaler.UnmarshalTerraform5Type(v)
}

func (v Value) Is(t Type) bool {
	if v.typ == nil || t == nil {
		return v.typ == nil && t == nil
	}
	return v.typ.Is(t)
}

func (v Value) IsKnown() bool {
	return v.value != UnknownValue
}

func (v Value) IsFullyKnown() bool {
	switch v.typ.(type) {
	case primitive:
		return v.IsKnown()
	case List, Set, Tuple:
		for _, val := range v.value.([]Value) {
			if !val.IsFullyKnown() {
				return false
			}
		}
		return true
	case Map, Object:
		for _, val := range v.value.(map[string]Value) {
			if !val.IsFullyKnown() {
				return false
			}
		}
		return true
	}
	panic(fmt.Sprintf("unknown type %T", v.typ))
}

func (v Value) IsNull() bool {
	return v.value == nil
}

func (v Value) MarshalMsgPack(t Type) ([]byte, error) {
	// always populate msgpack, as per
	// https://github.com/hashicorp/terraform/blob/doc-provider-value-wire-protocol/docs/plugin-protocol/object-wire-format.md
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)

	err := marshalMsgPack(v, t, Path{}, enc)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func unexpectedValueTypeError(p Path, expected, got interface{}, typ Type) error {
	return p.NewErrorf("unexpected value type %T, %s values must be of type %T", got, typ, expected)
}
