package tftypes

import (
	"errors"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func valuePointer(val Value) *Value {
	return &val
}

func TestValueDiffEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		diff1 ValueDiff
		diff2 ValueDiff
		equal bool
	}

	tests := map[string]testCase{
		"pathDiff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyInt(123),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: false,
		},
		"primitiveVal1Diff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test 3")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: false,
		},
		"primitiveVal2Diff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 3")),
			},
			equal: false,
		},
		"primitiveTypeDiff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(Number, big.NewFloat(123))),
				Value2: valuePointer(NewValue(Number, big.NewFloat(1234))),
			},
			equal: false,
		},
		"complexVal1Diff": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "testing"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			equal: false,
		},
		"complexVal2Diff": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test 1"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			equal: false,
		},
		"complexTypeDiff": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "testing"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "testing"),
					NewValue(String, "foo"),
				})),
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			equal: false,
		},
		"val1NilDiff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			equal: false,
		},
		"val2NilDiff": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
				})),
			},
			equal: false,
		},
		"allValsNilDiff": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
				})),
			},
			equal: false,
		},
		"primitiveEqual": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: true,
		},
		"complexEqual": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
				})),
				Value2: valuePointer(NewValue(List{
					ElementType: String,
				}, []Value{
					NewValue(String, "test"),
					NewValue(String, "test 2"),
					NewValue(String, "test 3"),
				})),
			},
			equal: true,
		},
		"val1NilEqual": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			equal: true,
		},
		"val2NilEqual": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: nil,
			},
			equal: true,
		},
		"allValsNilEqual": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
			},
			equal: true,
		},
		"val1EmptyEqual": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: &Value{},
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: &Value{},
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			equal: true,
		},
		"val2EmptyEqual": {
			diff1: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: &Value{},
			},
			diff2: ValueDiff{
				Path: NewAttributePath().WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: &Value{},
			},
			equal: true,
		},
		"allValsEmptyEqual": {
			diff1: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: &Value{},
				Value2: &Value{},
			},
			diff2: ValueDiff{
				Path:   NewAttributePath().WithElementKeyString("foo"),
				Value1: &Value{},
				Value2: &Value{},
			},
			equal: true,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			isEqual := test.diff1.Equal(test.diff2)
			if isEqual != test.equal {
				t.Fatalf("expected %v, got %v", test.equal, isEqual)
			}
			isEqual = test.diff2.Equal(test.diff1)
			if isEqual != test.equal {
				t.Fatalf("expected %v, got %v", test.equal, isEqual)
			}
		})
	}
}

func TestValueDiffDiff(t *testing.T) {
	t.Parallel()
	type testCase struct {
		val1  Value
		val2  Value
		diffs []ValueDiff
		err   error
	}

	tests := map[string]testCase{
		"val1Empty": {
			val1: Value{},
			val2: NewValue(String, "bar"),
			err:  errors.New("cannot diff value missing type"),
		},
		"val1EmptyTypeWithValue": {
			val1: Value{value: "foo"},
			val2: NewValue(String, "bar"),
			err:  errors.New("cannot diff value missing type"),
		},
		"val2Empty": {
			val1: NewValue(String, "foo"),
			val2: Value{},
			err:  errors.New("cannot diff value missing type"),
		},
		"val2EmptyTypeWithValue": {
			val1: NewValue(String, "foo"),
			val2: Value{value: "bar"},
			err:  errors.New("cannot diff value missing type"),
		},
		"valsEmptyNoDiff": {
			val1: Value{},
			val2: Value{},
		},
		"primitiveDiff": {
			val1: NewValue(String, "foo"),
			val2: NewValue(String, "bar"),
			diffs: []ValueDiff{
				{
					Value1: valuePointer(NewValue(String, "foo")),
					Value2: valuePointer(NewValue(String, "bar")),
				},
			},
		},
		"primitiveNoDiff": {
			val1: NewValue(String, "foo"),
			val2: NewValue(String, "foo"),
		},
		"primitiveTypeError": {
			val1: NewValue(String, "foo"),
			val2: NewValue(Bool, true),
			err:  errors.New("Can't diff values of different types"),
		},
		"listVal2Longer": {
			val1: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "foo"),
				NewValue(String, "bar"),
			}),
			val2: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "foo"),
				NewValue(String, "bar"),
				NewValue(String, "baz"),
			}),
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithElementKeyInt(2),
					Value1: nil,
					Value2: valuePointer(NewValue(String, "baz")),
				},
				{
					Value1: valuePointer(NewValue(List{
						ElementType: String,
					}, []Value{
						NewValue(String, "foo"),
						NewValue(String, "bar"),
					})),
					Value2: valuePointer(NewValue(List{
						ElementType: String,
					}, []Value{
						NewValue(String, "foo"),
						NewValue(String, "bar"),
						NewValue(String, "baz"),
					})),
				},
			},
		},
		"listVal1Longer": {
			val1: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "foo"),
				NewValue(String, "bar"),
				NewValue(String, "baz"),
			}),
			val2: NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "foo"),
				NewValue(String, "bar"),
			}),
			diffs: []ValueDiff{
				{
					Value1: valuePointer(NewValue(List{
						ElementType: String,
					}, []Value{
						NewValue(String, "foo"),
						NewValue(String, "bar"),
						NewValue(String, "baz"),
					})),
					Value2: valuePointer(NewValue(List{
						ElementType: String,
					}, []Value{
						NewValue(String, "foo"),
						NewValue(String, "bar"),
					})),
				},
				{
					Path:   NewAttributePath().WithElementKeyInt(2),
					Value1: valuePointer(NewValue(String, "baz")),
					Value2: nil,
				},
			},
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			diffs, err := test.val1.Diff(test.val2)
			if (err == nil && test.err != nil) || (test.err == nil && err != nil) || (test.err != nil && err != nil && test.err.Error() != err.Error()) {
				t.Errorf("Expected error to be %v, got %v", test.err, err)
			}

			if !cmp.Equal(diffs, test.diffs) {
				t.Errorf("Diff mismatch: %s", cmp.Diff(diffs, test.diffs))
			}

			// run the same test, but which value is val1
			_, err = test.val2.Diff(test.val1)
			if (err == nil && test.err != nil) || (test.err == nil && err != nil) || (test.err != nil && err != nil && test.err.Error() != err.Error()) {
				t.Errorf("Expected reversed error to be %v, got %v", test.err, err)
			}
		})
	}
}
