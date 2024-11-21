// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import (
	"math/big"
)

// TODO: doc
type NumberLowerBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberLowerBound) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (n NumberLowerBound) String() string {
	// TODO: implement
	return "todo - NumberLowerBound"
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
