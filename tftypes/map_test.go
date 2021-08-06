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

func TestMapUsableAs(t *testing.T) {
	t.Parallel()

	type testCase struct {
		m        Map
		other    Type
		expected bool
	}
	tests := map[string]testCase{
		"map-map-string-map-string": {
			m:        Map{AttributeType: Map{AttributeType: String}},
			other:    Map{AttributeType: String},
			expected: false,
		},
		"map-map-string-map-map-string": {
			m:        Map{AttributeType: Map{AttributeType: String}},
			other:    Map{AttributeType: Map{AttributeType: String}},
			expected: true,
		},
		"map-map-string-dpt": {
			m:        Map{AttributeType: Map{AttributeType: String}},
			other:    DynamicPseudoType,
			expected: true,
		},
		"map-map-string-map-dpt": {
			m:        Map{AttributeType: Map{AttributeType: String}},
			other:    Map{AttributeType: DynamicPseudoType},
			expected: true,
		},
		"map-map-string-map-map-dpt": {
			m:        Map{AttributeType: Map{AttributeType: String}},
			other:    Map{AttributeType: Map{AttributeType: DynamicPseudoType}},
			expected: true,
		},
		"map-string-dpt": {
			m:        Map{AttributeType: String},
			other:    DynamicPseudoType,
			expected: true,
		},
		"map-string-list-string": {
			m:        Map{AttributeType: String},
			other:    List{ElementType: String},
			expected: false,
		},
		"map-string-map-bool": {
			m:        Map{AttributeType: String},
			other:    Map{AttributeType: Bool},
			expected: false,
		},
		"map-string-map-dpt": {
			m:        Map{AttributeType: String},
			other:    Map{AttributeType: DynamicPseudoType},
			expected: true,
		},
		"map-string-map-string": {
			m:        Map{AttributeType: String},
			other:    Map{AttributeType: String},
			expected: true,
		},
		"map-string-object": {
			m:        Map{AttributeType: String},
			other:    Object{AttributeTypes: map[string]Type{"test": String}},
			expected: false,
		},
		"map-string-primitive": {
			m:        Map{AttributeType: String},
			other:    String,
			expected: false,
		},
		"map-string-set-string": {
			m:        Map{AttributeType: String},
			other:    Set{ElementType: String},
			expected: false,
		},
		"map-string-tuple-string": {
			m:        Map{AttributeType: String},
			other:    Tuple{ElementTypes: []Type{String}},
			expected: false,
		},
	}
	for name, tc := range tests {
		name, tc := name, tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res := tc.m.UsableAs(tc.other)
			if res != tc.expected {
				t.Fatalf("Expected result to be %v, got %v", tc.expected, res)
			}
		})
	}
}
