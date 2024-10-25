package refinement

import (
	"fmt"
)

type Key int64

func (k Key) String() string {
	// TODO: Not sure when this is used, double check the names
	switch k {
	case KeyNullness:
		return "nullness"
	case KeyStringPrefix:
		return "string_prefix"
	case KeyNumberLowerBound:
		return "number_lower_bound"
	case KeyNumberUpperBound:
		return "number_upper_bound"
	default:
		return fmt.Sprintf("unsupported refinement: %d", k)
	}
}

const (
	KeyNullness         = Key(1)
	KeyStringPrefix     = Key(2)
	KeyNumberLowerBound = Key(3)
	KeyNumberUpperBound = Key(4)
	// KeyCollectionLengthLowerBound = Key(5)
	// KeyCollectionLengthUpperBound = Key(6)
)

type Refinement interface {
	Equal(Refinement) bool
	String() string
	unimplementable() // prevents external implementations, all refinements are defined in the Terraform/HCL type system go-cty.
}

type Refinements map[Key]Refinement

func (r Refinements) Equal(o Refinements) bool {
	return false
}
func (r Refinements) String() string {
	// TODO: Not sure when this is used, should just aggregate and call all underlying refinements.String() method
	return "todo"
}
