// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

// Nullness represents an unknown value refinement that indicates the final value will definitely not be null (Nullness = false). This refinement
// can be applied to a value of any type (excluding DynamicPseudoType).
//
// While an unknown value can be refined to indicate that the final value will definitely be null (Nullness = true), there is no practical reason
// to do this. This option is exposed to maintain parity with Terraform's type system, while all practical usages of this refinement should collapse
// to known null values instead.
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

// Nullness returns the underlying refinement data indicating:
//   - When "false", the final value will definitely not be null.
//   - When "true", the final value will definitely be null.
//
// While an unknown value can be refined to indicate that the final value will definitely be null (Nullness = true), there is no practical reason
// to do this. This option is exposed to maintain parity with Terraform's type system, while all practical usages of this refinement should collapse
// to known null values instead.
func (n Nullness) Nullness() bool {
	return n.value
}

func (n Nullness) unimplementable() {}

// NewNullness returns the Nullness unknown value refinement that indicates the final value will definitely not be null (Nullness = false). This refinement
// can be applied to a value of any type (excluding DynamicPseudoType).
//
// While an unknown value can be refined to indicate that the final value will definitely be null (Nullness = true), there is no practical reason
// to do this. This option is exposed to maintain parity with Terraform's type system, while all practical usages of this refinement should collapse
// to known null values instead.
func NewNullness(value bool) Refinement {
	return Nullness{
		value: value,
	}
}
