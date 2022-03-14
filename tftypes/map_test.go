package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMapApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		m             Map
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"AttributeName": {
			m:             Map{},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt": {
			m:             Map{},
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyString-no-ElementType": {
			m:             Map{},
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: nil,
		},
		"ElementKeyString-ElementType": {
			m:             Map{ElementType: String},
			step:          ElementKeyString("test"),
			expectedType:  String,
			expectedError: nil,
		},
		"ElementKeyValue": {
			m:             Map{},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.m.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestMapEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		m1    Map
		m2    Map
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			m1:    Map{ElementType: String},
			m2:    Map{ElementType: String},
			equal: true,
		},
		"unequal": {
			m1:    Map{ElementType: String},
			m2:    Map{ElementType: Number},
			equal: false,
		},
		"equal-complex": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			m1:    Map{ElementType: String},
			m2:    Map{},
			equal: false,
		},
		"unequal-complex-empty": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2:    Map{ElementType: Object{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.m1.Equal(tc.m2)
			revRes := tc.m2.Equal(tc.m1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but m1.Equal(m2) is %v and m2.Equal(m1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestMapIs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		m1    Map
		m2    Map
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			m1:    Map{ElementType: String},
			m2:    Map{ElementType: String},
			equal: true,
		},
		"different-attributetype": {
			m1:    Map{ElementType: String},
			m2:    Map{ElementType: Number},
			equal: true,
		},
		"equal-complex": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"different-attributetype-complex": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: true,
		},
		"equal-empty": {
			m1:    Map{ElementType: String},
			m2:    Map{},
			equal: true,
		},
		"equal-complex-empty": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2:    Map{ElementType: Object{}},
			equal: true,
		},
		"equal-complex-nil": {
			m1: Map{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2:    Map{},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.m1.Is(tc.m2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestMapUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		m        Map
		other    Type
		expected bool
	}
	tests := map[string]testCase{
		"map-map-string-map-string": {
			m:        Map{ElementType: Map{ElementType: String}},
			other:    Map{ElementType: String},
			expected: false,
		},
		"map-map-string-map-map-string": {
			m:        Map{ElementType: Map{ElementType: String}},
			other:    Map{ElementType: Map{ElementType: String}},
			expected: true,
		},
		"map-map-string-dpt": {
			m:        Map{ElementType: Map{ElementType: String}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"map-map-string-map-dpt": {
			m:        Map{ElementType: Map{ElementType: String}},
			other:    Map{ElementType: DynamicPseudoType},
			expected: true,
		},
		"map-map-string-map-map-dpt": {
			m:        Map{ElementType: Map{ElementType: String}},
			other:    Map{ElementType: Map{ElementType: DynamicPseudoType}},
			expected: true,
		},
		"map-string-dpt": {
			m:        Map{ElementType: String},
			other:    DynamicPseudoType,
			expected: true,
		},
		"map-string-list-string": {
			m:        Map{ElementType: String},
			other:    List{ElementType: String},
			expected: false,
		},
		"map-string-map-bool": {
			m:        Map{ElementType: String},
			other:    Map{ElementType: Bool},
			expected: false,
		},
		"map-string-map-dpt": {
			m:        Map{ElementType: String},
			other:    Map{ElementType: DynamicPseudoType},
			expected: true,
		},
		"map-string-map-string": {
			m:        Map{ElementType: String},
			other:    Map{ElementType: String},
			expected: true,
		},
		"map-string-object": {
			m:        Map{ElementType: String},
			other:    Object{AttributeTypes: map[string]Type{"test": String}},
			expected: false,
		},
		"map-string-primitive": {
			m:        Map{ElementType: String},
			other:    String,
			expected: false,
		},
		"map-string-set-string": {
			m:        Map{ElementType: String},
			other:    Set{ElementType: String},
			expected: false,
		},
		"map-string-tuple-string": {
			m:        Map{ElementType: String},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.m.UsableAs(tc.other)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
