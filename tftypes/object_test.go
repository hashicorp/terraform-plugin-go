package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestObjectApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		object        Object
		step          AttributePathStep
		expectedType  interface{}
		expectedError error
	}{
		"AttributeName-no-AttributeTypes": {
			object:        Object{},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"AttributeName-AttributeTypes-found": {
			object:        Object{AttributeTypes: map[string]Type{"test": String}},
			step:          AttributeName("test"),
			expectedType:  String,
			expectedError: nil,
		},
		"AttributeName-AttributeTypes-not-found": {
			object:        Object{AttributeTypes: map[string]Type{"other": String}},
			step:          AttributeName("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyInt": {
			object:        Object{},
			step:          ElementKeyInt(123),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyString": {
			object:        Object{},
			step:          ElementKeyString("test"),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
		"ElementKeyValue": {
			object:        Object{},
			step:          ElementKeyValue(NewValue(String, "test")),
			expectedType:  nil,
			expectedError: ErrInvalidStep,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.object.ApplyTerraform5AttributePathStep(testCase.step)

			if !errors.Is(err, testCase.expectedError) {
				t.Errorf("expected error %q, got %s", testCase.expectedError, err)
			}

			if diff := cmp.Diff(got, testCase.expectedType); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestObjectEqual(t *testing.T) {
	t.Parallel()

	type testCase struct {
		o1    Object
		o2    Object
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			equal: true,
		},
		"unequal": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}},
			equal: false,
		},
		"unequal-lengths": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
			}},
			equal: false,
		},
		"unequal-different-keys": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"d": Bool,
			}},
			equal: false,
		},
		"unequal-optional-lengths": {
			o1: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			o2: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
					"b": {},
				},
			},
			equal: false,
		},
		"unequal-optional-attrs": {
			o1: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			o2: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"b": {},
				},
			},
			equal: false,
		},
		"equal-complex": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			equal: true,
		},
		"unequal-complex": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
					"d": DynamicPseudoType,
				}}}},
			equal: false,
		},
		"unequal-empty": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
			}},
			o2:    Object{},
			equal: false,
		},
		"unequal-complex-empty": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2:    Object{AttributeTypes: map[string]Type{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.o1.Equal(tc.o2)
			revRes := tc.o2.Equal(tc.o1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but o1.Equal(o2) is %v and o2.Equal(o1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestObjectIs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		o1    Object
		o2    Object
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			equal: true,
		},
		"different-attributetypes": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}},
			equal: true,
		},
		"equal-lengths": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
			}},
			equal: true,
		},
		"equal-different-keys": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}},
			o2: Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"d": Bool,
			}},
			equal: true,
		},
		"equal-optional-lengths": {
			o1: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			o2: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
					"b": {},
				},
			},
			equal: true,
		},
		"equal-optional-attrs": {
			o1: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"a": {},
				},
			},
			o2: Object{
				AttributeTypes: map[string]Type{
					"a": String,
					"b": Number,
					"c": Bool,
				},
				OptionalAttributes: map[string]struct{}{
					"b": {},
				},
			},
			equal: true,
		},
		"equal-complex": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			equal: true,
		},
		"different-attributetypes-complex": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
					"d": DynamicPseudoType,
				}}}},
			equal: true,
		},
		"equal-empty": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
			}},
			o2:    Object{AttributeTypes: map[string]Type{}},
			equal: true,
		},
		"equal-nil": {
			o1: Object{AttributeTypes: map[string]Type{
				"a": String,
			}},
			o2:    Object{},
			equal: true,
		},
		"equal-complex-empty": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2:    Object{AttributeTypes: map[string]Type{}},
			equal: true,
		},
		"equal-complex-nil": {
			o1: Object{AttributeTypes: map[string]Type{
				"list": List{ElementType: String},
				"object": Object{AttributeTypes: map[string]Type{
					"a": Number,
					"b": String,
					"c": Bool,
				}}}},
			o2:    Object{},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.o1.Is(tc.o2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestObjectUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		object      Object
		other       Type
		expected    bool
		shouldPanic bool
	}
	tests := map[string]testCase{
		"object-dpt": {
			object: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			other:    DynamicPseudoType,
			expected: true,
		},
		"object-primitive": {
			object: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			other:    String,
			expected: false,
		},
		"object-object-dpt": {
			object: Object{
				AttributeTypes: map[string]Type{
					"dpt":     String,
					"non-dpt": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"dpt":     DynamicPseudoType,
					"non-dpt": String,
				},
			},
			expected: true,
		},
		"object-object-type-inequality": {
			object: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"test": Number,
				},
			},
			expected: false,
		},
		"object-object-type-length": {
			object: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"test":  String,
					"test2": String,
				},
			},
			expected: false,
		},
		"object-object-type-missing": {
			object: Object{
				AttributeTypes: map[string]Type{
					"test": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"test2": String,
				},
			},
			expected: false,
		},
		"object-object-object-dpt": {
			object: Object{
				AttributeTypes: map[string]Type{
					"object": Object{
						AttributeTypes: map[string]Type{
							"dpt": String,
						},
					},
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"object": Object{
						AttributeTypes: map[string]Type{
							"dpt": DynamicPseudoType,
						},
					},
				},
			},
			expected: true,
		},
		"object-object-optional": {
			object: Object{
				AttributeTypes: map[string]Type{
					"optional": String,
					"required": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"optional": String,
					"required": String,
				},
				OptionalAttributes: map[string]struct{}{
					"optional": {},
				},
			},
			expected: true,
		},
		"object-object-optional-inequality": {
			object: Object{
				AttributeTypes: map[string]Type{
					"optional": String,
					"required": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"optional": Number,
					"required": String,
				},
				OptionalAttributes: map[string]struct{}{
					"optional": {},
				},
			},
			expected: false,
		},
		"object-object-required": {
			object: Object{
				AttributeTypes: map[string]Type{
					"required": String,
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"required": String,
				},
			},
			expected: true,
		},
		"object-OptionalAttributes-panic": {
			object: Object{
				AttributeTypes: map[string]Type{
					"optional": String,
					"required": String,
				},
				OptionalAttributes: map[string]struct{}{
					"optional": {},
				},
			},
			other: Object{
				AttributeTypes: map[string]Type{
					"optional": String,
					"required": String,
				},
				OptionalAttributes: map[string]struct{}{
					"optional": {},
				},
			},
			shouldPanic: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var gotPanic string
			var res bool
			func() {
				defer func() {
					if ex := recover(); ex != nil {
						if s, ok := ex.(string); ok {
							gotPanic = s
						} else {
							panic(ex)
						}
					}
				}()
				res = tc.object.UsableAs(tc.other)
			}()
			if (gotPanic != "") != tc.shouldPanic {
				if gotPanic != "" {
					t.Fatalf("Unexpected panic: %s", gotPanic)
				}
				t.Fatalf("Expected panic, but did not panic.")
			}
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
