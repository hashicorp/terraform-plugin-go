package tftypes

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
