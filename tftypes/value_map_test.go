package tftypes

import (
	"regexp"
	"testing"
)

func Test_newValue_map(t *testing.T) {
	t.Parallel()
	type testCase struct {
		typ      Type
		val      interface{}
		err      *regexp.Regexp
		expected Value
	}
	tests := map[string]testCase{
		"normal": {
			typ: Map{ElementType: String},
			val: map[string]Value{
				"a": NewValue(String, "hello"),
				"b": NewValue(String, "world"),
			},
			expected: Value{
				typ: Map{ElementType: String},
				value: map[string]Value{
					"a": {
						typ:   String,
						value: "hello",
					},
					"b": {
						typ:   String,
						value: "world",
					},
				},
			},
			err: nil,
		},
		"dynamic": {
			typ: Map{ElementType: DynamicPseudoType},
			val: map[string]Value{
				"a": NewValue(String, "hello"),
				"b": NewValue(String, "world"),
			},
			expected: Value{
				typ: Map{ElementType: DynamicPseudoType},
				value: map[string]Value{
					"a": {
						typ:   String,
						value: "hello",
					},
					"b": {
						typ:   String,
						value: "world",
					},
				},
			},
			err: nil,
		},
		"dynamic-different-types": {
			typ: Map{ElementType: DynamicPseudoType},
			val: map[string]Value{
				"a": NewValue(String, "hello"),
				"b": NewValue(Number, 123),
			},
			err: regexp.MustCompile(`maps must only contain one type of element, saw tftypes.String and tftypes.Number`),
		},
		"wrong-type": {
			typ: Map{ElementType: String},
			val: map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(Number, 123),
			},
			err: regexp.MustCompile(`ElementKeyString\("b"\): can't use tftypes.Number as tftypes.String`),
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
