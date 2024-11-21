// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import (
	"math/big"
)

// TODO: doc
type NumberUpperBound struct {
	inclusive bool
	value     *big.Float
}

func (n NumberUpperBound) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (n NumberUpperBound) String() string {
	// TODO: implement
	return "todo - NumberUpperBound"
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
