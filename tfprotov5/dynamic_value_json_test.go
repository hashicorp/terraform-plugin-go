package tfprotov5

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func TestDynamicValueJSON(t *testing.T) {
	t.Parallel()
	type testCase struct {
		value tftypes.Value
		typ   tftypes.Type
		json  string
	}
	tests := map[string]testCase{
		// Primitives
		"string": {
			value: tftypes.NewValue(tftypes.String, "hello"),
			typ:   tftypes.String,
			json:  `"hello"`,
		},
		"string-empty": {
			value: tftypes.NewValue(tftypes.String, ""),
			typ:   tftypes.String,
			json:  `""`,
		},
		"string-from-number": {
			value: tftypes.NewValue(tftypes.String, "15"),
			typ:   tftypes.String,
			json:  `15`,
		},
		"string-from-bool": {
			value: tftypes.NewValue(tftypes.String, "true"),
			typ:   tftypes.String,
			json:  "true",
		},
		"string-null": {
			value: tftypes.NewValue(tftypes.String, nil),
			typ:   tftypes.String,
			json:  "null",
		},
		"number-from-int": {
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(2)),
			typ:   tftypes.Number,
			json:  `2`,
		},
		"number-from-float": {
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(2.5)),
			typ:   tftypes.Number,
			json:  `2.5`,
		},
		"number-from-string": {
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(5)),
			typ:   tftypes.Number,
			json:  `"5"`,
		},
		"bool-true": {
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.Bool,
			json:  `true`,
		},
		"bool-false": {
			value: tftypes.NewValue(tftypes.Bool, false),
			typ:   tftypes.Bool,
			json:  `false`,
		},
		"bool-from-string": {
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.Bool,
			json:  `"true"`,
		},
		"bool-from-string_int": {
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.Bool,
			json:  `"1"`,
		},
		"bool-from-int": {
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.Bool,
			json:  `1`,
		},

		// Lists
		"list-of-bools": {
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.Bool, false),
			}),
			typ: tftypes.List{
				ElementType: tftypes.Bool,
			},
			json: `[true,false]`,
		},
		"list-empty": {
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{}),
			typ: tftypes.List{
				ElementType: tftypes.Bool,
			},
			json: `[]`,
		},
		"list-of-bools-from-strings": {
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.Bool, false),
			}),
			typ: tftypes.List{
				ElementType: tftypes.Bool,
			},
			json: `["true","false"]`,
		},

		// Sets
		"set-of-bools": {
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, false),
				tftypes.NewValue(tftypes.Bool, true),
			}),
			typ: tftypes.Set{
				ElementType: tftypes.Bool,
			},
			json: `[false,true]`,
		},
		"set-empty": {
			value: tftypes.NewValue(tftypes.Set{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{}),
			typ: tftypes.Set{
				ElementType: tftypes.Bool,
			},
			json: `[]`,
		},

		// Tuples
		"tuple-of-bool_number": {
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{
					tftypes.Bool,
					tftypes.Number,
				},
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.Number, big.NewFloat(5)),
			}),
			typ: tftypes.Tuple{
				ElementTypes: []tftypes.Type{
					tftypes.Bool,
					tftypes.Number,
				},
			},
			json: `[true,5]`,
		},
		"tuple-empty": {
			value: tftypes.NewValue(tftypes.Tuple{
				ElementTypes: []tftypes.Type{},
			}, []tftypes.Value{}),
			typ: tftypes.Tuple{
				ElementTypes: []tftypes.Type{},
			},
			json: `[]`,
		},

		// Maps
		"map-empty": {
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.Bool,
			}, map[string]tftypes.Value{}),
			typ: tftypes.Map{
				AttributeType: tftypes.Bool,
			},
			json: `{}`,
		},
		"map-of-bools": {
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.Bool,
			}, map[string]tftypes.Value{
				"yes": tftypes.NewValue(tftypes.Bool, true),
				"no":  tftypes.NewValue(tftypes.Bool, false),
			}),
			typ: tftypes.Map{
				AttributeType: tftypes.Bool,
			},
			json: `{"no":false,"yes":true}`,
		},
		"map-null": {
			value: tftypes.NewValue(tftypes.Map{
				AttributeType: tftypes.Bool,
			}, nil),
			typ: tftypes.Map{
				AttributeType: tftypes.Bool,
			},
			json: `null`,
		},

		// Objects
		"object-empty": {
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			}, map[string]tftypes.Value{}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
			json: `{}`,
		},
		"object-of-bool_number": {
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			}, map[string]tftypes.Value{
				"bool":   tftypes.NewValue(tftypes.Bool, true),
				"number": tftypes.NewValue(tftypes.Number, big.NewFloat(0)),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"bool":   tftypes.Bool,
					"number": tftypes.Number,
				},
			},
			json: `{"bool":true,"number":0}`,
		},

		// Encoding into dynamic produces type information wrapper
		"dynamic-bool": {
			value: tftypes.NewValue(tftypes.Bool, true),
			typ:   tftypes.DynamicPseudoType,
			json:  `{"value":true,"type":"bool"}`,
		},
		"dynamic-string": {
			value: tftypes.NewValue(tftypes.String, "hello"),
			typ:   tftypes.DynamicPseudoType,
			json:  `{"value":"hello","type":"string"}`,
		},
		"dynamic-number": {
			value: tftypes.NewValue(tftypes.Number, big.NewFloat(5)),
			typ:   tftypes.DynamicPseudoType,
			json:  `{"value":5,"type":"number"}`,
		},
		"dynamic-list-of-bools": {
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.Bool, false),
			}),
			typ:  tftypes.DynamicPseudoType,
			json: `{"value":[true,false],"type":["list","bool"]}`,
		},
		"list-of-dynamic-bools": {
			value: tftypes.NewValue(tftypes.List{
				ElementType: tftypes.Bool,
			}, []tftypes.Value{
				tftypes.NewValue(tftypes.Bool, true),
				tftypes.NewValue(tftypes.Bool, false),
			}),
			typ: tftypes.List{
				ElementType: tftypes.DynamicPseudoType,
			},
			json: `[{"value":true,"type":"bool"},{"value":false,"type":"bool"}]`,
		},
		"object-of-bool_dynamic": {
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"static":  tftypes.Bool,
					"dynamic": tftypes.DynamicPseudoType,
				},
			}, map[string]tftypes.Value{
				"static":  tftypes.NewValue(tftypes.Bool, true),
				"dynamic": tftypes.NewValue(tftypes.Bool, true),
			}),
			typ: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"static":  tftypes.Bool,
					"dynamic": tftypes.DynamicPseudoType,
				},
			},
			json: `{"dynamic":{"value":true,"type":"bool"},"static":true}`,
		},
		"dynamic-object": {
			value: tftypes.NewValue(tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"static":  tftypes.Bool,
					"dynamic": tftypes.Bool,
				},
			}, map[string]tftypes.Value{
				"static":  tftypes.NewValue(tftypes.Bool, true),
				"dynamic": tftypes.NewValue(tftypes.Bool, true),
			}),
			typ:  tftypes.DynamicPseudoType,
			json: `{"value":{"dynamic":true,"static":true},"type":["object",{"dynamic":"bool","static":"bool"}]}`,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dv := DynamicValue{
				JSON: []byte(test.json),
			}
			val, err := dv.Unmarshal(test.typ)
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}
			if diff := cmp.Diff(test.value, val, cmp.Comparer(numberComparer), tftypes.ValueComparer()); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
