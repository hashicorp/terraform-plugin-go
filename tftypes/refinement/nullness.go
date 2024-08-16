package refinement

type nullness struct {
	Value bool
}

func (n nullness) Equal(Refinement) bool {
	return false
}

func (n nullness) String() string {
	return "todo - nullness"
}

func (n nullness) unimplementable() {}

func Nullness(value bool) Refinement {
	return nullness{
		Value: value,
	}
}
