// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import (
	"fmt"
	"math/big"
)

// TODO: doc
type NumberUpperBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberUpperBound) Equal(other Refinement) bool {
	otherVal, ok := other.(NumberUpperBound)
	if !ok {
		return false
	}

	return n.IsInclusive() == otherVal.IsInclusive() && n.UpperBound().Cmp(otherVal.UpperBound()) == 0
}

func (n NumberUpperBound) String() string {
	rangeDescription := "inclusive"
	if !n.IsInclusive() {
		rangeDescription = "exclusive"
	}

	return fmt.Sprintf("upper bound = %s (%s)", n.UpperBound().String(), rangeDescription)
}

// TODO: doc
func (n NumberUpperBound) IsInclusive() bool {
	return n.inclusive
}

// TODO: doc
func (n NumberUpperBound) UpperBound() *big.Float {
	return n.value
}

func (n NumberUpperBound) unimplementable() {}

// TODO: doc
func NewNumberUpperBound(value *big.Float, inclusive bool) Refinement {
	return NumberUpperBound{
		value:     value,
		inclusive: inclusive,
	}
}
