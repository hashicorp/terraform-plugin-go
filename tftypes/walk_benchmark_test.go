package tftypes

import "testing"

func BenchmarkTransform1000(b *testing.B) {
	benchmarkTransform(b, 1000)
}

// This benchmark iterates through an entire set of objects, which is one of the
// most expensive, but common, use cases.
func benchmarkTransform(b *testing.B, elements int) {
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

	for n := 0; n < b.N; n++ {
		_, err := Transform(
			value,
			func(_ *AttributePath, value Value) (Value, error) {
				return value, nil
			},
		)

		if err != nil {
			b.Fatalf("unexpected Transform error: %s", err)
		}
	}
}
