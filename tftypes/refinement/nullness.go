package refinement

type Nullness struct {
	value bool
}

func (n Nullness) Equal(Refinement) bool {
	return false
}

func (n Nullness) String() string {
	return "todo - Nullness"
}

func (n Nullness) Nullness() bool {
	return n.value
}

func (n Nullness) unimplementable() {}

func NewNullness(value bool) Refinement {
	return Nullness{
		value: value,
	}
}
