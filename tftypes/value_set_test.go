package tftypes

import (
	"regexp"
	"testing"
)

func Test_newValue_set(t *testing.T) {
	t.Parallel()
	type testCase struct {
		typ      Type
		val      interface{}
		err      *regexp.Regexp
		expected Value
	}
	tests := map[string]testCase{
		"normal": {
			typ: Set{ElementType: String},
			val: []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			},
			expected: Value{
				typ: Set{ElementType: String},
				value: []Value{
					{
						typ:   String,
						value: "hello",
					},
					{
						typ:   String,
						value: "world",
					},
				},
			},
			err: nil,
		},
		"dynamic": {
			typ: Set{ElementType: DynamicPseudoType},
			val: []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			},
			expected: Value{
				typ: Set{ElementType: DynamicPseudoType},
				value: []Value{
					{
						typ:   String,
						value: "hello",
					},
					{
						typ:   String,
						value: "world",
					},
				},
			},
			err: nil,
		},
		"dynamic-different-types": {
			typ: Set{ElementType: DynamicPseudoType},
			val: []Value{
				NewValue(String, "hello"),
				NewValue(Number, 123),
			},
			err: regexp.MustCompile(`sets must only contain one type of element, saw tftypes.String and tftypes.Number`),
		},
		"wrong-type": {
			typ: Set{ElementType: String},
			val: []Value{
				NewValue(String, "foo"),
				NewValue(Number, 123),
			},
			err: regexp.MustCompile(`ElementKeyValue\(tftypes.Number<\"123\">\): can't use tftypes.Number as tftypes.String`),
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
