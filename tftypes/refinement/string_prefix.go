// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// TODO: doc
type StringPrefix struct {
	value string
}

func (s StringPrefix) Equal(Refinement) bool {
	// TODO: implement
	return false
}

func (s StringPrefix) String() string {
	// TODO: implement
	return "todo - stringPrefix"
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
