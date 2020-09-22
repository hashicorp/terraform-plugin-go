package tftypes

import "fmt"

type Map struct {
	AttributeType Type
}

func (m Map) Is(t Type) bool {
	_, ok := t.(Map)
	return ok
}

func (m Map) String() string {
	return "tftypes.Map"
}

func (m Map) private() {}

func (m Map) MarshalJSON() ([]byte, error) {
	attributeType, err := m.AttributeType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.Map's attribute type %T to JSON: %w", m.AttributeType, err)
	}
	return []byte(`["map",` + string(attributeType) + `]`), nil
}
