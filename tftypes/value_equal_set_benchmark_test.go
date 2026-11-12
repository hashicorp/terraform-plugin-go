// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import (
	"fmt"
	"testing"
)

func benchmarkSetNestedValue(n int) Value {
	elemType := Object{AttributeTypes: map[string]Type{
		"name":  String,
		"note":  String,
		"count": Number,
	}}
	elems := make([]Value, n)
	for i := range elems {
		elems[i] = NewValue(elemType, map[string]Value{
			"name":  NewValue(String, fmt.Sprintf("element-%d", i)),
			"note":  NewValue(String, nil),
			"count": NewValue(Number, nil),
		})
	}
	setType := Set{ElementType: elemType}
	return NewValue(
		Object{AttributeTypes: map[string]Type{"items": setType}},
		map[string]Value{"items": NewValue(setType, elems)},
	)
}

func BenchmarkValueEqualSetNested(b *testing.B) {
	for _, n := range []int{100, 1000, 3251} {
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B) {
			v1 := benchmarkSetNestedValue(n)
			v2 := benchmarkSetNestedValue(n)
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if !v1.Equal(v2) {
					b.Fatal("expected equal values")
				}
			}
		})
	}
}
