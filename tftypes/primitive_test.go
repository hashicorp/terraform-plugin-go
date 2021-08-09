package tftypes

import (
	"testing"
)

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
