package tftypes

import "testing"

func TestMapEqual(t *testing.T) {
	type testCase struct {
		m1    Map
		m2    Map
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			m1:    Map{AttributeType: String},
			m2:    Map{AttributeType: String},
			equal: true,
		},
		"unequal": {
			m1:    Map{AttributeType: String},
			m2:    Map{AttributeType: Number},
			equal: false,
		},
		"equal-complex": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			m1:    Map{AttributeType: String},
			m2:    Map{},
			equal: false,
		},
		"unequal-complex-empty": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2:    Map{AttributeType: Object{}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
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
	type testCase struct {
		m1    Map
		m2    Map
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			m1:    Map{AttributeType: String},
			m2:    Map{AttributeType: String},
			equal: true,
		},
		"unequal": {
			m1:    Map{AttributeType: String},
			m2:    Map{AttributeType: Number},
			equal: false,
		},
		"equal-complex": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			equal: true,
		},
		"unequal-complex": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}},
			equal: false,
		},
		"unequal-empty": {
			m1:    Map{AttributeType: String},
			m2:    Map{},
			equal: true,
		},
		"unequal-complex-empty": {
			m1: Map{AttributeType: Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}},
			m2:    Map{AttributeType: Object{}},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.m1.Is(tc.m2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}
