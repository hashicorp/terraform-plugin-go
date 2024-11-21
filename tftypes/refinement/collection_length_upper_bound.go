// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// TODO: doc
type CollectionLengthUpperBound struct {
	value int64
}

func (n CollectionLengthUpperBound) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (n CollectionLengthUpperBound) String() string {
	// TODO: implement
	return "todo - CollectionLengthUpperBound"
}

// TODO: doc
func (n CollectionLengthUpperBound) UpperBound() int64 {
	return n.value
}

func (n CollectionLengthUpperBound) unimplementable() {}

// TODO: doc
func NewCollectionLengthUpperBound(value int64) Refinement {
	return CollectionLengthUpperBound{
		value: value,
	}
}
