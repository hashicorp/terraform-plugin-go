// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import "fmt"

// TODO: doc
type CollectionLengthUpperBound struct {
	value int64
}

func (n CollectionLengthUpperBound) Equal(other Refinement) bool {
	otherVal, ok := other.(CollectionLengthUpperBound)
	if !ok {
		return false
	}

	return n.UpperBound() == otherVal.UpperBound()
}

func (n CollectionLengthUpperBound) String() string {
	return fmt.Sprintf("length upper bound = %d", n.UpperBound())
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
