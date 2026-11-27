// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tftypes

import (
	"hash/maphash"
	"math"
	"math/big"
	"sort"
	"sync"
)

// setElementIndex lets an ElementKeyValue step find its element without
// comparing the step against every element of the set. It only narrows the
// candidates: every hit is still confirmed with deepEqual, and any case it
// cannot answer falls back to the linear scan.
type setElementIndex struct {
	once    sync.Once
	size    int
	buckets map[uint64][]int
}

var setElementIndexSeed = maphash.MakeSeed()

// lookup returns the first element of elems that deepEquals step. The second
// return value is false when the index cannot answer, in which case the caller
// must scan elems itself.
func (idx *setElementIndex) lookup(elementType Type, step Value, elems []Value) (Value, bool) {
	if len(elems) == 0 || step.Type() == nil {
		return Value{}, false
	}

	idx.once.Do(func() {
		// Dynamic types can hold values of differing concrete types, which the
		// scan reports as errors. Leave those to the scan.
		if typeContainsDynamic(elementType) {
			return
		}

		buckets := make(map[uint64][]int, len(elems))

		for i, el := range elems {
			if el.Type() == nil || !el.Type().Equal(elems[0].Type()) {
				return
			}

			h, ok := hashValue(el)

			if !ok {
				return
			}

			buckets[h] = append(buckets[h], i)
		}

		idx.size = len(elems)
		idx.buckets = buckets
	})

	if idx.buckets == nil || idx.size != len(elems) || !step.Type().Equal(elems[0].Type()) {
		return Value{}, false
	}

	h, ok := hashValue(step)

	if !ok {
		return Value{}, false
	}

	for _, i := range idx.buckets[h] {
		equal, err := step.deepEqual(elems[i])

		if err != nil {
			return Value{}, false
		}

		if equal {
			return elems[i], true
		}
	}

	return Value{}, false
}

func typeContainsDynamic(t Type) bool {
	switch typ := t.(type) {
	case primitive:
		return typ.name == DynamicPseudoType.name
	case List:
		return typeContainsDynamic(typ.ElementType)
	case Set:
		return typeContainsDynamic(typ.ElementType)
	case Map:
		return typeContainsDynamic(typ.ElementType)
	case Tuple:
		for _, elementType := range typ.ElementTypes {
			if typeContainsDynamic(elementType) {
				return true
			}
		}

		return false
	case Object:
		for _, attributeType := range typ.AttributeTypes {
			if typeContainsDynamic(attributeType) {
				return true
			}
		}

		return false
	}

	return true
}

const (
	hashTagUnknown uint64 = 0x9e3779b97f4a7c15
	hashTagNull    uint64 = 0xc2b2ae3d27d4eb4f
	hashTagList    uint64 = 0x165667b19e3779f9
	hashTagSet     uint64 = 0x27d4eb2f165667c5
	hashTagMap     uint64 = 0x85ebca77c2b2ae63
)

func hashMix(h, v uint64) uint64 {
	h ^= v + 0x9e3779b97f4a7c15 + (h << 6) + (h >> 2)
	return h
}

// hashValue returns a hash that is equal for any two values deepEqual reports
// as equal. The second return value is false for values it does not handle.
func hashValue(val Value) (uint64, bool) {
	if val.Type() == nil {
		return 0, false
	}

	if !val.IsKnown() {
		return hashTagUnknown, true
	}

	if val.IsNull() {
		return hashTagNull, true
	}

	switch v := val.value.(type) {
	case string:
		return maphash.String(setElementIndexSeed, v), true
	case bool:
		if v {
			return 1, true
		}

		return 2, true
	case *big.Float:
		// deepEqual compares numbers with Cmp, so hash the numeric value and
		// not the representation.
		f, _ := v.Float64()

		if f == 0 {
			f = 0
		}

		return math.Float64bits(f), true
	case []Value:
		if _, ok := val.Type().(Set); ok {
			hashes := make([]uint64, 0, len(v))

			for _, el := range v {
				h, ok := hashValue(el)

				if !ok {
					return 0, false
				}

				hashes = append(hashes, h)
			}

			sort.Slice(hashes, func(i, j int) bool { return hashes[i] < hashes[j] })

			result := hashTagSet

			for i, h := range hashes {
				if i > 0 && h == hashes[i-1] {
					continue
				}

				result = hashMix(result, h)
			}

			return result, true
		}

		result := hashTagList

		for _, el := range v {
			h, ok := hashValue(el)

			if !ok {
				return 0, false
			}

			result = hashMix(result, h)
		}

		return result, true
	case map[string]Value:
		result := hashTagMap

		for k, el := range v {
			h, ok := hashValue(el)

			if !ok {
				return 0, false
			}

			result += hashMix(maphash.String(setElementIndexSeed, k), h)
		}

		return result, true
	}

	return 0, false
}
