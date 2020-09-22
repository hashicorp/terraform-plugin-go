package tftypes

import "encoding/json"

type Tuple struct {
	ElementTypes []Type
}

func (tu Tuple) Is(t Type) bool {
	_, ok := t.(Tuple)
	return ok
}

func (t Tuple) String() string {
	return "tftypes.Tuple"
}

func (t Tuple) private() {}

func (t Tuple) MarshalJSON() ([]byte, error) {
	elements, err := json.Marshal(t.ElementTypes)
	if err != nil {
		return nil, err
	}
	return []byte(`["tuple",` + string(elements) + `]`), nil
}
