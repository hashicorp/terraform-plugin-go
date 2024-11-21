// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// TODO: doc
type CollectionLengthLowerBound struct {
	value int64
}

func (n CollectionLengthLowerBound) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (n CollectionLengthLowerBound) String() string {
	// TODO: implement
	return "todo - CollectionLengthLowerBound"
}

// TODO: doc
func (n CollectionLengthLowerBound) LowerBound() int64 {
	return n.value
}

func (n CollectionLengthLowerBound) unimplementable() {}

// TODO: doc
func NewCollectionLengthLowerBound(value int64) Refinement {
	return CollectionLengthLowerBound{
		value: value,
	}
}
