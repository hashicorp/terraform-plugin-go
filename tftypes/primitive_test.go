package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrimitiveApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		primitive     primitive
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"Bool-AttributeName": {
			primitive:     Bool,
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Bool-ElementKeyInt": {
			primitive:     Bool,
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Bool-ElementKeyString": {
			primitive:     Bool,
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Bool-ElementKeyValue": {
			primitive:     Bool,
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"DynamicPseudoType-AttributeName": {
			primitive:     DynamicPseudoType,
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"DynamicPseudoType-ElementKeyInt": {
			primitive:     DynamicPseudoType,
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"DynamicPseudoType-ElementKeyString": {
			primitive:     DynamicPseudoType,
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"DynamicPseudoType-ElementKeyValue": {
			primitive:     DynamicPseudoType,
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Number-AttributeName": {
			primitive:     Number,
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Number-ElementKeyInt": {
			primitive:     Number,
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Number-ElementKeyString": {
			primitive:     Number,
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"Number-ElementKeyValue": {
			primitive:     Number,
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"String-AttributeName": {
			primitive:     String,
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"String-ElementKeyInt": {
			primitive:     String,
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"String-ElementKeyString": {
			primitive:     String,
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"String-ElementKeyValue": {
			primitive:     String,
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.primitive.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPrimitiveEqual(t *testing.T) {
	type testCase struct {
		p1    primitive
		p2    primitive
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			p1:    String,
			p2:    String,
			equal: true,
		},
		"unequal": {
			p1:    String,
			p2:    Number,
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.p1.Equal(tc.p2)
			revRes := tc.p2.Equal(tc.p1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but p1.Equal(p2) is %v and p2.Equal(p1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestPrimitiveIs(t *testing.T) {
	type testCase struct {
		p1    primitive
		p2    primitive
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			p1:    String,
			p2:    String,
			equal: true,
		},
		"unequal": {
			p1:    String,
			p2:    Number,
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.p1.Is(tc.p2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestPrimitiveUsableAs(t *testing.T) {
	type testCase struct {
		p        primitive
		o        Type
		expected bool
	}
	tests := map[string]testCase{
		"string-string": {
			p:        String,
			o:        String,
			expected: true,
		},
		"string-number": {
			p:        String,
			o:        Number,
			expected: false,
		},
		"string-bool": {
			p:        String,
			o:        Bool,
			expected: false,
		},
		"string-dpt": {
			p:        String,
			o:        DynamicPseudoType,
			expected: true,
		},
		"string-list": {
			p:        String,
			o:        List{ElementType: String},
			expected: false,
		},
		"dpt-string": {
			p:        DynamicPseudoType,
			o:        String,
			expected: false,
		},
		"dpt-number": {
			p:        DynamicPseudoType,
			o:        Number,
			expected: false,
		},
		"dpt-bool": {
			p:        DynamicPseudoType,
			o:        Bool,
			expected: false,
		},
		"dpt-dpt": {
			p:        DynamicPseudoType,
			o:        DynamicPseudoType,
			expected: true,
		},
		"dpt-list": {
			p:        DynamicPseudoType,
			o:        List{ElementType: String},
			expected: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.p.UsableAs(tc.o)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
