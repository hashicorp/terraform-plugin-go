package tftypes

import (
	"math/big"
	"regexp"
	"testing"
)

func Test_newValue_object(t *testing.T) {
	t.Parallel()
	type testCase struct {
		typ      Type
		val      interface{}
		err      *regexp.Regexp
		expected Value
	}
	tests := map[string]testCase{
		"normal": {
			typ: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			val: map[string]Value{
				"a": NewValue(String, "hello"),
				"b": NewValue(Number, 123),
				"c": NewValue(Bool, true),
			},
			expected: Value{
				typ: Object{AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				}},
				value: map[string]Value{
					"a": {
						typ:   String,
						value: "hello",
					},
					"b": {
						typ:   Number,
						value: big.NewFloat(123),
					},
					"c": {
						typ:   Bool,
						value: true,
					},
				},
			},
			err: nil,
		},
		"optional-included": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			val: map[string]Value{
				"a": NewValue(String, "hello"),
				"b": NewValue(Number, 123),
				"c": NewValue(Bool, true),
			},
			expected: Value{
				typ: Object{
					AttributeTypes: map[string]Type{
						"a": String,
						"b": Number,
						"c": Bool,
					},
					OptionalAttributes: map[string]struct{}{
						"a": {},
					},
				},
				value: map[string]Value{
					"a": {
						typ:   String,
						value: "hello",
					},
					"b": {
						typ:   Number,
						value: big.NewFloat(123),
					},
					"c": {
						typ:   Bool,
						value: true,
					},
				},
			},
			err: nil,
		},
		"optional-excluded": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			val: map[string]Value{
				"b": NewValue(Number, 123),
				"c": NewValue(Bool, true),
			},
			expected: Value{
				typ: Object{
					AttributeTypes: map[string]Type{
						"a": String,
						"b": Number,
						"c": Bool,
					},
					OptionalAttributes: map[string]struct{}{
						"a": {},
					},
				},
				value: map[string]Value{
					"b": {
						typ:   Number,
						value: big.NewFloat(123),
					},
					"c": {
						typ:   Bool,
						value: true,
					},
				},
			},
			err: nil,
		},
		"missing-attribute": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
			},
			val: map[string]Value{
				"b": NewValue(Number, 123),
				"c": NewValue(Bool, true),
			},
			err: regexp.MustCompile(`required attribute "a" not set`),
		},
		"invalid-attribute": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
			},
			val: map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(Number, 123),
				"c": NewValue(Bool, false),
				"d": NewValue(Bool, true),
			},
			err: regexp.MustCompile(`can't set a value on "d" in tftypes.NewValue, key not part of the object type`),
		},
		"attribute-wrong-type": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
			},
			val: map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(Number, 123),
				"c": NewValue(String, "false"),
			},
			err: regexp.MustCompile(`AttributeName\("c"\): can't use tftypes.String as tftypes.Bool`),
		},
		"dynamic-attribute": {
			typ: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": DynamicPseudoType,
				},
			},
			val: map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(Number, 123),
				"c": NewValue(String, "false"),
			},
			expected: Value{
				typ: Object{AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": DynamicPseudoType,
				}},
				value: map[string]Value{
					"a": {
						typ:   String,
						value: "foo",
					},
					"b": {
						typ:   Number,
						value: big.NewFloat(123),
					},
					"c": {
						typ:   String,
						value: "false",
					},
				},
			},
			err: nil,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res, err := newValue(test.typ, test.val)
			if err == nil && test.err != nil {
				t.Errorf("Expected error to match %q, got nil", test.err)
			} else if err != nil && test.err == nil {
				t.Errorf("Expected error to be nil, got %q", err)
			} else if err != nil && test.err != nil && !test.err.MatchString(err.Error()) {
				t.Errorf("Expected error to match %q, got %q", test.err, err.Error())
			}
			if !res.Equal(test.expected) {
				t.Errorf("Expected value to be %s, got %s", test.expected, res)
			}
		})
	}
}
