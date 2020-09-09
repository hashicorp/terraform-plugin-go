package tftypes

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
