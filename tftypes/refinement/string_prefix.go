// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import "fmt"

// TODO: doc
type StringPrefix struct {
	value string
}

func (s StringPrefix) Equal(other Refinement) bool {
	otherVal, ok := other.(StringPrefix)
	if !ok {
		return false
	}

	return s.PrefixValue() == otherVal.PrefixValue()
}

func (s StringPrefix) String() string {
	return fmt.Sprintf("prefix = %q", s.PrefixValue())
}

// TODO: doc
func (s StringPrefix) PrefixValue() string {
	return s.value
}

func (s StringPrefix) unimplementable() {}

// TODO: doc
func NewStringPrefix(value string) Refinement {
	return StringPrefix{
		value: value,
	}
}
