package tftypes

import (
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func numberComparer(i, j *big.Float) bool {
	return (i == nil && j == nil) || (i != nil && j != nil && i.Cmp(j) == 0)
}

func valueComparer(i, j Value) bool {
	if !i.typ.Is(j.typ) {
		return false
	}
	return cmp.Equal(i.value, j.value, cmp.Comparer(numberComparer), cmp.Comparer(valueComparer))
}

func TestValueAs(t *testing.T) {
	t.Parallel()
	type testCase struct {
		in       Value
		as       interface{}
		expected interface{}
	}

	strPointer := func(in string) *string {
		return &in
	}
	strPointerPointer := func(in *string) **string {
		return &in
	}
	numberPointerPointer := func(in *big.Float) **big.Float {
		return &in
	}
	boolPointer := func(in bool) *bool {
		return &in
	}
	boolPointerPointer := func(in *bool) **bool {
		return &in
	}
	mapPointer := func(in map[string]Value) *map[string]Value {
		return &in
	}
	mapPointerPointer := func(in *map[string]Value) **map[string]Value {
		return &in
	}
	slicePointer := func(in []Value) *[]Value {
		return &in
	}
	slicePointerPointer := func(in *[]Value) **[]Value {
		return &in
	}
	tests := map[string]testCase{
		"string": {
			in:       NewValue(String, "hello"),
			as:       strPointer(""),
			expected: strPointer("hello"),
		},
		"string-null": {
			in:       NewValue(String, nil),
			as:       strPointer("this value should be removed"),
			expected: strPointer(""),
		},
		"string-pointer": {
			in:       NewValue(String, "hello"),
			as:       strPointerPointer(strPointer("")),
			expected: strPointerPointer(strPointer("hello")),
		},
		"string-pointer-null": {
			in:       NewValue(String, nil),
			as:       strPointerPointer(strPointer("this value should be removed")),
			expected: strPointerPointer(nil),
		},
		"number": {
			in:       NewValue(Number, big.NewFloat(123)),
			as:       big.NewFloat(0),
			expected: big.NewFloat(123),
		},
		"number-null": {
			in:       NewValue(Number, nil),
			as:       big.NewFloat(123),
			expected: big.NewFloat(0),
		},
		"number-pointer": {
			in:       NewValue(Number, big.NewFloat(123)),
			as:       numberPointerPointer(big.NewFloat(0)),
			expected: numberPointerPointer(big.NewFloat(123)),
		},
		"number-pointer-null": {
			in:       NewValue(Number, nil),
			as:       numberPointerPointer(big.NewFloat(123)),
			expected: numberPointerPointer(nil),
		},
		"bool": {
			in:       NewValue(Bool, true),
			as:       boolPointer(false),
			expected: boolPointer(true),
		},
		"bool-null": {
			in:       NewValue(Bool, nil),
			as:       boolPointer(true),
			expected: boolPointer(false),
		},
		"bool-pointer": {
			in:       NewValue(Bool, true),
			as:       boolPointerPointer(boolPointer(false)),
			expected: boolPointerPointer(boolPointer(true)),
		},
		"bool-pointer-null": {
			in:       NewValue(Bool, nil),
			as:       boolPointerPointer(boolPointer(true)),
			expected: boolPointerPointer(nil),
		},
		"map": {
			in: NewValue(Map{AttributeType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			as: mapPointer(map[string]Value{}),
			expected: mapPointer(map[string]Value{
				"hello": NewValue(String, "world"),
			}),
		},
		"map-null": {
			in: NewValue(Map{AttributeType: String}, nil),
			as: mapPointer(map[string]Value{
				"a": NewValue(String, "this should be removed"),
			}),
			expected: mapPointer(map[string]Value{}),
		},
		"map-pointer": {
			in: NewValue(Map{AttributeType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			as: mapPointerPointer(mapPointer(map[string]Value{})),
			expected: mapPointerPointer(mapPointer(map[string]Value{
				"hello": NewValue(String, "world"),
			})),
		},
		"map-pointer-null": {
			in: NewValue(Map{AttributeType: String}, nil),
			as: mapPointerPointer(mapPointer(map[string]Value{
				"a": NewValue(String, "this should be removed"),
			})),
			expected: mapPointerPointer(nil),
		},
		"list": {
			in:       NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointer([]Value{}),
			expected: slicePointer([]Value{NewValue(String, "hello")}),
		},
		"list-null": {
			in:       NewValue(List{ElementType: String}, nil),
			as:       slicePointer([]Value{NewValue(String, "hello")}),
			expected: slicePointer([]Value{}),
		},
		"list-pointer": {
			in:       NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointerPointer(slicePointer([]Value{})),
			expected: slicePointerPointer(slicePointer([]Value{NewValue(String, "hello")})),
		},
		"list-pointer-null": {
			in:       NewValue(List{ElementType: String}, nil),
			as:       slicePointerPointer(slicePointer([]Value{NewValue(String, "hello")})),
			expected: slicePointerPointer(nil),
		},
		"set": {
			in:       NewValue(Set{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointer([]Value{}),
			expected: slicePointer([]Value{NewValue(String, "hello")}),
		},
		"set-null": {
			in:       NewValue(Set{ElementType: String}, nil),
			as:       slicePointer([]Value{NewValue(String, "hello")}),
			expected: slicePointer([]Value{}),
		},
		"set-pointer": {
			in:       NewValue(Set{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointerPointer(slicePointer([]Value{})),
			expected: slicePointerPointer(slicePointer([]Value{NewValue(String, "hello")})),
		},
		"set-pointer-null": {
			in:       NewValue(Set{ElementType: String}, nil),
			as:       slicePointerPointer(slicePointer([]Value{NewValue(String, "hello")})),
			expected: slicePointerPointer(nil),
		},
		"object": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
			as: mapPointer(map[string]Value{}),
			expected: mapPointer(map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
		},
		"object-null": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, nil),
			as: mapPointer(map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
			expected: mapPointer(map[string]Value{}),
		},
		"object-pointer": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
			as: mapPointerPointer(mapPointer(map[string]Value{})),
			expected: mapPointerPointer(mapPointer(map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			})),
		},
		"object-pointer-null": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, nil),
			as: mapPointerPointer(mapPointer(map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			})),
			expected: mapPointerPointer(nil),
		},
		"tuple": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			as: slicePointer([]Value{}),
			expected: slicePointer([]Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
		},
		"tuple-null": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, nil),
			as: slicePointer([]Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			expected: slicePointer([]Value{}),
		},
		"tuple-pointer": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			as: slicePointerPointer(slicePointer([]Value{})),
			expected: slicePointerPointer(slicePointer([]Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			})),
		},
		"tuple-pointer-null": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, nil),
			as: slicePointerPointer(slicePointer([]Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			})),
			expected: slicePointerPointer(nil),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			err := test.in.As(test.as)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			if diff := cmp.Diff(test.expected, test.as,
				cmp.Comparer(numberComparer),
				cmp.Comparer(valueComparer)); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestValueIsKnown(t *testing.T) {
	t.Parallel()
	type testCase struct {
		value      Value
		known      bool
		fullyKnown bool
	}
	tests := map[string]testCase{
		"string-known": {
			value:      NewValue(String, "hello"),
			known:      true,
			fullyKnown: true,
		},
		"string-unknown": {
			value:      NewValue(String, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"number-known": {
			value:      NewValue(Number, big.NewFloat(123)),
			known:      true,
			fullyKnown: true,
		},
		"number-unknown": {
			value:      NewValue(Number, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"bool-known": {
			value:      NewValue(Bool, true),
			known:      true,
			fullyKnown: true,
		},
		"bool-unknown": {
			value:      NewValue(Bool, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"list-string-known": {
			value:      NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			known:      true,
			fullyKnown: true,
		},
		"list-string-partially-known": {
			value:      NewValue(List{ElementType: String}, []Value{NewValue(String, UnknownValue)}),
			known:      true,
			fullyKnown: false,
		},
		"list-string-unknown": {
			value:      NewValue(List{ElementType: String}, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"set-string-known": {
			value:      NewValue(Set{ElementType: String}, []Value{NewValue(String, "hello")}),
			known:      true,
			fullyKnown: true,
		},
		"set-string-partially-known": {
			value:      NewValue(Set{ElementType: String}, []Value{NewValue(String, UnknownValue)}),
			known:      true,
			fullyKnown: false,
		},
		"set-string-unknown": {
			value:      NewValue(Set{ElementType: String}, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"map-string-known": {
			value:      NewValue(Map{AttributeType: String}, map[string]Value{"foo": NewValue(String, "hello")}),
			known:      true,
			fullyKnown: true,
		},
		"map-string-partially-known": {
			value:      NewValue(Map{AttributeType: String}, map[string]Value{"foo": NewValue(String, UnknownValue)}),
			known:      true,
			fullyKnown: false,
		},
		"map-string-unknown": {
			value:      NewValue(Map{AttributeType: String}, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"object-string_number_bool-known": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
			known:      true,
			fullyKnown: true,
		},
		"object-string_number_bool-partially-known": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, UnknownValue),
				"baz": NewValue(Bool, true),
			}),
			known:      true,
			fullyKnown: false,
		},
		"object-string_number_bool-unknown": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"tuple-string_number_bool-known": {
			value: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			known:      true,
			fullyKnown: true,
		},
		"tuple-string_number_bool-partially-known": {
			value: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, UnknownValue),
				NewValue(Bool, true),
			}),
			known:      true,
			fullyKnown: false,
		},
		"tuple-string_number_bool-unknown": {
			value: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, UnknownValue),
			known:      false,
			fullyKnown: false,
		},
		"complicated-known": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						AttributeType: String,
					}},
				}},
			}}, map[string]Value{
				"foo": NewValue(Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						AttributeType: String,
					}},
				}}, []Value{
					NewValue(String, "hello"),
					NewValue(Bool, false),
					NewValue(List{ElementType: Map{
						AttributeType: String,
					}}, []Value{
						NewValue(Map{
							AttributeType: String,
						}, map[string]Value{
							"red":    NewValue(String, "orange"),
							"yellow": NewValue(String, "green"),
							"blue":   NewValue(String, nil),
						}),
						NewValue(Map{
							AttributeType: String,
						}, map[string]Value{
							"a": NewValue(String, "apple"),
							"b": NewValue(String, "banana"),
							"c": NewValue(String, "chili"),
						}),
					}),
				}),
			}),
			known:      true,
			fullyKnown: true,
		},
		"complicated-unknown": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						AttributeType: String,
					}},
				}},
			}}, map[string]Value{
				"foo": NewValue(Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						AttributeType: String,
					}},
				}}, []Value{
					NewValue(String, "hello"),
					NewValue(Bool, false),
					NewValue(List{ElementType: Map{
						AttributeType: String,
					}}, []Value{
						NewValue(Map{
							AttributeType: String,
						}, map[string]Value{
							"red":    NewValue(String, "orange"),
							"yellow": NewValue(String, "green"),
							"blue":   NewValue(String, nil),
						}),
						NewValue(Map{
							AttributeType: String,
						}, map[string]Value{
							"a": NewValue(String, "apple"),
							"b": NewValue(String, UnknownValue),
							"c": NewValue(String, "chili"),
						}),
					}),
				}),
			}),
			known:      true,
			fullyKnown: false,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			known := test.value.IsKnown()
			fullyKnown := test.value.IsFullyKnown()

			if test.known != known {
				t.Errorf("expected known to be %v, is %v", test.known, known)
			}
			if test.fullyKnown != fullyKnown {
				t.Errorf("expected fully known to be %v, is %v", test.fullyKnown, fullyKnown)
			}
		})
	}
}
