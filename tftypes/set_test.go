package tftypes

import "testing"

func TestSetEqual(t *testing.T) {
	type testCase struct {
		s1    Set
		s2    Set
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: String},
			equal: true,
		},
		"unequal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: Number},
			equal: false,
		},
		"equal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			s1:    Set{ElementType: String},
			s2:    Set{},
			equal: false,
		},
		"unequal-complex-empty": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2:    Set{ElementType: Object{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.s1.Equal(tc.s2)
			revRes := tc.s2.Equal(tc.s1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but s1.Equal(s2) is %v and s2.Equal(s1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestSetIs(t *testing.T) {
	type testCase struct {
		s1    Set
		s2    Set
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: String},
			equal: true,
		},
		"unequal": {
			s1:    Set{ElementType: String},
			s2:    Set{ElementType: Number},
			equal: false,
		},
		"equal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			s1:    Set{ElementType: String},
			s2:    Set{},
			equal: true,
		},
		"unequal-complex-empty": {
			s1: Set{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			s2:    Set{ElementType: Object{}},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.s1.Is(tc.s2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}
