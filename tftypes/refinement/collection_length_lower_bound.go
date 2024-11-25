// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import "fmt"

// TODO: doc
type CollectionLengthLowerBound struct {
	value int64
}

func (n CollectionLengthLowerBound) Equal(other Refinement) bool {
	otherVal, ok := other.(CollectionLengthLowerBound)
	if !ok {
		return false
	}

	return n.LowerBound() == otherVal.LowerBound()
}

func (n CollectionLengthLowerBound) String() string {
	return fmt.Sprintf("length lower bound = %d", n.LowerBound())
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
