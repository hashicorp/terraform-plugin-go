// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

// TestSetElementIndexMatchesScan checks that an ElementKeyValue step resolves
// to the same element, or the same error, with and without the index.
func TestSetElementIndexMatchesScan(t *testing.T) {
	t.Parallel()

	inner := Set{ElementType: String}
	elemType := Object{AttributeTypes: map[string]Type{
		"name":  String,
		"count": Number,
		"flag":  Bool,
		"tags":  inner,
		"list":  List{ElementType: Number},
	}}

	rng := rand.New(rand.NewSource(1))

	randomString := func() Value {
		switch rng.Intn(6) {
		case 0:
			return NewValue(String, nil)
		case 1:
			return NewValue(String, UnknownValue)
		}
		return NewValue(String, fmt.Sprintf("s%d", rng.Intn(4)))
	}
	randomNumber := func() Value {
		switch rng.Intn(6) {
		case 0:
			return NewValue(Number, nil)
		case 1:
			return NewValue(Number, UnknownValue)
		case 2:
			// Same numeric value at a different precision.
			return NewValue(Number, new(big.Float).SetPrec(200).SetInt64(int64(rng.Intn(3))))
		}
		return NewValue(Number, big.NewFloat(float64(rng.Intn(3))))
	}
	randomElement := func() Value {
		var tags []Value
		for i := rng.Intn(4); i > 0; i-- {
			tags = append(tags, randomString())
		}
		var list []Value
		for i := rng.Intn(3); i > 0; i-- {
			list = append(list, randomNumber())
		}
		return NewValue(elemType, map[string]Value{
			"name":  randomString(),
			"count": randomNumber(),
			"flag":  NewValue(Bool, rng.Intn(2) == 0),
			"tags":  NewValue(inner, tags),
			"list":  NewValue(List{ElementType: Number}, list),
		})
	}

	for round := 0; round < 200; round++ {
		var elems []Value
		for i := rng.Intn(40); i > 0; i-- {
			elems = append(elems, randomElement())
		}

		indexed := NewValue(Set{ElementType: elemType}, elems)
		scanned := Value{typ: indexed.typ, value: indexed.value}

		if indexed.setIndex == nil {
			t.Fatal("expected NewValue to attach an index to a Set")
		}

		steps := append([]Value{}, elems...)
		for i := 0; i < 20; i++ {
			steps = append(steps, randomElement())
		}
		steps = append(steps, NewValue(String, "wrong type"))

		for _, step := range steps {
			got, gotErr := indexed.ApplyTerraform5AttributePathStep(ElementKeyValue(step))
			want, wantErr := scanned.ApplyTerraform5AttributePathStep(ElementKeyValue(step))

			if (gotErr == nil) != (wantErr == nil) || (gotErr != nil && gotErr.Error() != wantErr.Error()) {
				t.Fatalf("round %d: error mismatch for %s: index %v, scan %v", round, step, gotErr, wantErr)
			}

			if gotErr != nil {
				continue
			}

			if !got.(Value).Equal(want.(Value)) {
				t.Fatalf("round %d: element mismatch for %s: index %s, scan %s", round, step, got, want)
			}
		}
	}
}

func TestSetElementIndexSkipsDynamicTypes(t *testing.T) {
	t.Parallel()

	elemType := Object{AttributeTypes: map[string]Type{"value": DynamicPseudoType}}
	a := NewValue(elemType, map[string]Value{"value": NewValue(String, "a")})
	b := NewValue(elemType, map[string]Value{"value": NewValue(Number, big.NewFloat(1))})
	set := NewValue(Set{ElementType: elemType}, []Value{a, b})

	if _, found := set.setIndex.lookup(elemType, a, []Value{a, b}); found {
		t.Fatal("expected the index to leave dynamic element types to the scan")
	}
}
