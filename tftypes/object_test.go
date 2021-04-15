package tftypes

import "testing"

func TestObjectEqual(t *testing.T) {
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
			equal: true,
		},
		"unequal-complex-empty": {
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
			res := tc.o1.Is(tc.o2)
			if res != tc.equal {
				t.Errorf("Expected result to be %v, got %v", tc.equal, res)
			}
		})
	}
}
