package tftypes

type Set struct {
	ElementType Type
}

func (s Set) Is(t Type) bool {
	_, ok := t.(Set)
	return ok
}

func (s Set) String() string {
	return "tftypes.Set"
}

func (s Set) private() {}
