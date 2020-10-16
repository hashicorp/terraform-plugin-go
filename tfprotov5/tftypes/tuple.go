package tftypes

import "encoding/json"

type Tuple struct {
	ElementTypes []Type
}

func (tu Tuple) Is(t Type) bool {
	v, ok := t.(Tuple)
	if !ok {
		return false
	}
	if v.ElementTypes != nil {
		if len(v.ElementTypes) != len(tu.ElementTypes) {
			return false
		}
		for pos, typ := range tu.ElementTypes {
			if !typ.Is(v.ElementTypes[pos]) {
				return false
			}
		}
	}
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
