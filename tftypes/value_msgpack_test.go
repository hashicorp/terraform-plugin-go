// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tftypes

import (
	"encoding/hex"
	"errors"
	"math"
	"math/big"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tftypes/refinement"
)

// Hex encoding of the long prefix refinements used in this test
var longPrefixRefinement = "c801050c8201c202d9ff7072656669783a2f2f303132333435363738392d303132333435363738392d303132333435363738392d30313" +
	"2333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d30313" +
	"2333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d30313" +
	"2333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d303132333435363738392d30313" +
	"2333435363738392d30313233"

func TestValueFromMsgPack(t *testing.T) {
	t.Parallel()
	type testCase struct {
		hex   string
		value Value
		typ   Type
	}
	bigNumber, _, err := big.ParseFloat("9999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999999", 10, 512, big.ToNearestEven)
	if err != nil {
		t.Fatalf("error parsing big number: %s", err)
	}
	awkwardFraction, _, err := big.ParseFloat("0.8", 10, 512, big.ToNearestEven)
	if err != nil {
		t.Fatalf("error parsing awkward fraction: %s", err)
	}

	// integer under 64 bits which rounds incorrectly if parsed as a float64
	uint64AsFloat, _ := new(big.Float).SetString("9223372036854775808")

	tests := map[string]testCase{
		"hello-string": {
			hex:   "a568656c6c6f",
			value: NewValue(String, "hello"),
			typ:   String,
		},
		"empty-string": {
			hex:   "a0",
			value: NewValue(String, ""),
			typ:   String,
		},
		"null-string": {
			hex:   "c0",
			value: NewValue(String, nil),
			typ:   String,
		},
		"unknown-string": {
			hex:   "d40000",
			value: NewValue(String, UnknownValue),
			typ:   String,
		},
		"true-bool": {
			hex:   "c3",
			value: NewValue(Bool, true),
			typ:   Bool,
		},
		"false-bool": {
			hex:   "c2",
			value: NewValue(Bool, false),
			typ:   Bool,
		},
		"null-bool": {
			hex:   "c0",
			value: NewValue(Bool, nil),
			typ:   Bool,
		},
		"unknown-bool": {
			hex:   "d40000",
			value: NewValue(Bool, UnknownValue),
			typ:   Bool,
		},
		"int-number": {
			hex:   "01",
			value: NewValue(Number, big.NewFloat(1)),
			typ:   Number,
		},
		"int64-positive-number": {
			hex:   "cf7fffffffffffffff",
			value: NewValue(Number, new(big.Float).SetInt64(math.MaxInt64)),
			typ:   Number,
		},
		"int64-negative-number": {
			hex:   "d38000000000000000",
			value: NewValue(Number, new(big.Float).SetInt64(math.MinInt64)),
			typ:   Number,
		},
		"uint64-number": {
			hex:   "b43138343436373434303733373039353531363135",
			value: NewValue(Number, new(big.Float).SetUint64(math.MaxUint64)),
			typ:   Number,
		},
		"float-number": {
			hex:   "cb3ff8000000000000",
			value: NewValue(Number, big.NewFloat(1.5)),
			typ:   Number,
		},
		"float64-positive-number": {
			// Because MaxFloat64 is an integer value, it must be encoded as an integer to ensure we don't lose precision when decoding the value
			hex:   "da0135313739373639333133343836323331353730303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030303030",
			value: NewValue(Number, new(big.Float).SetFloat64(math.MaxFloat64)),
			typ:   Number,
		},
		"float64-negative-number": {
			hex:   "cb0000000000000001",
			value: NewValue(Number, new(big.Float).SetFloat64(math.SmallestNonzeroFloat64)),
			typ:   Number,
		},
		"big-number": {
			hex:   "d96439393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939",
			value: NewValue(Number, bigNumber),
			typ:   Number,
		},
		"awkward-fraction-number": {
			hex:   "a3302e38",
			value: NewValue(Number, awkwardFraction),
			typ:   Number,
		},
		"positive-infinity-number": {
			hex:   "cb7ff0000000000000",
			value: NewValue(Number, big.NewFloat(math.Inf(1))),
			typ:   Number,
		},
		"negative-infinity-number": {
			hex:   "cbfff0000000000000",
			value: NewValue(Number, big.NewFloat(math.Inf(-1))),
			typ:   Number,
		},
		"large-uint64": {
			hex:   "b339323233333732303336383534373735383038",
			value: NewValue(Number, uint64AsFloat),
			typ:   Number,
		},
		"dynamic-bool": {
			hex:   "92c40622626f6f6c22c3",
			value: NewValue(Bool, true),
			typ:   DynamicPseudoType,
		},
		"dynamic-list": {
			hex: "92c4115b226c697374222c22737472696e67225d91a568656c6c6f",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: DynamicPseudoType,
		},
		"dynamic-map": {
			hex: "92c4105b226d6170222c22737472696e67225d81a568656c6c6fa568656c6c6f",
			value: NewValue(Map{
				ElementType: String,
			}, map[string]Value{
				"hello": NewValue(String, "hello"),
			}),
			typ: DynamicPseudoType,
		},
		"dynamic-number": {
			hex:   "92c408226e756d6265722201",
			value: NewValue(Number, big.NewFloat(1)),
			typ:   DynamicPseudoType,
		},
		"dynamic-object": {
			hex: "81a86772656574696e6792c40822737472696e6722a568656c6c6f",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"greeting": DynamicPseudoType,
				},
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"greeting": DynamicPseudoType,
				},
			},
		},
		"dynamic-set": {
			hex: "92c4105b22736574222c22737472696e67225d91a568656c6c6f",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: DynamicPseudoType,
		},
		"dynamic-string": {
			hex:   "92c40822737472696e6722a568656c6c6f",
			value: NewValue(String, "hello"),
			typ:   DynamicPseudoType,
		},
		"list-dynamic": {
			hex: "9192c40822737472696e6722a568656c6c6f",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: List{ElementType: DynamicPseudoType},
		},
		"list-string-hello": {
			hex: "91a568656c6c6f",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: List{ElementType: String},
		},
		"list-string-unknown": {
			hex: "91d40000",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, UnknownValue),
			}),
			typ: List{ElementType: String},
		},
		"list-string-null-string": {
			hex: "91c0",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, nil),
			}),
			typ: List{ElementType: String},
		},
		"list-string-null": {
			hex: "c0",
			value: NewValue(List{
				ElementType: String,
			}, nil),
			typ: List{ElementType: String},
		},
		"list-string-empty": {
			hex: "90",
			value: NewValue(List{
				ElementType: String,
			}, []Value{}),
			typ: List{ElementType: String},
		},
		"set-dynamic": {
			hex: "9192c40822737472696e6722a568656c6c6f",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: Set{ElementType: DynamicPseudoType},
		},
		"set-string-hello": {
			hex: "91a568656c6c6f",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: Set{ElementType: String},
		},
		"set-string-unknown": {
			hex: "91d40000",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{
				NewValue(String, UnknownValue),
			}),
			typ: Set{ElementType: String},
		},
		"set-string-null-string": {
			hex: "91c0",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{
				NewValue(String, nil),
			}),
			typ: Set{ElementType: String},
		},
		"set-string-empty": {
			hex: "90",
			value: NewValue(Set{
				ElementType: String,
			}, []Value{}),
			typ: Set{ElementType: String},
		},
		"map-dynamic": {
			hex: "81a86772656574696e6792c40822737472696e6722a568656c6c6f",
			value: NewValue(Map{
				ElementType: DynamicPseudoType,
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Map{ElementType: DynamicPseudoType},
		},
		"map-string-hello": {
			hex: "81a86772656574696e67a568656c6c6f",
			value: NewValue(Map{
				ElementType: String,
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Map{ElementType: String},
		},
		"map-string-unknown": {
			hex: "81a86772656574696e67d40000",
			value: NewValue(Map{
				ElementType: String,
			}, map[string]Value{
				"greeting": NewValue(String, UnknownValue),
			}),
			typ: Map{ElementType: String},
		},
		"map-string-null": {
			hex: "81a86772656574696e67c0",
			value: NewValue(Map{
				ElementType: String,
			}, map[string]Value{
				"greeting": NewValue(String, nil),
			}),
			typ: Map{ElementType: String},
		},
		"map-string-empty": {
			hex: "80",
			value: NewValue(Map{
				ElementType: String,
			}, map[string]Value{}),
			typ: Map{ElementType: String},
		},
		"tuple-dynamic": {
			hex: "9292c40822737472696e6722a568656c6c6f92c40822737472696e6722a5776f726c64",
			value: NewValue(Tuple{
				ElementTypes: []Type{DynamicPseudoType, DynamicPseudoType},
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			}),
			typ: Tuple{ElementTypes: []Type{DynamicPseudoType, DynamicPseudoType}},
		},
		"tuple-string-hello": {
			hex: "91a568656c6c6f",
			value: NewValue(Tuple{
				ElementTypes: []Type{String},
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: Tuple{ElementTypes: []Type{String}},
		},
		"tuple-string-unknown": {
			hex: "91d40000",
			value: NewValue(Tuple{
				ElementTypes: []Type{String},
			}, []Value{
				NewValue(String, UnknownValue),
			}),
			typ: Tuple{ElementTypes: []Type{String}},
		},
		"tuple-string-null": {
			hex: "91c0",
			value: NewValue(Tuple{
				ElementTypes: []Type{String},
			}, []Value{
				NewValue(String, nil),
			}),
			typ: Tuple{ElementTypes: []Type{String}},
		},
		"tuple-empty": {
			hex: "90",
			value: NewValue(Tuple{
				ElementTypes: []Type{},
			}, []Value{}),
			typ: Tuple{ElementTypes: []Type{}},
		},
		"object-dynamic": {
			hex: "81a86772656574696e6792c40822737472696e6722a568656c6c6f",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"greeting": DynamicPseudoType,
				},
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"greeting": DynamicPseudoType,
				},
			},
		},
		"object-string-hello": {
			hex: "81a86772656574696e67a568656c6c6f",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			},
		},
		"object-string-unknown": {
			hex: "81a86772656574696e67d40000",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			}, map[string]Value{
				"greeting": NewValue(String, UnknownValue),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			},
		},
		"object-string-null": {
			hex: "81a86772656574696e67c0",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			}, map[string]Value{
				"greeting": NewValue(String, nil),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"greeting": String,
				},
			},
		},
		"object-string-multi-null": {
			hex: "82a161c0a162c0",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": String,
				},
			}, map[string]Value{
				"a": NewValue(String, nil),
				"b": NewValue(String, nil),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": String,
				},
			},
		},
		"object-string-multi-unknown": {
			hex: "82a161d40000a162d40000",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": String,
				},
			}, map[string]Value{
				"a": NewValue(String, UnknownValue),
				"b": NewValue(String, UnknownValue),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": String,
				},
			},
		},
		"object-empty": {
			hex: "80",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{},
			}, map[string]Value{}),
			typ: Object{
				AttributeTypes: map[string]Type{},
			},
		},
		"string-dynamic-null": {
			hex:   "92c40822737472696e6722c0",
			value: NewValue(String, nil),
			typ:   DynamicPseudoType,
		},
		"dynamic-null": {
			hex:   "c0",
			value: NewValue(DynamicPseudoType, nil),
			typ:   DynamicPseudoType,
		},
		"object-dynamic-null": {
			hex: "81a161c0",
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"a": DynamicPseudoType,
				},
			}, map[string]Value{
				"a": NewValue(DynamicPseudoType, nil),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": DynamicPseudoType,
				},
			},
		},
		"dynamic-unknown": {
			hex:   "d40000",
			value: NewValue(DynamicPseudoType, UnknownValue),
			typ:   DynamicPseudoType,
		},
		"dynamic-list-string-hello": {
			hex: "9192c40822737472696e6722a568656c6c6f",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
			}),
			typ: List{ElementType: DynamicPseudoType},
		},
		"dynamic-list-string-null": {
			hex: "9192c40822737472696e6722c0",
			value: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, nil),
			}),
			typ: List{ElementType: DynamicPseudoType},
		},
		"dynamic-list-unknown": {
			hex: "91d40000",
			value: NewValue(List{
				ElementType: DynamicPseudoType,
			}, []Value{
				NewValue(DynamicPseudoType, UnknownValue),
			}),
			typ: List{ElementType: DynamicPseudoType},
		},
		"unknown-dynamic-refinements-ignored": {
			hex: "d40000",
			value: NewValue(DynamicPseudoType, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: DynamicPseudoType,
		},
		"unknown-bool-with-nullness-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(Bool, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: Bool,
		},
		"unknown-number-with-nullness-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: Number,
		},
		"unknown-string-with-nullness-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: String,
		},
		"unknown-list-with-nullness-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(List{ElementType: String}, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: List{ElementType: String},
		},
		"unknown-object-with-nullness-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(Object{AttributeTypes: map[string]Type{"attr": String}}, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
			}),
			typ: Object{AttributeTypes: map[string]Type{"attr": String}},
		},
		"unknown-string-with-empty-prefix-refinement": {
			hex: "c7030c8101c2",
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:     refinement.NewNullness(false),
				refinement.KeyStringPrefix: refinement.NewStringPrefix(""),
			}),
			typ: String,
		},
		"unknown-string-with-prefix-refinement": {
			hex: "c70e0c8201c202a97072656669783a2f2f",
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:     refinement.NewNullness(false),
				refinement.KeyStringPrefix: refinement.NewStringPrefix("prefix://"),
			}),
			typ: String,
		},
		"unknown-string-with-long-prefix-refinement-one": {
			// This prefix will be cutoff at 256 bytes, so it will be equal to the other long prefix test.
			hex: longPrefixRefinement,
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
				refinement.KeyStringPrefix: refinement.NewStringPrefix(
					"prefix://0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-" +
						"0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-" +
						"0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-thiswillbecutoff1",
				),
			}),
			typ: String,
		},
		"unknown-string-with-long-prefix-refinement-two": {
			// This prefix will be cutoff at 256 bytes, so it will be equal to the other long prefix test.
			hex: longPrefixRefinement,
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness: refinement.NewNullness(false),
				refinement.KeyStringPrefix: refinement.NewStringPrefix(
					"prefix://0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-" +
						"0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-" +
						"0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-thiswillbecutoff2",
				),
			}),
			typ: String,
		},
		"unknown-number-with-bound-refinements-integers-inclusive": {
			hex: "c70b0c8301c2039201c3049205c3",
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:         refinement.NewNullness(false),
				refinement.KeyNumberLowerBound: refinement.NewNumberLowerBound(big.NewFloat(1), true),
				refinement.KeyNumberUpperBound: refinement.NewNumberUpperBound(big.NewFloat(5), true),
			}),
			typ: Number,
		},
		"unknown-number-with-bound-refinements-integers-exclusive": {
			hex: "c70b0c8301c2039201c2049205c2",
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:         refinement.NewNullness(false),
				refinement.KeyNumberLowerBound: refinement.NewNumberLowerBound(big.NewFloat(1), false),
				refinement.KeyNumberUpperBound: refinement.NewNumberUpperBound(big.NewFloat(5), false),
			}),
			typ: Number,
		},
		"unknown-number-with-bound-refinements-float-inclusive": {
			hex: "c71b0c8301c20392cb3ff3ae147ae147aec30492cb4016ae147ae147aec3",
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:         refinement.NewNullness(false),
				refinement.KeyNumberLowerBound: refinement.NewNumberLowerBound(big.NewFloat(1.23), true),
				refinement.KeyNumberUpperBound: refinement.NewNumberUpperBound(big.NewFloat(5.67), true),
			}),
			typ: Number,
		},
		"unknown-number-with-bound-refinements-float-exclusive": {
			hex: "c71b0c8301c20392cb3ff3ae147ae147aec20492cb4016ae147ae147aec2",
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:         refinement.NewNullness(false),
				refinement.KeyNumberLowerBound: refinement.NewNumberLowerBound(big.NewFloat(1.23), false),
				refinement.KeyNumberUpperBound: refinement.NewNumberUpperBound(big.NewFloat(5.67), false),
			}),
			typ: Number,
		},
		"unknown-list-with-bound-refinements": {
			hex: "c7070c8301c205010605",
			value: NewValue(List{ElementType: String}, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:                   refinement.NewNullness(false),
				refinement.KeyCollectionLengthLowerBound: refinement.NewCollectionLengthLowerBound(1),
				refinement.KeyCollectionLengthUpperBound: refinement.NewCollectionLengthUpperBound(5),
			}),
			typ: List{ElementType: String},
		},
		"unknown-map-with-bound-refinements": {
			hex: "c7070c8301c205000604",
			value: NewValue(Map{ElementType: Number}, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:                   refinement.NewNullness(false),
				refinement.KeyCollectionLengthLowerBound: refinement.NewCollectionLengthLowerBound(0),
				refinement.KeyCollectionLengthUpperBound: refinement.NewCollectionLengthUpperBound(4),
			}),
			typ: Map{ElementType: Number},
		},
		"unknown-set-with-bound-refinements": {
			hex: "c7070c8301c205020606",
			value: NewValue(Set{ElementType: Bool}, UnknownValue).Refine(refinement.Refinements{
				refinement.KeyNullness:                   refinement.NewNullness(false),
				refinement.KeyCollectionLengthLowerBound: refinement.NewCollectionLengthLowerBound(2),
				refinement.KeyCollectionLengthUpperBound: refinement.NewCollectionLengthUpperBound(6),
			}),
			typ: Set{ElementType: Bool},
		},
		"unknown-with-invalid-refinement-type": {
			hex: "d40000",
			value: NewValue(Bool, UnknownValue).Refine(refinement.Refinements{
				// This refinement will be ignored since only strings will attempt to encode this
				refinement.KeyStringPrefix: refinement.NewStringPrefix("ignored"),
			}),
			typ: Bool,
		},
		"unknown-with-invalid-refinement-data": {
			hex: "d40000",
			value: NewValue(Bool, UnknownValue).Refine(refinement.Refinements{
				// This refinement will be ignored since we don't know how to encode it
				refinement.Key(100): refinement.NewStringPrefix("ignored"),
			}),
			typ: Bool,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := test.value.MarshalMsgPack(test.typ)
			if err != nil {
				t.Fatalf("unexpected error marshaling: %s", err)
			}
			res := hex.EncodeToString(got)
			if res != test.hex {
				t.Errorf("expected msgpack to be %q, got %q", test.hex, res)
			}

			b, err := hex.DecodeString(test.hex)
			if err != nil {
				t.Fatalf("unexpected error parsing hex: %s", err)
			}

			val, err := ValueFromMsgPack(b, test.typ)
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}

			if test.value.String() != val.String() {
				t.Errorf("Unexpected results (-wanted +got): %s", cmp.Diff(test.value, val))
			}
		})
	}
}

func TestMarshalMsgPack_error(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		value         Value
		typ           Type
		expectedError error
	}{
		"unknown-with-invalid-nullness-refinement": {
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				// String prefix is invalid on KeyNullness
				refinement.KeyNullness: refinement.NewStringPrefix("invalid"),
			}),
			typ:           String,
			expectedError: errors.New("error encoding Nullness value refinement: unexpected refinement data of type refinement.StringPrefix"),
		},
		"unknown-with-invalid-prefix-refinement": {
			value: NewValue(String, UnknownValue).Refine(refinement.Refinements{
				// Nullness is invalid on KeyStringPrefix
				refinement.KeyStringPrefix: refinement.NewNullness(false),
			}),
			typ:           String,
			expectedError: errors.New("error encoding StringPrefix value refinement: unexpected refinement data of type refinement.Nullness"),
		},
		"unknown-with-invalid-number-lowerbound-refinement": {
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				// NumberUpperBound is invalid on KeyNumberLowerBound
				refinement.KeyNumberLowerBound: refinement.NewNumberUpperBound(big.NewFloat(1), true),
			}),
			typ:           Number,
			expectedError: errors.New("error encoding NumberLowerBound value refinement: unexpected refinement data of type refinement.NumberUpperBound"),
		},
		"unknown-with-invalid-number-upperbound-refinement": {
			value: NewValue(Number, UnknownValue).Refine(refinement.Refinements{
				// NumberLowerBound is invalid on KeyNumberUpperBound
				refinement.KeyNumberUpperBound: refinement.NewNumberLowerBound(big.NewFloat(1), true),
			}),
			typ:           Number,
			expectedError: errors.New("error encoding NumberUpperBound value refinement: unexpected refinement data of type refinement.NumberLowerBound"),
		},
		"unknown-with-invalid-collection-lowerbound-refinement": {
			value: NewValue(List{ElementType: String}, UnknownValue).Refine(refinement.Refinements{
				// CollectionLengthUpperBound is invalid on KeyCollectionLengthLowerBound
				refinement.KeyCollectionLengthLowerBound: refinement.NewCollectionLengthUpperBound(1),
			}),
			typ:           List{ElementType: String},
			expectedError: errors.New("error encoding CollectionLengthLowerBound value refinement: unexpected refinement data of type refinement.CollectionLengthUpperBound"),
		},
		"unknown-with-invalid-collection-upperbound-refinement": {
			value: NewValue(Map{ElementType: String}, UnknownValue).Refine(refinement.Refinements{
				// CollectionLengthLowerBound is invalid on KeyCollectionLengthUpperBound
				refinement.KeyCollectionLengthUpperBound: refinement.NewCollectionLengthLowerBound(1),
			}),
			typ:           Map{ElementType: String},
			expectedError: errors.New("error encoding CollectionLengthUpperBound value refinement: unexpected refinement data of type refinement.CollectionLengthLowerBound"),
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := test.value.MarshalMsgPack(test.typ)
			if err == nil {
				t.Fatalf("got no error, wanted err: %s", test.expectedError)
			}

			if !strings.Contains(err.Error(), test.expectedError.Error()) {
				t.Fatalf("wanted error %q, got error: %s", test.expectedError.Error(), err.Error())
			}
		})
	}
}
