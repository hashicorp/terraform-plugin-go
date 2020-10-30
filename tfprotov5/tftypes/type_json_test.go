package tftypes

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nsf/jsondiff"
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
			typ:  Map{AttributeType: String},
		},
		"object-string_number_bool": {
			json: `["object",{"foo":"string","bar":"number","baz":"bool"}]`,
			typ: Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}},
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
	jsondiffopts := jsondiff.DefaultConsoleOptions()
	for name, test := range testCases {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			typ, err := ParseJSONType([]byte(test.json))
			if err != nil {
				t.Fatalf("unexpected error parsing JSON: %s", err)
			}
			if !typ.Is(test.typ) {
				t.Fatalf("Unexpected parsing results (-wanted +got): %s", cmp.Diff(test.typ, typ))
			}

			json, err := typ.MarshalJSON()
			if err != nil {
				t.Fatalf("unexpected error generating JSON: %s", err)
			}
			diff, diffStr := jsondiff.Compare([]byte(test.json), json, &jsondiffopts)
			if diff != jsondiff.FullMatch {
				t.Fatalf("unexpected JSON generating results (got => expected): %s", diffStr)
			}
		})
	}
}
