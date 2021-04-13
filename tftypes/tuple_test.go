package tftypes

import "testing"

func TestTupleEqual(t *testing.T) {
	type testCase struct {
		t1    Tuple
		t2    Tuple
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			equal: true,
		},
		"unequal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{Number, String, Bool}},
			equal: false,
		},
		"unequal-lengths": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number}},
			equal: false,
		},
		"equal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			equal: true,
		},
		"unequal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}}},
			equal: false,
		},
		"unequal-empty": {
			t1:    Tuple{ElementTypes: []Type{String}},
			t2:    Tuple{},
			equal: false,
		},
		"unequal-complex-empty": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2:    Tuple{ElementTypes: []Type{Object{}}},
			equal: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.t1.Equal(tc.t2)
			revRes := tc.t2.Equal(tc.t1)
			if res != revRes {
				t.Errorf("Expected Equal to be commutative, but t1.Equal(t2) is %v and t2.Equal(t1) is %v", res, revRes)
			}
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}

func TestTupleIs(t *testing.T) {
	type testCase struct {
		t1    Tuple
		t2    Tuple
		equal bool
	}
	tests := map[string]testCase{
		"equal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			equal: true,
		},
		"unequal": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{Number, String, Bool}},
			equal: false,
		},
		"unequal-lengths": {
			t1:    Tuple{ElementTypes: []Type{String, Number, Bool}},
			t2:    Tuple{ElementTypes: []Type{String, Number}},
			equal: false,
		},
		"equal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			equal: true,
		},
		"unequal-complex": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
				"d": DynamicPseudoType,
			}}}},
			equal: false,
		},
		"unequal-empty": {
			t1:    Tuple{ElementTypes: []Type{String}},
			t2:    Tuple{},
			equal: true,
		},
		"unequal-complex-empty": {
			t1: Tuple{ElementTypes: []Type{Object{AttributeTypes: map[string]Type{
				"a": Number,
				"b": String,
				"c": Bool,
			}}}},
			t2:    Tuple{ElementTypes: []Type{Object{}}},
			equal: true,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			res := tc.t1.Is(tc.t2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}
