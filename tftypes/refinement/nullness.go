// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// TODO: doc
type Nullness struct {
	value bool
}

func (n Nullness) Equal(other Refinement) bool {
	otherVal, ok := other.(Nullness)
	if !ok {
		return false
	}

	return n.Nullness() == otherVal.Nullness()
}

func (n Nullness) String() string {
	if n.value {
		// This case should never happen, as an unknown value that is definitely null should be
		// represented as a known null value.
		return "null"
	}

	return "not null"
}

// TODO: doc
func (n Nullness) Nullness() bool {
	return n.value
}

func (n Nullness) unimplementable() {}

// TODO: doc
func NewNullness(value bool) Refinement {
	return Nullness{
		value: value,
	}
}
