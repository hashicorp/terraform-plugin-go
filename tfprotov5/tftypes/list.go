package tftypes

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
