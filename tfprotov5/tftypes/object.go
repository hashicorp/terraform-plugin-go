package tftypes

import "encoding/json"

type Object struct {
	AttributeTypes map[string]Type
}

func (o Object) Is(t Type) bool {
	_, ok := t.(Object)
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
