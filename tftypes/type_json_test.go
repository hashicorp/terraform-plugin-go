package tftypes

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// Reference: https://github.com/google/go-cmp/issues/224
var cmpTransformJSON = cmp.FilterValues(
	func(x, y []byte) bool {
		return json.Valid(x) && json.Valid(y)
	},
	cmp.Transformer("ParseJSON", func(in []byte) (out interface{}) {
		if err := json.Unmarshal(in, &out); err != nil {
			panic(err) // should never occur given previous filter to ensure valid JSON
		}

		return out
	}),
)

func TestTypeJSON(t *testing.T) {
	t.Parallel()
	type testCase struct {
		json string
		typ  Type
	}

	testCases := map[string]testCase{
		"string": {
			json: `"string"`,
			typ:  String,
		},
		"number": {
			json: `"number"`,
			typ:  Number,
		},
		"bool": {
			json: `"bool"`,
			typ:  Bool,
		},
		"dynamic": {
			json: `"dynamic"`,
			typ:  DynamicPseudoType,
		},
		"list-string": {
			json: `["list","string"]`,
			typ:  List{ElementType: String},
		},
		"set-string": {
			json: `["set","string"]`,
			typ:  Set{ElementType: String},
		},
		"map-string": {
			json: `["map","string"]`,
			typ:  Map{ElementType: String},
		},
		"object-string_number_bool": {
			json: `["object",{"foo":"string","bar":"number","baz":"bool"}]`,
			typ: Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}},
		},
		"object-string_number_bool-optional_string": {
			json: `["object",{"foo":"string","bar":"number","baz":"bool"},["foo"]]`,
			typ: Object{
				AttributeTypes: map[string]Type{
					"foo": String,
					"bar": Number,
					"baz": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"foo": {},
				},
			},
		},
		"object-string_number_bool-optional_string-number": {
			json: `["object",{"foo":"string","bar":"number","baz":"bool"},["bar", "foo"]]`,
			typ: Object{
				AttributeTypes: map[string]Type{
					"foo": String,
					"bar": Number,
					"baz": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"foo": {},
					"bar": {},
				},
			},
		},
		"object-string_number_bool-optional_string-number-bool": {
			json: `["object",{"foo":"string","bar":"number","baz":"bool"},["bar","baz","foo"]]`,
			typ: Object{
				AttributeTypes: map[string]Type{
					"foo": String,
					"bar": Number,
					"baz": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"foo": {},
					"bar": {},
					"baz": {},
				},
			},
		},
		"tuple-string_number_bool": {
			json: `["tuple",["string","number","bool"]]`,
			typ: Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}},
		},
		"object-very_complicated": {
			json: `["object",{
				"foo":["tuple",[
					["object",{"a":"number","b":"string","c":"bool"}],
					"string",
					"number",
					["list",[
						"tuple", ["string","bool"]
					]]
				]],
				"bar":["set",["object",{
					"red":"string","blue":["list","string"],"green":"number"
				}]]
			}]`,
			typ: Object{AttributeTypes: map[string]Type{
				"foo": Tuple{ElementTypes: []Type{
					Object{AttributeTypes: map[string]Type{
						"a": Number,
						"b": String,
						"c": Bool,
					}},
					String,
					Number,
					List{ElementType: Tuple{ElementTypes: []Type{
						String, Bool,
					}}},
				}},
				"bar": Set{ElementType: Object{
					AttributeTypes: map[string]Type{
						"red": String,
						"blue": List{
							ElementType: String,
						},
						"green": Number,
					},
				}},
			}},
		},
	}

	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			typ, err := ParseJSONType([]byte(test.json))
			if err != nil {
				t.Fatalf("unexpected error parsing JSON: %s", err)
			}
			if !typ.Equal(test.typ) {
				t.Fatalf("Unexpected parsing results (-wanted +got): %s", cmp.Diff(test.typ, typ))
			}

			typJSON, err := typ.MarshalJSON()
			if err != nil {
				t.Fatalf("unexpected error generating JSON: %s", err)
			}

			if diff := cmp.Diff([]byte(test.json), typJSON, cmpTransformJSON); diff != "" {
				t.Fatalf("unexpected generated JSON difference: %s", diff)
			}
		})
	}
}
