package refinement

import (
	"math/big"
)

type NumberUpperBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberUpperBound) Equal(Refinement) bool {
	return false
}

func (n NumberUpperBound) String() string {
	return "todo - NumberUpperBound"
}

func (n NumberUpperBound) IsInclusive() bool {
	return n.inclusive
}

func (n NumberUpperBound) UpperBound() *big.Float {
	return n.value
}

func (n NumberUpperBound) unimplementable() {}

func NewNumberUpperBound(value *big.Float, inclusive bool) Refinement {
	return NumberUpperBound{
		value:     value,
		inclusive: inclusive,
	}
}
