package refinement

type Key int64

func (k Key) String() string {
	return "todo"
}

const (
	KeyNullness = Key(1)
	// KeyStringPrefix               = Key(2)
	// KeyNumberLowerBound           = Key(3)
	// KeyNumberUpperBound           = Key(4)
	// KeyCollectionLengthLowerBound = Key(5)
	// KeyCollectionLengthUpperBound = Key(6)
)

type Refinement interface {
	Equal(Refinement) bool
	String() string
	unimplementable() // prevent external implementations
}

type Refinements map[Key]Refinement

func (r Refinements) Equal(o Refinements) bool {
	return false
}
func (r Refinements) String() string {
	return "todo"
}
