package tftypes

import (
	"math/big"
	"testing"
)

func TestNewValue_number(t *testing.T) {
	t.Parallel()
	type testCase struct {
		result   Value
		expected Value
	}

	// helpers to get (potentially nil) pointers to specific numeric types
	uintPtr := func(i uint, null bool) *uint {
		if null {
			return nil
		}
		return &i
	}
	uint8Ptr := func(i uint8, null bool) *uint8 {
		if null {
			return nil
		}
		return &i
	}
	uint16Ptr := func(i uint16, null bool) *uint16 {
		if null {
			return nil
		}
		return &i
	}
	uint32Ptr := func(i uint32, null bool) *uint32 {
		if null {
			return nil
		}
		return &i
	}
	uint64Ptr := func(i uint64, null bool) *uint64 {
		if null {
			return nil
		}
		return &i
	}
	intPtr := func(i int, null bool) *int {
		if null {
			return nil
		}
		return &i
	}
	int8Ptr := func(i int8, null bool) *int8 {
		if null {
			return nil
		}
		return &i
	}
	int16Ptr := func(i int16, null bool) *int16 {
		if null {
			return nil
		}
		return &i
	}
	int32Ptr := func(i int32, null bool) *int32 {
		if null {
			return nil
		}
		return &i
	}
	int64Ptr := func(i int64, null bool) *int64 {
		if null {
			return nil
		}
		return &i
	}
	float64Ptr := func(i float64, null bool) *float64 {
		if null {
			return nil
		}
		return &i
	}

	tests := map[string]testCase{
		"unknown": {
			expected: Value{typ: Number, value: UnknownValue},
			result:   NewValue(Number, UnknownValue),
		},
		"null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, nil),
		},
		"*big.Float-nil": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, (*big.Float)(nil)),
		},
		"*big.Float": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, big.NewFloat(123)),
		},
		"uint": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint(123)),
		},
		"*uint": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uintPtr(123, false)),
		},
		"*uint-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, uintPtr(0, true)),
		},
		"uint8": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint8(123)),
		},
		"*uint8": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint8Ptr(123, false)),
		},
		"*uint8-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, uint8Ptr(0, true)),
		},
		"uint16": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint16(123)),
		},
		"*uint16": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint16Ptr(123, false)),
		},
		"*uint16-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, uint16Ptr(0, true)),
		},
		"uint32": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint32(123)),
		},
		"*uint32": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint32Ptr(123, false)),
		},
		"*uint32-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, uint32Ptr(0, true)),
		},
		"uint64": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint64(123)),
		},
		"*uint64": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, uint64Ptr(123, false)),
		},
		"*uint64-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, uint64Ptr(0, true)),
		},
		"int": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int(123)),
		},
		"*int": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, intPtr(123, false)),
		},
		"*int-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, intPtr(0, true)),
		},
		"int8": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int8(123)),
		},
		"*int8": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int8Ptr(123, false)),
		},
		"*int8-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, int8Ptr(0, true)),
		},
		"int16": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int16(123)),
		},
		"*int16": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int16Ptr(123, false)),
		},
		"*int16-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, int16Ptr(0, true)),
		},
		"int32": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int32(123)),
		},
		"*int32": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int32Ptr(123, false)),
		},
		"*int32-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, int32Ptr(0, true)),
		},
		"int64": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int64(123)),
		},
		"*int64": {
			expected: Value{typ: Number, value: big.NewFloat(123)},
			result:   NewValue(Number, int64Ptr(123, false)),
		},
		"*int64-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, int64Ptr(0, true)),
		},
		"float64": {
			expected: Value{typ: Number, value: big.NewFloat(123.456)},
			result:   NewValue(Number, float64(123.456)),
		},
		"*float64": {
			expected: Value{typ: Number, value: big.NewFloat(123.456)},
			result:   NewValue(Number, float64Ptr(123.456, false)),
		},
		"*float64-null": {
			expected: Value{typ: Number, value: nil},
			result:   NewValue(Number, float64Ptr(0, true)),
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if !test.expected.Equal(test.result) {
				t.Errorf("Expected %s to be equal to %s, wasn't", test.expected, test.result)
			}
		})
	}
}
