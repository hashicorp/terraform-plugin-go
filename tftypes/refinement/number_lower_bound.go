// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import (
	"fmt"
	"math/big"
)

// TODO: doc
type NumberLowerBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberLowerBound) Equal(other Refinement) bool {
	otherVal, ok := other.(NumberLowerBound)
	if !ok {
		return false
	}

	return n.IsInclusive() == otherVal.IsInclusive() && n.LowerBound().Cmp(otherVal.LowerBound()) == 0
}

func (n NumberLowerBound) String() string {
	rangeDescription := "inclusive"
	if !n.IsInclusive() {
		rangeDescription = "exclusive"
	}

	return fmt.Sprintf("lower bound = %s (%s)", n.LowerBound().String(), rangeDescription)
}

// TODO: doc
func (n NumberLowerBound) IsInclusive() bool {
	return n.inclusive
}

// TODO: doc
func (n NumberLowerBound) LowerBound() *big.Float {
	return n.value
}

func (n NumberLowerBound) unimplementable() {}

// TODO: doc
func NewNumberLowerBound(value *big.Float, inclusive bool) Refinement {
	return NumberLowerBound{
		value:     value,
		inclusive: inclusive,
	}
}
