package tftypes

import (
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
		"known": {
			typ: DynamicPseudoType,
			val: "hello",
			err: regexp.MustCompile(`cannot have DynamicPseudoType with known value, DynamicPseudoType can only contain null or unknown values`),
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
