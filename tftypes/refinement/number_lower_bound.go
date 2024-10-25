package refinement

import (
	"math/big"
)

type NumberLowerBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberLowerBound) Equal(Refinement) bool {
	return false
}

func (n NumberLowerBound) String() string {
	return "todo - NumberLowerBound"
}

func (n NumberLowerBound) IsInclusive() bool {
	return n.inclusive
}

func (n NumberLowerBound) LowerBound() *big.Float {
	return n.value
}

func (n NumberLowerBound) unimplementable() {}

func NewNumberLowerBound(value *big.Float, inclusive bool) Refinement {
	return NumberLowerBound{
		value:     value,
		inclusive: inclusive,
	}
}
