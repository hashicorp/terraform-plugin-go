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
		p           primitive
		o           Type
		expected    bool
		shouldPanic bool
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
			p:           DynamicPseudoType,
			o:           String,
			shouldPanic: true,
		},
		"dpt-number": {
			p:           DynamicPseudoType,
			o:           String,
			shouldPanic: true,
		},
		"dpt-bool": {
			p:           DynamicPseudoType,
			o:           Bool,
			shouldPanic: true,
		},
		"dpt-dpt": {
			p:           DynamicPseudoType,
			o:           DynamicPseudoType,
			shouldPanic: true,
		},
		"dpt-list": {
			p:           DynamicPseudoType,
			o:           List{ElementType: String},
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
				res = tc.p.UsableAs(tc.o)
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
