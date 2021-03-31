package tftypes

import "fmt"

// Map is a Terraform type representing an unordered collection of elements,
// all of the same type, each identifiable with a unique string key.
type Map struct {
	AttributeType Type

	// used to make this type uncomparable
	// see https://golang.org/ref/spec#Comparison_operators
	// this enforces the use of Is, instead
	_ []struct{}
}

// Is returns whether `t` is a Map type or not. If `t` is an instance of the
// Map type and its AttributeType property is not nil, it will only return true
// if its AttributeType is considered the same type as `m`'s AttributeType.
func (m Map) Is(t Type) bool {
	v, ok := t.(Map)
	if !ok {
		return false
	}
	if v.AttributeType != nil {
		return m.AttributeType.Is(v.AttributeType)
	}
	return ok
}

func (m Map) String() string {
	return "tftypes.Map[" + m.AttributeType.String() + "]"
}

func (m Map) private() {}

func (m Map) supportedGoTypes() []string {
	return []string{"map[string]tftypes.Value"}
}

// MarshalJSON returns a JSON representation of the full type signature of `m`,
// including its AttributeType.
//
// Deprecated: this is not meant to be called by third-party code.
func (m Map) MarshalJSON() ([]byte, error) {
	attributeType, err := m.AttributeType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.Map's attribute type %T to JSON: %w", m.AttributeType, err)
	}
	return []byte(`["map",` + string(attributeType) + `]`), nil
}

func valueFromMap(typ Type, in interface{}) (Value, error) {
	switch value := in.(type) {
	case map[string]Value:
		for k, v := range value {
			if !v.Type().Is(typ) && !typ.Is(DynamicPseudoType) {
				// TODO: make this an attribute path error?
				return Value{}, fmt.Errorf("tftypes.NewValue can't use type %s as a value for %q in %s. Expected type is %s.", v.Type(), k, Map{AttributeType: typ}, typ)
			}
		}
		return Value{
			typ:   Map{AttributeType: typ},
			value: value,
		}, nil
	default:
		return Value{}, fmt.Errorf("tftypes.NewValue can't use %T as a tftypes.Map. Expected types are: %s", in, formattedSupportedGoTypes(Map{}))
	}
}
