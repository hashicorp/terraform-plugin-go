// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// TODO: doc
type Nullness struct {
	value bool
}

func (n Nullness) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (n Nullness) String() string {
	// TODO: implement
	return "todo - Nullness"
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
