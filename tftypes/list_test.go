package tftypes

import "testing"

func TestListEqual(t *testing.T) {
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
			equal: true,
		},
		"unequal-complex-empty": {
			l1: List{ElementType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			l2:    List{ElementType: Object{}},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.l1.Is(tc.l2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}
