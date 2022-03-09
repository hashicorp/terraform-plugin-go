package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestListApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		list          List
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"AttributeName": {
			list:          List{},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt-no-ElementType": {
			list:          List{},
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: nil,
		},
		"ElementKeyInt-ElementType-found": {
			list:          List{ElementType: String},
			step:          ElementKeyInt(123),
			expectedType:  String,
			expectedError: nil,
		},
		"ElementKeyInt-ElementType-negative": {
			list:          List{ElementType: String},
			step:          ElementKeyInt(-1),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyString": {
			list:          List{},
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyValue": {
			list:          List{},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.list.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestListEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		l1    List
		l2    List
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			l1:    List{ElementType: String},
			l2:    List{ElementType: String},
			equal: true,
		},
		"unequal": {
			l1:    List{ElementType: String},
			l2:    List{ElementType: Number},
			equal: false,
		},
		"equal-complex": {
			l1: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			l2: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			l1: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			l2: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			l1:    List{ElementType: String},
			l2:    List{},
			equal: false,
		},
		"unequal-complex-empty": {
			l1: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			l2:    List{ElementType: Object{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.l1.Equal(tc.l2)
			revRes := tc.l2.Equal(tc.l1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but l1.Equal(l2) is %v and l2.Equal(l1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestListIs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		list  List
		other Type
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			list:  List{ElementType: String},
			other: List{ElementType: String},
			equal: true,
		},
		"different-elementtype": {
			list:  List{ElementType: String},
			other: List{ElementType: Number},
			equal: true,
		},
		"equal-complex": {
			list: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			other: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"different-elementtype-complex": {
			list: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			other: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: true,
		},
		"equal-empty": {
			list:  List{ElementType: String},
			other: List{},
			equal: true,
		},
		"unequal-set": {
			list:  List{ElementType: String},
			other: Set{ElementType: String},
			equal: false,
		},
		"equal-complex-empty": {
			list: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			other: List{ElementType: Object{}},
			equal: true,
		},
		"equal-complex-nil": {
			list: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			other: List{},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.list.Is(tc.other)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestListUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		list     List
		other    Type
		expected bool
	}
	tests := map[string]testCase{
		"list-list-string-list-string": {
			list:     List{ElementType: List{ElementType: String}},
			other:    List{ElementType: String},
			expected: false,
		},
		"list-list-string-list-list-string": {
			list:     List{ElementType: List{ElementType: String}},
			other:    List{ElementType: List{ElementType: String}},
			expected: true,
		},
		"list-list-string-dpt": {
			list:     List{ElementType: List{ElementType: String}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"list-list-string-list-dpt": {
			list:     List{ElementType: List{ElementType: String}},
			other:    List{ElementType: DynamicPseudoType},
			expected: true,
		},
		"list-list-string-list-list-dpt": {
			list:     List{ElementType: List{ElementType: String}},
			other:    List{ElementType: List{ElementType: DynamicPseudoType}},
			expected: true,
		},
		"list-string-dpt": {
			list:     List{ElementType: String},
			other:    DynamicPseudoType,
			expected: true,
		},
		"list-string-list-bool": {
			list:     List{ElementType: String},
			other:    List{ElementType: Bool},
			expected: false,
		},
		"list-string-list-dpt": {
			list:     List{ElementType: String},
			other:    List{ElementType: DynamicPseudoType},
			expected: true,
		},
		"list-string-list-string": {
			list:     List{ElementType: String},
			other:    List{ElementType: String},
			expected: true,
		},
		"list-string-map": {
			list:     List{ElementType: String},
			other:    Map{ElementType: String},
			expected: false,
		},
		"list-string-object": {
			list:     List{ElementType: String},
			other:    Object{AttributeTypes: map[string]Type{"test": String}},
			expected: false,
		},
		"list-string-primitive": {
			list:     List{ElementType: String},
			other:    String,
			expected: false,
		},
		"list-string-set-string": {
			list:     List{ElementType: String},
			other:    Set{ElementType: String},
			expected: false,
		},
		"list-string-tuple-string": {
			list:     List{ElementType: String},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.list.UsableAs(tc.other)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
