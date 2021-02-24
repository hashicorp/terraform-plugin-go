package tftypes

import (
	"math/big"
	"testing"
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyInt(123),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: false,
		},
		"primitiveVal1Diff": {
			diff1: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test 3")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: false,
		},
		"primitiveVal2Diff": {
			diff1: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 3")),
			},
			equal: false,
		},
		"primitiveTypeDiff": {
			diff1: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(Number, big.NewFloat(123))),
				Value2: valuePointer(NewValue(Number, big.NewFloat(1234))),
			},
			equal: false,
		},
		"complexVal1Diff": {
			diff1: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: nil,
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			diff2: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(String, "test")),
				Value2: valuePointer(NewValue(String, "test 2")),
			},
			equal: true,
		},
		"complexEqual": {
			diff1: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: nil,
				Value2: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, true),
					NewValue(Bool, false),
				})),
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
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
				Path: AttributePath{}.WithElementKeyString("foo"),
				Value1: valuePointer(NewValue(List{
					ElementType: Bool,
				}, []Value{
					NewValue(Bool, false),
					NewValue(Bool, true),
				})),
				Value2: nil,
			},
			diff2: ValueDiff{
				Path: AttributePath{}.WithElementKeyString("foo"),
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
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
			},
			diff2: ValueDiff{
				Path:   AttributePath{}.WithElementKeyString("foo"),
				Value1: nil,
				Value2: nil,
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
	t.Error("not implemented")
}
