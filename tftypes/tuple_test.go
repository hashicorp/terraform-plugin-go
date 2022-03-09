package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTupleApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		tuple         Tuple
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"AttributeName": {
			tuple:         Tuple{},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt-no-ElementTypes": {
			tuple:         Tuple{},
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt-ElementTypes-found": {
			tuple:         Tuple{ElementTypes: []Type{String}},
			step:          ElementKeyInt(0),
			expectedType:  String,
			expectedError: nil,
		},
		"ElementKeyInt-ElementTypes-negative": {
			tuple:         Tuple{ElementTypes: []Type{String}},
			step:          ElementKeyInt(-1),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt-ElementTypes-overflow": {
			tuple:         Tuple{ElementTypes: []Type{String}},
			step:          ElementKeyInt(1),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyString": {
			tuple:         Tuple{},
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyValue": {
			tuple:         Tuple{},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.tuple.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestTupleEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		t1    Tuple
		t2    Tuple
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			equal: true,
		},
		"unequal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{Number, String, Bool}},
			equal: false,
		},
		"unequal-lengths": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number}},
			equal: false,
		},
		"equal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			equal: true,
		},
		"unequal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}}},
			equal: false,
		},
		"unequal-empty": {
			t1:    Tuple{ElementTypes: []Type{String}},
			t2:    Tuple{},
			equal: false,
		},
		"unequal-complex-empty": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2:    Tuple{ElementTypes: []Type{Object{}}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.t1.Equal(tc.t2)
			revRes := tc.t2.Equal(tc.t1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but t1.Equal(t2) is %v and t2.Equal(t1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestTupleIs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		t1    Tuple
		t2    Tuple
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			equal: true,
		},
		"different-elementtypes": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{Number, String, Bool}},
			equal: true,
		},
		"equal-lengths": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number}},
			equal: true,
		},
		"equal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			equal: true,
		},
		"different-elementtypes-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}}},
			equal: true,
		},
		"equal-empty": {
			t1:    Tuple{ElementTypes: []Type{String}},
			t2:    Tuple{},
			equal: true,
		},
		"equal-complex-empty": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2:    Tuple{ElementTypes: []Type{Object{}}},
			equal: true,
		},
		"equal-complex-nil": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2:    Tuple{},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.t1.Is(tc.t2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestTupleUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		tuple    Tuple
		other    Type
		expected bool
	}
	tests := map[string]testCase{
		"tuple-tuple-string-tuple-string": {
			tuple:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: false,
		},
		"tuple-tuple-string-tuple-tuple-string": {
			tuple:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			other:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			expected: true,
		},
		"tuple-tuple-string-dpt": {
			tuple:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"tuple-tuple-string-tuple-dpt": {
			tuple:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			other:    Tuple{ElementTypes: []Type{DynamicPseudoType}},
			expected: true,
		},
		"tuple-tuple-string-tuple-tuple-dpt": {
			tuple:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{String}}}},
			other:    Tuple{ElementTypes: []Type{Tuple{ElementTypes: []Type{DynamicPseudoType}}}},
			expected: true,
		},
		"tuple-string-dpt": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"tuple-string-list-string": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    List{ElementType: String},
			expected: false,
		},
		"tuple-string-map": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Map{ElementType: String},
			expected: false,
		},
		"tuple-string-object": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Object{AttributeTypes: map[string]Type{"test": String}},
			expected: false,
		},
		"tuple-string-primitive": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    String,
			expected: false,
		},
		"tuple-string-set-string": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Set{ElementType: String},
			expected: false,
		},
		"tuple-string-tuple-bool": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Tuple{ElementTypes: []Type{Bool}},
			expected: false,
		},
		"tuple-string-tuple-bool-string": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Tuple{ElementTypes: []Type{Bool, String}},
			expected: false,
		},
		"tuple-string-tuple-dpt": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Tuple{ElementTypes: []Type{DynamicPseudoType}},
			expected: true,
		},
		"tuple-string-tuple-string": {
			tuple:    Tuple{ElementTypes: []Type{String}},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.tuple.UsableAs(tc.other)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
