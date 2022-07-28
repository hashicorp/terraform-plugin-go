package tftypes

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestValueFromJSON(t *testing.T) {
	t.Parallel()
	type testCase struct {
		value         Value
		typ           Type
		json          string
		expectedError error
	}
	tests := map[string]testCase{
		// Primitives
		"string": {
			value: NewValue(String, "hello"),
			typ:   String,
			json:  `"hello"`,
		},
		"string-empty": {
			value: NewValue(String, ""),
			typ:   String,
			json:  `""`,
		},
		"string-from-number": {
			value: NewValue(String, "15"),
			typ:   String,
			json:  `15`,
		},
		"string-from-bool": {
			value: NewValue(String, "true"),
			typ:   String,
			json:  "true",
		},
		"string-null": {
			value: NewValue(String, nil),
			typ:   String,
			json:  "null",
		},
		"number-from-int": {
			value: NewValue(Number, big.NewFloat(2)),
			typ:   Number,
			json:  `2`,
		},
		"number-from-float": {
			value: NewValue(Number, big.NewFloat(2.5)),
			typ:   Number,
			json:  `2.5`,
		},
		"number-from-string": {
			value: NewValue(Number, big.NewFloat(5)),
			typ:   Number,
			json:  `"5"`,
		},
		"bool-true": {
			value: NewValue(Bool, true),
			typ:   Bool,
			json:  `true`,
		},
		"bool-false": {
			value: NewValue(Bool, false),
			typ:   Bool,
			json:  `false`,
		},
		"bool-from-string": {
			value: NewValue(Bool, true),
			typ:   Bool,
			json:  `"true"`,
		},
		"bool-from-string_int": {
			value: NewValue(Bool, true),
			typ:   Bool,
			json:  `"1"`,
		},
		"bool-from-int": {
			value: NewValue(Bool, true),
			typ:   Bool,
			json:  `1`,
		},

		// Lists
		"list-of-bools": {
			value: NewValue(List{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ: List{
				ElementType: Bool,
			},
			json: `[true,false]`,
		},
		"list-empty": {
			value: NewValue(List{
				ElementType: Bool,
			}, []Value{}),
			typ: List{
				ElementType: Bool,
			},
			json: `[]`,
		},
		"list-of-bools-from-strings": {
			value: NewValue(List{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ: List{
				ElementType: Bool,
			},
			json: `["true","false"]`,
		},

		// Sets
		"set-of-bools": {
			value: NewValue(Set{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, false),
				NewValue(Bool, true),
			}),
			typ: Set{
				ElementType: Bool,
			},
			json: `[false,true]`,
		},
		"set-empty": {
			value: NewValue(Set{
				ElementType: Bool,
			}, []Value{}),
			typ: Set{
				ElementType: Bool,
			},
			json: `[]`,
		},

		// Tuples
		"tuple-of-bool_number": {
			value: NewValue(Tuple{
				ElementTypes: []Type{
					Bool,
					Number,
				},
			}, []Value{
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(5)),
			}),
			typ: Tuple{
				ElementTypes: []Type{
					Bool,
					Number,
				},
			},
			json: `[true,5]`,
		},
		"tuple-empty": {
			value: NewValue(Tuple{
				ElementTypes: []Type{},
			}, []Value{}),
			typ: Tuple{
				ElementTypes: []Type{},
			},
			json: `[]`,
		},

		// Maps
		"map-empty": {
			value: NewValue(Map{
				ElementType: Bool,
			}, map[string]Value{}),
			typ: Map{
				ElementType: Bool,
			},
			json: `{}`,
		},
		"map-of-bools": {
			value: NewValue(Map{
				ElementType: Bool,
			}, map[string]Value{
				"yes": NewValue(Bool, true),
				"no":  NewValue(Bool, false),
			}),
			typ: Map{
				ElementType: Bool,
			},
			json: `{"no":false,"yes":true}`,
		},
		"map-null": {
			value: NewValue(Map{
				ElementType: Bool,
			}, nil),
			typ: Map{
				ElementType: Bool,
			},
			json: `null`,
		},

		// Objects
		"object-empty": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{},
			}, map[string]Value{}),
			typ: Object{
				AttributeTypes: map[string]Type{},
			},
			json: `{}`,
		},
		"object-attribute-key-token-error": {
			value: Value{},
			typ: Object{
				AttributeTypes: map[string]Type{},
			},
			json: `{{}}`,
			expectedError: AttributePathError{
				Path: NewAttributePath(),
				err:  fmt.Errorf("error reading object attribute key token: invalid character '{'"),
			},
		},
		"object-attribute-key-missing-error": {
			value: Value{},
			typ: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			json: `{"not-test": "test-value"}`,
			expectedError: AttributePathError{
				Path: NewAttributePath().WithAttributeName("not-test"),
				err:  fmt.Errorf("unsupported attribute \"not-test\""),
			},
		},
		"object-of-bool_number": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			}, map[string]Value{
				"bool":   NewValue(Bool, true),
				"number": NewValue(Number, big.NewFloat(0)),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			},
			json: `{"bool":true,"number":0}`,
		},

		// Encoding into dynamic produces type information wrapper
		"dynamic-bool": {
			value: NewValue(Bool, true),
			typ:   DynamicPseudoType,
			json:  `{"value":true,"type":"bool"}`,
		},
		"dynamic-string": {
			value: NewValue(String, "hello"),
			typ:   DynamicPseudoType,
			json:  `{"value":"hello","type":"string"}`,
		},
		"dynamic-number": {
			value: NewValue(Number, big.NewFloat(5)),
			typ:   DynamicPseudoType,
			json:  `{"value":5,"type":"number"}`,
		},
		"dynamic-list-of-bools": {
			value: NewValue(List{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ:  DynamicPseudoType,
			json: `{"value":[true,false],"type":["list","bool"]}`,
		},
		"list-of-dynamic-bools": {
			value: NewValue(List{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ: List{
				ElementType: DynamicPseudoType,
			},
			json: `[{"value":true,"type":"bool"},{"value":false,"type":"bool"}]`,
		},
		"dynamic-set-of-bools": {
			value: NewValue(Set{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ:  DynamicPseudoType,
			json: `{"value":[true,false],"type":["set","bool"]}`,
		},
		"set-of-dynamic-bools": {
			value: NewValue(Set{
				ElementType: Bool,
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ: Set{
				ElementType: DynamicPseudoType,
			},
			json: `[{"value":true,"type":"bool"},{"value":false,"type":"bool"}]`,
		},
		"dynamic-tuple-of-bools": {
			value: NewValue(Tuple{
				ElementTypes: []Type{Bool, Bool},
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ:  DynamicPseudoType,
			json: `{"value":[true,false],"type":["tuple",["bool","bool"]]}`,
		},
		"tuple-of-dynamic-bools": {
			value: NewValue(Tuple{
				ElementTypes: []Type{DynamicPseudoType, DynamicPseudoType},
			}, []Value{
				NewValue(Bool, true),
				NewValue(Bool, false),
			}),
			typ: Tuple{
				ElementTypes: []Type{DynamicPseudoType, DynamicPseudoType},
			},
			json: `[{"value":true,"type":"bool"},{"value":false,"type":"bool"}]`,
		},
		"dynamic-map-of-bools": {
			value: NewValue(Map{
				ElementType: Bool,
			}, map[string]Value{
				"true":  NewValue(Bool, true),
				"false": NewValue(Bool, false),
			}),
			typ:  DynamicPseudoType,
			json: `{"value":{"true":true,"false":false},"type":["map","bool"]}`,
		},
		"map-of-dynamic-bools": {
			value: NewValue(Map{
				ElementType: DynamicPseudoType,
			}, map[string]Value{
				"true":  NewValue(Bool, true),
				"false": NewValue(Bool, false),
			}),
			typ: Map{
				ElementType: DynamicPseudoType,
			},
			json: `{"true":{"value":true,"type":"bool"},"false":{"value":false,"type":"bool"}}`,
		},
		"object-of-bool_dynamic": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"static":  Bool,
					"dynamic": DynamicPseudoType,
				},
			}, map[string]Value{
				"static":  NewValue(Bool, true),
				"dynamic": NewValue(Bool, true),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"static":  Bool,
					"dynamic": DynamicPseudoType,
				},
			},
			json: `{"dynamic":{"value":true,"type":"bool"},"static":true}`,
		},
		"dynamic-object": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"static":  Bool,
					"dynamic": Bool,
				},
			}, map[string]Value{
				"static":  NewValue(Bool, true),
				"dynamic": NewValue(Bool, true),
			}),
			typ:  DynamicPseudoType,
			json: `{"value":{"dynamic":true,"static":true},"type":["object",{"dynamic":"bool","static":"bool"}]}`,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			val, err := ValueFromJSON([]byte(test.json), test.typ)
			if diff := cmp.Diff(test.expectedError, err); diff != "" {
				t.Errorf("unexpected error difference: %s", diff)
			}
			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}

func TestValueFromJSONWithOpts(t *testing.T) {
	t.Parallel()
	type testCase struct {
		value Value
		typ   Type
		json  string
	}
	tests := map[string]testCase{
		"object-of-bool-number": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			}, map[string]Value{
				"bool":   NewValue(Bool, true),
				"number": NewValue(Number, big.NewFloat(0)),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			},
			json: `{"bool":true,"number":0}`,
		},
		"object-with-missing-attribute": {
			value: NewValue(Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			}, map[string]Value{
				"bool":   NewValue(Bool, true),
				"number": NewValue(Number, big.NewFloat(0)),
			}),
			typ: Object{
				AttributeTypes: map[string]Type{
					"bool":   Bool,
					"number": Number,
				},
			},
			json: `{"bool":true,"number":0,"unknown":"whatever"}`,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			val, err := ValueFromJSONWithOpts([]byte(test.json), test.typ, ValueFromJSONOpts{
				IgnoreUndefinedAttributes: true,
			})
			if err != nil {
				t.Fatalf("unexpected error unmarshaling: %s", err)
			}
			if diff := cmp.Diff(test.value, val); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
