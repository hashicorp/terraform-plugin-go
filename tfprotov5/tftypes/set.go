package tftypes

import "fmt"

type Set struct {
	ElementType Type
}

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
	return "tftypes.Set"
}

func (s Set) private() {}

func (s Set) MarshalJSON() ([]byte, error) {
	elementType, err := s.ElementType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.Set's element type %T to JSON: %w", s.ElementType, err)
	}
	return []byte(`["set",` + string(elementType) + `]`), nil
}
