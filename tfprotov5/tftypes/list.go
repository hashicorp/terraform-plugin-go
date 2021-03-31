package tftypes

import (
	"fmt"
)

// List is a Terraform type representing an ordered collection of elements, all
// of the same type.
type List struct {
	ElementType Type

	// used to make this type uncomparable
	// see https://golang.org/ref/spec#Comparison_operators
	// this enforces the use of Is, instead
	_ []struct{}
}

// Is returns whether `t` is a List type or not. If `t` is an instance of the
// List type and its ElementType property is nil, it will return true. If `t`'s
// ElementType property is not nil, it will only return true if its ElementType
// is considered the same type as `l`'s ElementType.
func (l List) Is(t Type) bool {
	v, ok := t.(List)
	if !ok {
		return false
	}
	if v.ElementType != nil {
		return l.ElementType.Is(v.ElementType)
	}
	return ok
}

func (l List) String() string {
	return "tftypes.List[" + l.ElementType.String() + "]"
}

func (l List) private() {}

func (l List) supportedGoTypes() []string {
	return []string{"[]tftypes.Value"}
}

func valueFromList(typ Type, in interface{}) (Value, error) {
	switch value := in.(type) {
	case []Value:
		for pos, v := range value {
			if !v.Type().Is(typ) && !typ.Is(DynamicPseudoType) {
				return Value{}, fmt.Errorf("tftypes.NewValue can't use type %s as a value in position %d of %s. Expected type is %s.", v.Type(), pos, List{ElementType: typ}, typ)
			}
		}
		return Value{
			typ:   List{ElementType: typ},
			value: value,
		}, nil
	default:
		return Value{}, fmt.Errorf("tftypes.NewValue can't use %T as a tftypes.List. Expected types are: %s", in, formattedSupportedGoTypes(List{}))
	}
}

// MarshalJSON returns a JSON representation of the full type signature of `l`,
// including its ElementType.
//
// Deprecated: this is not meant to be called by third-party code.
func (l List) MarshalJSON() ([]byte, error) {
	elementType, err := l.ElementType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.List's element type %T to JSON: %w", l.ElementType, err)
	}
	return []byte(`["list",` + string(elementType) + `]`), nil
}
