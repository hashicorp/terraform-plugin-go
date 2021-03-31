package tfprotov6

import (
	"encoding/hex"
	"math"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestDynamicValueMsgPack(t *testing.T) {
	t.Parallel()
	type testCase struct {
		hex   string
		value tftypes.Value
		typ   tftypes.Type
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
			value: tftypes.NewValue(tftypes.String, "hello"),
			typ:   tftypes.String,
		},
		"empty-string": {
			hex:   "a0",
			value: tftypes.NewValue(tftypes.String, ""),
			typ:   tftypes.String,
		},
		"null-string": {
			hex:   "c0",
			value: tftypes.NewValue(tftypes.String, nil),
			typ:   tftypes.String,
		},
		"unknown-string": {
			hex:   "d40000",
			value: tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			typ:   tftypes.String,
		},
		"true-bool": {
			hex:   "c3",
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.Bool,
		},
		"false-bool": {
			hex:   "c2",
			value: tftypes.NewValue(tftypes.Bool, false),
			typ:   tftypes.Bool,
		},
		"null-bool": {
			hex:   "c0",
			value: tftypes.NewValue(tftypes.Bool, nil),
			typ:   tftypes.Bool,
		},
		"unknown-bool": {
			hex:   "d40000",
			value: tftypes.NewValue(tftypes.Bool, tftypes.UnknownValue),
			typ:   tftypes.Bool,
		},
		"int-number": {
			hex:   "01",
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(1)),
			typ:   tftypes.Number,
		},
		"float-number": {
			hex:   "cb3ff8000000000000",
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(1.5)),
			typ:   tftypes.Number,
		},
		"big-number": {
			hex:   "d96439393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939393939",
			value: tftypes.NewValue(tftypes.Number, bigNumber),
			typ:   tftypes.Number,
		},
		"awkward-fraction-number": {
			hex:   "a3302e38",
			value: tftypes.NewValue(tftypes.Number, awkwardFraction),
			typ:   tftypes.Number,
		},
		"positive-infinity-number": {
			hex:   "cb7ff0000000000000",
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(math.Inf(1))),
			typ:   tftypes.Number,
		},
		"negative-infinity-number": {
			hex:   "cbfff0000000000000",
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(math.Inf(-1))),
			typ:   tftypes.Number,
		},
		"list-string-hello": {
			hex: "91a568656c6c6f",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.List{ElementType: tftypes.String},
		},
		"list-string-unknown": {
			hex: "91d40000",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.List{ElementType: tftypes.String},
		},
		"list-string-null-string": {
			hex: "91c0",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.List{ElementType: tftypes.String},
		},
		"list-string-null": {
			hex: "c0",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, nil),
			typ: tftypes.List{ElementType: tftypes.String},
		},
		"list-string-empty": {
			hex: "90",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
			typ: tftypes.List{ElementType: tftypes.String},
		},
		"set-string-hello": {
			hex: "91a568656c6c6f",
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.Set{ElementType: tftypes.String},
		},
		"set-string-unknown": {
			hex: "91d40000",
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.Set{ElementType: tftypes.String},
		},
		"set-string-null-string": {
			hex: "91c0",
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.Set{ElementType: tftypes.String},
		},
		"set-string-empty": {
			hex: "90",
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.String,
			}, []tftypes.Value{}),
			typ: tftypes.Set{ElementType: tftypes.String},
		},
		"map-string-hello": {
			hex: "81a86772656574696e67a568656c6c6f",
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.Map{AttributeType: tftypes.String},
		},
		"map-string-unknown": {
			hex: "81a86772656574696e67d40000",
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.Map{AttributeType: tftypes.String},
		},
		"map-string-null": {
			hex: "81a86772656574696e67c0",
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.Map{AttributeType: tftypes.String},
		},
		"map-string-empty": {
			hex: "80",
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.String,
			}, map[string]tftypes.Value{}),
			typ: tftypes.Map{AttributeType: tftypes.String},
		},
		"tuple-string-hello": {
			hex: "91a568656c6c6f",
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{tftypes.String},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.Tuple{ElementTypes: []tftypes.Type{tftypes.String}},
		},
		"tuple-string-unknown": {
			hex: "91d40000",
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{tftypes.String},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.Tuple{ElementTypes: []tftypes.Type{tftypes.String}},
		},
		"tuple-string-null": {
			hex: "91c0",
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{tftypes.String},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.Tuple{ElementTypes: []tftypes.Type{tftypes.String}},
		},
		"tuple-empty": {
			hex: "90",
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{},
			}, []tftypes.Value{}),
			typ: tftypes.Tuple{ElementTypes: []tftypes.Type{}},
		},
		"object-string-hello": {
			hex: "81a86772656574696e67a568656c6c6f",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			},
		},
		"object-string-unknown": {
			hex: "81a86772656574696e67d40000",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			},
		},
		"object-string-null": {
			hex: "81a86772656574696e67c0",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"greeting": tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"greeting": tftypes.String,
				},
			},
		},
		"object-string-multi-null": {
			hex: "82a161c0a162c0",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
					"b": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.String, nil),
				"b": tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
					"b": tftypes.String,
				},
			},
		},
		"object-string-multi-unknown": {
			hex: "82a161d40000a162d40000",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
					"b": tftypes.String,
				},
			}, map[string]tftypes.Value{
				"a": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
				"b": tftypes.NewValue(tftypes.String, tftypes.UnknownValue),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"a": tftypes.String,
					"b": tftypes.String,
				},
			},
		},
		"object-empty": {
			hex: "80",
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			}, map[string]tftypes.Value{}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"string-dynamic-null": {
			hex:   "92c40822737472696e6722c0",
			value: tftypes.NewValue(tftypes.String, nil),
			typ:   tftypes.DynamicPseudoType,
		},
		"dynamic-unknown": {
			hex:   "d40000",
			value: tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
			typ:   tftypes.DynamicPseudoType,
		},
		"dynamic-list-string-hello": {
			hex: "9192c40822737472696e6722a568656c6c6f",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, "hello"),
			}),
			typ: tftypes.List{ElementType: tftypes.DynamicPseudoType},
		},
		"dynamic-list-string-null": {
			hex: "9192c40822737472696e6722c0",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.String,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.String, nil),
			}),
			typ: tftypes.List{ElementType: tftypes.DynamicPseudoType},
		},
		"dynamic-list-unknown": {
			hex: "91d40000",
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.DynamicPseudoType,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.DynamicPseudoType, tftypes.UnknownValue),
			}),
			typ: tftypes.List{ElementType: tftypes.DynamicPseudoType},
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
			dv := DynamicValue{
				MsgPack: b,
			}
			val, err := dv.Unmarshal(test.typ)
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}

			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
