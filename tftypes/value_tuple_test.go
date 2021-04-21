package tftypes

import (
	"math/big"
	"regexp"
	"testing"
)

func Test_newValue_tuple(t *testing.T) {
	t.Parallel()
	type testCase struct {
		typ      Type
		val      interface{}
		err      *regexp.Regexp
		expected Value
	}
	tests := map[string]testCase{
		"normal": {
			typ: Tuple{ElementTypes: []Type{
				String,
				Number,
				Bool,
			}},
			val: []Value{
				NewValue(String, "hello"),
				NewValue(Number, 123),
				NewValue(Bool, true),
			},
			expected: Value{
				typ: Tuple{ElementTypes: []Type{
					String,
					Number,
					Bool,
				}},
				value: []Value{
					{
						typ:   String,
						value: "hello",
					},
					{
						typ:   Number,
						value: big.NewFloat(123),
					},
					{
						typ:   Bool,
						value: true,
					},
				},
			},
			err: nil,
		},
		"missing-element": {
			typ: Tuple{
				ElementTypes: []Type{
					String,
					Number,
					Bool,
				},
			},
			val: []Value{
				NewValue(String, "hello"),
				NewValue(Number, 123),
			},
			err: regexp.MustCompile(`can't create a tftypes.Value with 2 elements, type tftypes.Tuple\[tftypes.String, tftypes.Number, tftypes.Bool\] requires 3 elements`),
		},
		"extra-element": {
			typ: Tuple{
				ElementTypes: []Type{
					String,
					Number,
					Bool,
				},
			},
			val: []Value{
				NewValue(String, "foo"),
				NewValue(Number, 123),
				NewValue(Bool, false),
				NewValue(Bool, true),
			},
			err: regexp.MustCompile(`can't create a tftypes.Value with 4 elements, type tftypes.Tuple\[tftypes.String, tftypes.Number, tftypes.Bool\] requires 3 elements`),
		},
		"element-wrong-type": {
			typ: Tuple{
				ElementTypes: []Type{
					String,
					Number,
					Bool,
				},
			},
			val: []Value{
				NewValue(String, "foo"),
				NewValue(Number, 123),
				NewValue(String, "false"),
			},
			err: regexp.MustCompile(`ElementKeyInt\(2\): can't use tftypes.String as tftypes.Bool`),
		},
		"dynamic-element": {
			typ: Tuple{
				ElementTypes: []Type{
					String,
					Number,
					DynamicPseudoType,
				},
			},
			val: []Value{
				NewValue(String, "foo"),
				NewValue(Number, 123),
				NewValue(String, "false"),
			},
			expected: Value{
				typ: Tuple{ElementTypes: []Type{
					String,
					Number,
					DynamicPseudoType,
				}},
				value: []Value{
					{
						typ:   String,
						value: "foo",
					},
					{
						typ:   Number,
						value: big.NewFloat(123),
					},
					{
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
