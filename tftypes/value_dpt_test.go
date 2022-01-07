package tftypes

import (
	"math/big"
	"regexp"
	"testing"
)

func Test_newValue_dpt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		typ      Type
		val      interface{}
		err      *regexp.Regexp
		expected Value
	}
	tests := map[string]testCase{
		"*big.Float": {
			typ: DynamicPseudoType,
			val: big.NewFloat(123),
			expected: Value{
				typ: DynamicPseudoType,
				value: Value{
					typ:   Number,
					value: big.NewFloat(123),
				},
			},
		},
		"bool": {
			typ: DynamicPseudoType,
			val: true,
			expected: Value{
				typ:   DynamicPseudoType,
				value: true,
			},
		},
		"float64": {
			typ: DynamicPseudoType,
			val: float64(123),
			expected: Value{
				typ: DynamicPseudoType,
				value: Value{
					typ:   Number,
					value: big.NewFloat(123),
				},
			},
		},
		"int": {
			typ: DynamicPseudoType,
			val: 123,
			expected: Value{
				typ: DynamicPseudoType,
				value: Value{
					typ:   Number,
					value: big.NewFloat(123),
				},
			},
		},
		"int64": {
			typ: DynamicPseudoType,
			val: int64(123),
			expected: Value{
				typ: DynamicPseudoType,
				value: Value{
					typ:   Number,
					value: big.NewFloat(123),
				},
			},
		},
		"object": {
			typ: DynamicPseudoType,
			val: map[string]Value{
				"testkey": NewValue(String, "testvalue"),
			},
			expected: Value{
				typ: DynamicPseudoType,
				value: map[string]Value{
					"testkey": {
						typ:   String,
						value: "testvalue",
					},
				},
			},
		},
		"string": {
			typ: DynamicPseudoType,
			val: "test",
			expected: Value{
				typ:   DynamicPseudoType,
				value: "test",
			},
		},
		"tuple": {
			typ: DynamicPseudoType,
			val: []Value{NewValue(String, "test")},
			expected: Value{
				typ: DynamicPseudoType,
				value: []Value{
					{
						typ:   String,
						value: "test",
					},
				},
			},
		},
		"null": {
			typ: DynamicPseudoType,
			val: nil,
			expected: Value{
				typ:   DynamicPseudoType,
				value: nil,
			},
		},
		"unknown": {
			typ: DynamicPseudoType,
			val: UnknownValue,
			expected: Value{
				typ:   DynamicPseudoType,
				value: UnknownValue,
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
