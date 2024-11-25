// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package refinement

import (
	"fmt"
	"sort"
	"strings"
)

type Key int64

func (k Key) String() string {
	// TODO: Not sure when this is used, double check the names
	switch k {
	case KeyNullness:
		return "nullness"
	case KeyStringPrefix:
		return "string_prefix"
	case KeyNumberLowerBound:
		return "number_lower_bound"
	case KeyNumberUpperBound:
		return "number_upper_bound"
	default:
		return fmt.Sprintf("unsupported refinement: %d", k)
	}
}

const (
	// KeyNullness represents a refinement that specifies whether the final value will not be null.
	//
	// MAINTINAER NOTE: In practice, this refinement data will only contain "false", indicating the final value
	// cannot be null. If the refinement data was ever set to "true", that would indicate the final value will be null, in which
	// case the value is not unknown, it is known and should not have any refinement data.
	//
	// This refinement is relevant for all types except tftypes.DynamicPseudoType.
	KeyNullness = Key(1)

	// KeyStringPrefix represents a refinement that specifies a known prefix of a final string value.
	//
	// This refinement is only relevant for tftypes.String.
	KeyStringPrefix = Key(2)

	// KeyNumberLowerBound represents a refinement that specifies the lower bound of possible values for a final number value.
	// The refinement data contains a boolean which indicates whether the bound is inclusive.
	//
	// This refinement is only relevant for tftypes.Number.
	KeyNumberLowerBound = Key(3)

	// KeyNumberUpperBound represents a refinement that specifies the upper bound of possible values for a final number value.
	// The refinement data contains a boolean which indicates whether the bound is inclusive.
	//
	// This refinement is only relevant for tftypes.Number.
	KeyNumberUpperBound = Key(4)

	// KeyCollectionLengthLowerBound represents a refinement that specifies the lower bound of possible length for a final collection value.
	//
	// This refinement is only relevant for tftypes.List, tftypes.Set, and tftypes.Map.
	KeyCollectionLengthLowerBound = Key(5)

	// KeyCollectionLengthUpperBound represents a refinement that specifies the upper bound of possible length for a final collection value.
	//
	// This refinement is only relevant for tftypes.List, tftypes.Set, and tftypes.Map.
	KeyCollectionLengthUpperBound = Key(6)
)

// TODO: docs
type Refinement interface {
	// Equal should return true if the Refinement is considered equivalent to the
	// Refinement passed as an argument.
	Equal(Refinement) bool

	// String should return a human-friendly version of the Refinement.
	String() string

	unimplementable() // prevents external implementations, all refinements are defined in the Terraform/HCL type system go-cty.
}

// TODO: docs
type Refinements map[Key]Refinement

func (r Refinements) Equal(other Refinements) bool {
	if len(r) != len(other) {
		return false
	}

	for key, refnVal := range r {
		otherRefnVal, ok := other[key]
		if !ok {
			// Didn't find a refinement at the same key
			return false
		}

		if !refnVal.Equal(otherRefnVal) {
			// Refinement data is not equal
			return false
		}
	}

	return true
}
func (r Refinements) String() string {
	var res strings.Builder

	keys := make([]Key, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(a, b int) bool { return keys[a] < keys[b] })
	for pos, key := range keys {
		if pos != 0 {
			res.WriteString(", ")
		}
		res.WriteString(r[key].String())
	}

	return res.String()
}
