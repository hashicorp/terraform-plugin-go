package tftypes

import (
	"fmt"
	"strconv"
)

const (
	String = Type('S')
)

type Unmarshaler interface {
	UnmarshalTerraform5Type(t Type, v interface{}) error
}

type ErrUnhandledType Type

func (e ErrUnhandledType) Error() string {
	return fmt.Sprintf("unhandled Terraform type %s", Type(e).String())
}

// RawValue represents a form of a Terraform type that can be parsed into a Go
// type.
type RawValue struct {
	Type  Type
	Value interface{}
}

func (r RawValue) Unmarshal(dst interface{}) error {
	if unmarshaler, ok := dst.(Unmarshaler); ok {
		return unmarshaler.UnmarshalTerraform5Type(r.Type, r.Value)
	}
	switch dst.(type) {
	// TODO: cast to primitives as a default behavior
	}
	return ErrUnhandledType(r.Type)
}

// Type is used to identify Terraform types.
type Type rune

func (t Type) String() string {
	switch t {
	// TODO: return string version of type constants
	case String:
		return "tftypes.String"
	}
	return "unknown type " + strconv.QuoteRune(rune(t))
}
