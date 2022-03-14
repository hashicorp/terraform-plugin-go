package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSetApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		set           Set
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"AttributeName": {
			set:           Set{},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt": {
			set:           Set{},
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyString": {
			set:           Set{},
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyValue-no-ElementType": {
			set:           Set{},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: nil,
		},
		"ElementKeyValue-ElementType": {
			set:           Set{ElementType: String},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  String,
			expectedError: nil,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.set.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSetEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		s1    Set
		s2    Set
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: String},
			equal: true,
		},
		"unequal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: Number},
			equal: false,
		},
		"equal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			s1:    Set{ElementType: String},
			s2:    Set{},
			equal: false,
		},
		"unequal-complex-empty": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2:    Set{ElementType: Object{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.s1.Equal(tc.s2)
			revRes := tc.s2.Equal(tc.s1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but s1.Equal(s2) is %v and s2.Equal(s1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestSetIs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		s1    Set
		s2    Set
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: String},
			equal: true,
		},
		"different-elementtype": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: Number},
			equal: true,
		},
		"equal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"different-elementtype-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: true,
		},
		"equal-empty": {
			s1:    Set{ElementType: String},
			s2:    Set{},
			equal: true,
		},
		"equal-complex-empty": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2:    Set{ElementType: Object{}},
			equal: true,
		},
		"equal-complex-nil": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2:    Set{},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.s1.Is(tc.s2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestSetUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		set      Set
		other    Type
		expected bool
	}
	tests := map[string]testCase{
		"set-set-string-set-string": {
			set:      Set{ElementType: Set{ElementType: String}},
			other:    Set{ElementType: String},
			expected: false,
		},
		"set-set-string-set-set-string": {
			set:      Set{ElementType: Set{ElementType: String}},
			other:    Set{ElementType: Set{ElementType: String}},
			expected: true,
		},
		"set-set-string-dpt": {
			set:      Set{ElementType: Set{ElementType: String}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"set-set-string-set-dpt": {
			set:      Set{ElementType: Set{ElementType: String}},
			other:    Set{ElementType: DynamicPseudoType},
			expected: true,
		},
		"set-set-string-set-set-dpt": {
			set:      Set{ElementType: Set{ElementType: String}},
			other:    Set{ElementType: Set{ElementType: DynamicPseudoType}},
			expected: true,
		},
		"set-string-dpt": {
			set:      Set{ElementType: String},
			other:    DynamicPseudoType,
			expected: true,
		},
		"set-string-set-bool": {
			set:      Set{ElementType: String},
			other:    Set{ElementType: Bool},
			expected: false,
		},
		"set-string-set-dpt": {
			set:      Set{ElementType: String},
			other:    Set{ElementType: DynamicPseudoType},
			expected: true,
		},
		"set-string-list-string": {
			set:      Set{ElementType: String},
			other:    List{ElementType: String},
			expected: false,
		},
		"set-string-map": {
			set:      Set{ElementType: String},
			other:    Map{ElementType: String},
			expected: false,
		},
		"set-string-object": {
			set:      Set{ElementType: String},
			other:    Object{AttributeTypes: map[string]Type{"test": String}},
			expected: false,
		},
		"set-string-primitive": {
			set:      Set{ElementType: String},
			other:    String,
			expected: false,
		},
		"set-string-set-string": {
			set:      Set{ElementType: String},
			other:    Set{ElementType: String},
			expected: true,
		},
		"set-string-tuple-string": {
			set:      Set{ElementType: String},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.set.UsableAs(tc.other)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
