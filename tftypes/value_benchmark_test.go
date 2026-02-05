// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import "testing"

func BenchmarkValueApplyTerraform5AttributePathStep1000(b *testing.B) {
	benchmarkValueApplyTerraform5AttributePathStep(b, 1000)
}

// This benchmark iterates through an entire set of objects, which is one of the
// most expensive, but common, use cases.
func benchmarkValueApplyTerraform5AttributePathStep(b *testing.B, elements int) {
	// Set of objects is one of the most expensive operations
	objectType := Object{
		AttributeTypes: map[string]Type{
			"element_index": Number, // guaranteed to be different each element
			"test_string":   String,
		},
	}
	setType := Set{
		ElementType: objectType,
	}

	setElements := make([]Value, elements)

	for index := range setElements {
		setElements[index] = NewValue(
			objectType,
			map[string]Value{
				"element_index": NewValue(Number, index),
				"test_string":   NewValue(String, "test value"),
			},
		)
	}

	value := NewValue(
		setType,
		setElements,
	)

	// ensure iteration occurs through whole set
	step := ElementKeyValue(setElements[elements-1].Copy())

	for n := 0; n < b.N; n++ {
		_, err := value.ApplyTerraform5AttributePathStep(step)

		if err != nil {
			b.Fatalf("unexpected ApplyTerraform5AttributePathStep error: %s", err)
		}
	}
}
