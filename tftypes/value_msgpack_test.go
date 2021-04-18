package tftypes

import (
	"encoding/hex"
	"math"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

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
		"float-number": {
			hex:   "cb3ff8000000000000",
			value: NewValue(Number, big.NewFloat(1.5)),
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
		"map-string-hello": {
			hex: "81a86772656574696e67a568656c6c6f",
			value: NewValue(Map{
				AttributeType: String,
			}, map[string]Value{
				"greeting": NewValue(String, "hello"),
			}),
			typ: Map{AttributeType: String},
		},
		"map-string-unknown": {
			hex: "81a86772656574696e67d40000",
			value: NewValue(Map{
				AttributeType: String,
			}, map[string]Value{
				"greeting": NewValue(String, UnknownValue),
			}),
			typ: Map{AttributeType: String},
		},
		"map-string-null": {
			hex: "81a86772656574696e67c0",
			value: NewValue(Map{
				AttributeType: String,
			}, map[string]Value{
				"greeting": NewValue(String, nil),
			}),
			typ: Map{AttributeType: String},
		},
		"map-string-empty": {
			hex: "80",
			value: NewValue(Map{
				AttributeType: String,
			}, map[string]Value{}),
			typ: Map{AttributeType: String},
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
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := test.value.MarshalMsgPack(test.typ) //nolint:staticcheck
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

			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
