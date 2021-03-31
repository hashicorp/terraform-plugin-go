package tftypes

import (
	"fmt"
)

// Set is a Terraform type representing an unordered collection of unique
// elements, all of the same type.
type Set struct {
	ElementType Type

	// used to make this type uncomparable
	// see https://golang.org/ref/spec#Comparison_operators
	// this enforces the use of Is, instead
	_ []struct{}
}

// Is returns whether `t` is a Set type or not. If `t` is an instance of the
// Set type and its ElementType property is nil, it will return true. If `t`'s
// ElementType property is not nil, it will only return true if its ElementType
// is considered the same type as `s`'s ElementType.
func (s Set) Is(t Type) bool {
	v, ok := t.(Set)
	if !ok {
		return false
	}
	if v.ElementType != nil {
		return s.ElementType.Is(v.ElementType)
	}
	return ok
}

func (s Set) String() string {
	return "tftypes.Set[" + s.ElementType.String() + "]"
}

func (s Set) private() {}

func (s Set) supportedGoTypes() []string {
	return []string{"[]tftypes.Value"}
}

func valueFromSet(typ Type, in interface{}) (Value, error) {
	switch value := in.(type) {
	case []Value:
		for pos, v := range value {
			if !v.Type().Is(typ) && !typ.Is(DynamicPseudoType) {
				return Value{}, fmt.Errorf("tftypes.NewValue can't use type %s as a value in position %d of %s. Expected type is %s.", v.Type(), pos, Set{ElementType: typ}, typ)
			}
		}
		return Value{
			typ:   Set{ElementType: typ},
			value: value,
		}, nil
	default:
		return Value{}, fmt.Errorf("tftypes.NewValue can't use %T as a tftypes.Set. Expected types are: %s", in, formattedSupportedGoTypes(Set{}))
	}
}

// MarshalJSON returns a JSON representation of the full type signature of `s`,
// including its ElementType.
//
// Deprecated: this is not meant to be called by third-party code.
func (s Set) MarshalJSON() ([]byte, error) {
	elementType, err := s.ElementType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.Set's element type %T to JSON: %w", s.ElementType, err)
	}
	return []byte(`["set",` + string(elementType) + `]`), nil
}
