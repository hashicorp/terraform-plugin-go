package tftypes

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
