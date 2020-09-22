package tftypes

import "fmt"

type List struct {
	ElementType Type
}

func (l List) Is(t Type) bool {
	_, ok := t.(List)
	return ok
}

func (l List) String() string {
	return "tftypes.List"
}

func (l List) private() {}

func (l List) MarshalJSON() ([]byte, error) {
	elementType, err := l.ElementType.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("error marshaling tftypes.List's element type %T to JSON: %w", l.ElementType, err)
	}
	return []byte(`["list",` + string(elementType) + `]`), nil
}
