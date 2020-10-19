package tftypes

import "encoding/json"

type Object struct {
	AttributeTypes map[string]Type
}

func (o Object) Is(t Type) bool {
	v, ok := t.(Object)
	if !ok {
		return false
	}
	if v.AttributeTypes != nil {
		if len(o.AttributeTypes) != len(v.AttributeTypes) {
			return false
		}
		for k, typ := range o.AttributeTypes {
			if _, ok := v.AttributeTypes[k]; !ok {
				return false
			}
			if !typ.Is(v.AttributeTypes[k]) {
				return false
			}
		}
	}
	return ok
}

func (o Object) String() string {
	return "tftypes.Object"
}

func (o Object) private() {}

func (o Object) MarshalJSON() ([]byte, error) {
	attrs, err := json.Marshal(o.AttributeTypes)
	if err != nil {
		return nil, err
	}
	return []byte(`["object",` + string(attrs) + `]`), nil
}
