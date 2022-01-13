package tftypes

import (
	"errors"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func numberComparer(i, j *big.Float) bool {
	return (i == nil && j == nil) || (i != nil && j != nil && i.Cmp(j) == 0)
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
		"string-pointer-in": {
			in:       NewValue(String, strPointer("hello")),
			as:       strPointerPointer(strPointer("")),
			expected: strPointerPointer(strPointer("hello")),
		},
		"string-pointer-null": {
			in:       NewValue(String, nil),
			as:       strPointerPointer(strPointer("this value should be removed")),
			expected: strPointerPointer(nil),
		},
		"string-pointer-string-null": {
			in:       NewValue(String, (*string)(nil)),
			as:       strPointerPointer(strPointer("this value should be removed")),
			expected: strPointerPointer(nil),
		},
		"uninstantiated-string-pointer": {
			in:       NewValue(String, "hello"),
			as:       strPointerPointer(nil),
			expected: strPointerPointer(strPointer("hello")),
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
		"uninstantiated-number-pointer": {
			in:       NewValue(Number, big.NewFloat(123)),
			as:       numberPointerPointer(nil),
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
		"bool-pointer-in": {
			in:       NewValue(Bool, boolPointer(true)),
			as:       boolPointerPointer(boolPointer(false)),
			expected: boolPointerPointer(boolPointer(true)),
		},
		"bool-pointer-null": {
			in:       NewValue(Bool, nil),
			as:       boolPointerPointer(boolPointer(true)),
			expected: boolPointerPointer(nil),
		},
		"bool-pointer-bool-null": {
			in:       NewValue(Bool, (*bool)(nil)),
			as:       boolPointerPointer(boolPointer(true)),
			expected: boolPointerPointer(nil),
		},
		"uninstantiated-bool-pointer": {
			in:       NewValue(Bool, true),
			as:       boolPointerPointer(nil),
			expected: boolPointerPointer(boolPointer(true)),
		},
		"map": {
			in: NewValue(Map{ElementType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			as: mapPointer(map[string]Value{}),
			expected: mapPointer(map[string]Value{
				"hello": NewValue(String, "world"),
			}),
		},
		"map-null": {
			in: NewValue(Map{ElementType: String}, nil),
			as: mapPointer(map[string]Value{
				"a": NewValue(String, "this should be removed"),
			}),
			expected: mapPointer(map[string]Value{}),
		},
		"map-pointer": {
			in: NewValue(Map{ElementType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			as: mapPointerPointer(mapPointer(map[string]Value{})),
			expected: mapPointerPointer(mapPointer(map[string]Value{
				"hello": NewValue(String, "world"),
			})),
		},
		"map-pointer-null": {
			in: NewValue(Map{ElementType: String}, nil),
			as: mapPointerPointer(mapPointer(map[string]Value{
				"a": NewValue(String, "this should be removed"),
			})),
			expected: mapPointerPointer(nil),
		},
		"uninstantiated-map-pointer": {
			in: NewValue(Map{ElementType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			as: mapPointerPointer(nil),
			expected: mapPointerPointer(mapPointer(map[string]Value{
				"hello": NewValue(String, "world"),
			})),
		},
		"list": {
			in:       NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointer([]Value{}),
			expected: slicePointer([]Value{NewValue(String, "hello")}),
		},
		"list-dynamic": {
			in:       NewValue(List{ElementType: DynamicPseudoType}, []Value{NewValue(String, "hello")}),
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
		"list-object": {
			in: NewValue(List{ElementType: Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}}, []Value{NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			})}),
			as: slicePointer([]Value{}),
			expected: slicePointer([]Value{NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			})}),
		},
		"uninstantiated-list": {
			in:       NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointerPointer(nil),
			expected: slicePointerPointer(slicePointer([]Value{NewValue(String, "hello")})),
		},
		"set": {
			in:       NewValue(Set{ElementType: String}, []Value{NewValue(String, "hello")}),
			as:       slicePointer([]Value{}),
			expected: slicePointer([]Value{NewValue(String, "hello")}),
		},
		"set-dynamic": {
			in:       NewValue(Set{ElementType: DynamicPseudoType}, []Value{NewValue(String, "hello")}),
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
		"object-dynamic": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": DynamicPseudoType,
				"bar": DynamicPseudoType,
				"baz": DynamicPseudoType,
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

		"object-list-dynamic": {
			in: NewValue(Object{
				AttributeTypes: map[string]Type{
					"result": List{ElementType: DynamicPseudoType},
				},
			}, map[string]Value{
				"result": NewValue(List{ElementType: Object{
					AttributeTypes: map[string]Type{
						"testcol": String,
					},
				}}, []Value{}),
			}),
			as: mapPointer(map[string]Value{}),
			expected: mapPointer(map[string]Value{
				"result": NewValue(List{ElementType: Object{
					AttributeTypes: map[string]Type{
						"testcol": String,
					},
				}}, []Value{}),
			}),
		},
		"dynamic": {
			in:       NewValue(DynamicPseudoType, "hello"),
			as:       strPointer(""),
			expected: strPointer("hello"),
		},
		"dynamic-null": {
			in:       NewValue(DynamicPseudoType, nil),
			as:       strPointer(""),
			expected: strPointer(""),
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
				cmp.Comparer(numberComparer)); diff != "" {
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
		"dynamic-unknown": {
			value:      NewValue(DynamicPseudoType, UnknownValue),
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
			value:      NewValue(Map{ElementType: String}, map[string]Value{"foo": NewValue(String, "hello")}),
			known:      true,
			fullyKnown: true,
		},
		"map-string-partially-known": {
			value:      NewValue(Map{ElementType: String}, map[string]Value{"foo": NewValue(String, UnknownValue)}),
			known:      true,
			fullyKnown: false,
		},
		"map-string-unknown": {
			value:      NewValue(Map{ElementType: String}, UnknownValue),
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
						ElementType: String,
					}},
				}},
			}}, map[string]Value{
				"foo": NewValue(Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						ElementType: String,
					}},
				}}, []Value{
					NewValue(String, "hello"),
					NewValue(Bool, false),
					NewValue(List{ElementType: Map{
						ElementType: String,
					}}, []Value{
						NewValue(Map{
							ElementType: String,
						}, map[string]Value{
							"red":    NewValue(String, "orange"),
							"yellow": NewValue(String, "green"),
							"blue":   NewValue(String, nil),
						}),
						NewValue(Map{
							ElementType: String,
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
						ElementType: String,
					}},
				}},
			}}, map[string]Value{
				"foo": NewValue(Tuple{ElementTypes: []Type{
					String, Bool, List{ElementType: Map{
						ElementType: String,
					}},
				}}, []Value{
					NewValue(String, "hello"),
					NewValue(Bool, false),
					NewValue(List{ElementType: Map{
						ElementType: String,
					}}, []Value{
						NewValue(Map{
							ElementType: String,
						}, map[string]Value{
							"red":    NewValue(String, "orange"),
							"yellow": NewValue(String, "green"),
							"blue":   NewValue(String, nil),
						}),
						NewValue(Map{
							ElementType: String,
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
		"dynamic-null": {
			value:      NewValue(DynamicPseudoType, nil),
			known:      true,
			fullyKnown: true,
		},
		"object-null": {
			value: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
			}}, nil),
			known:      true,
			fullyKnown: true,
		},
		"map-null": {
			value:      NewValue(Map{ElementType: String}, nil),
			known:      true,
			fullyKnown: true,
		},
		"tuple-null": {
			value:      NewValue(Tuple{ElementTypes: []Type{String}}, nil),
			known:      true,
			fullyKnown: true,
		},
		"list-null": {
			value:      NewValue(List{ElementType: String}, nil),
			known:      true,
			fullyKnown: true,
		},
		"set-null": {
			value:      NewValue(Set{ElementType: String}, nil),
			known:      true,
			fullyKnown: true,
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

func TestValueEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		val1  Value
		val2  Value
		equal bool
	}
	tests := map[string]testCase{
		"stringEqual": {
			val1:  NewValue(String, "hello"),
			val2:  NewValue(String, "hello"),
			equal: true,
		},
		"stringDiff": {
			val1:  NewValue(String, "hello"),
			val2:  NewValue(String, "world"),
			equal: false,
		},
		"boolEqual": {
			val1:  NewValue(Bool, true),
			val2:  NewValue(Bool, true),
			equal: true,
		},
		"boolDiff": {
			val1:  NewValue(Bool, false),
			val2:  NewValue(Bool, true),
			equal: false,
		},
		"numberEqual": {
			val1:  NewValue(Number, big.NewFloat(123)),
			val2:  NewValue(Number, big.NewFloat(0).SetInt64(123)),
			equal: true,
		},
		"numberDiff": {
			val1:  NewValue(Number, big.NewFloat(1)),
			val2:  NewValue(Number, big.NewFloat(2)),
			equal: false,
		},
		"unknownEqual": {
			val1:  NewValue(String, UnknownValue),
			val2:  NewValue(String, UnknownValue),
			equal: true,
		},
		"unknownDiff": {
			val1:  NewValue(String, UnknownValue),
			val2:  NewValue(String, "world"),
			equal: false,
		},
		"listEqual": {
			val1: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			equal: true,
		},
		"listDiffLength": {
			val1: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "abc"),
			}),
			equal: false,
		},
		"listDiffElement": {
			val1: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world!"),
				NewValue(String, "abc"),
			}),
			equal: false,
		},
		"setEqual": {
			val1: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			equal: true,
		},
		"setDiffLength": {
			val1: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "abc"),
			}),
			equal: false,
		},
		"setDiffElement": {
			val1: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
				NewValue(String, "abc"),
			}),
			val2: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world!"),
				NewValue(String, "abc"),
			}),
			equal: false,
		},
		"tupleEqual": {
			val1: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "b"),
					NewValue(String, "c"),
				}),
			}),
			val2: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "b"),
					NewValue(String, "c"),
				}),
			}),
			equal: true,
		},
		"tupleDiff": {
			val1: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "b"),
					NewValue(String, "c"),
				}),
			}),
			val2: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello, world"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "b"),
					NewValue(String, "c"),
				}),
			}),
			equal: false,
		},
		"tupleDiffNested": {
			val1: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "b"),
					NewValue(String, "c"),
				}),
			}),
			val2: NewValue(Tuple{ElementTypes: []Type{
				String, Bool, Number, List{ElementType: String},
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Bool, true),
				NewValue(Number, big.NewFloat(123)),
				NewValue(List{ElementType: String}, []Value{
					NewValue(String, "a"),
					NewValue(String, "d"),
					NewValue(String, "c"),
				}),
			}),
			equal: false,
		},
		"mapEqual": {
			val1: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			val2: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			equal: true,
		},
		"mapDiffLength": {
			val1: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			val2: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			equal: false,
		},
		"mapDiffValue": {
			val1: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			val2: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(4)),
			}),
			equal: false,
		},
		"mapDiffKeys": {
			val1: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			val2: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":  NewValue(Number, big.NewFloat(1)),
				"two":  NewValue(Number, big.NewFloat(2)),
				"four": NewValue(Number, big.NewFloat(3)),
			}),
			equal: false,
		},
		"objectEqual": {
			val1: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true), NewValue(Bool, false),
				}),
			}),
			val2: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true), NewValue(Bool, false),
				}),
			}),
			equal: true,
		},
		"objectDiff": {
			val1: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true), NewValue(Bool, false),
				}),
			}),
			val2: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world!"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true), NewValue(Bool, false),
				}),
			}),
			equal: false,
		},
		"objectDiffNested": {
			val1: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true), NewValue(Bool, false),
				}),
			}),
			val2: NewValue(Object{AttributeTypes: map[string]Type{
				"one":   Number,
				"hello": String,
				"true":  Bool,
				"set":   Set{ElementType: Bool},
			}}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"hello": NewValue(String, "world"),
				"true":  NewValue(Bool, true),
				"set": NewValue(Set{ElementType: Bool}, []Value{
					NewValue(Bool, true),
				}),
			}),
			equal: false,
		},
		"nullEqual": {
			val1:  NewValue(String, nil),
			val2:  NewValue(String, nil),
			equal: true,
		},
		"nullDiff": {
			val1:  NewValue(String, nil),
			val2:  NewValue(String, "hello"),
			equal: false,
		},
		"empty": {
			val1:  Value{},
			val2:  Value{},
			equal: true,
		},
		"emptyDiff": {
			val1: Value{},
			val2: NewValue(String, "hello"),
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if result := test.val1.Equal(test.val2); result != test.equal {
				t.Errorf("expected %v, got %v comparing %s and %s", test.equal, result, test.val1, test.val2)
			}
			if result := test.val2.Equal(test.val1); result != test.equal {
				t.Errorf("expected %v, got %v comparing %s and %s", test.equal, result, test.val2, test.val1)
			}
		})
	}
}

func TestValueApplyTerraform5AttributePathStep(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val  Value
		step AttributePathStep
		res  interface{}
		err  error
	}
	tests := map[string]testCase{
		"string_attr": {
			val:  NewValue(String, "hello"),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"string_eki": {
			val:  NewValue(String, "hello"),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"string_eks": {
			val:  NewValue(String, "hello"),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"string_ekv": {
			val:  NewValue(String, "hello"),
			step: ElementKeyValue(NewValue(String, "hello")),
			err:  ErrInvalidStep,
		},
		"number_attr": {
			val:  NewValue(Number, big.NewFloat(123)),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"number_eki": {
			val:  NewValue(Number, big.NewFloat(123)),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"number_eks": {
			val:  NewValue(Number, big.NewFloat(123)),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"number_ekv": {
			val:  NewValue(Number, big.NewFloat(123)),
			step: ElementKeyValue(NewValue(Number, big.NewFloat(123))),
			err:  ErrInvalidStep,
		},
		"bool_attr": {
			val:  NewValue(Bool, true),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"bool_eki": {
			val:  NewValue(Bool, true),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"bool_eks": {
			val:  NewValue(Bool, true),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"bool_ekv": {
			val:  NewValue(Bool, true),
			step: ElementKeyValue(NewValue(Bool, true)),
			err:  ErrInvalidStep,
		},
		"unknown_attr": {
			val:  NewValue(String, UnknownValue),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"unknown_eki": {
			val:  NewValue(String, UnknownValue),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"unknown_eks": {
			val:  NewValue(String, UnknownValue),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"unknown_ekv": {
			val:  NewValue(String, UnknownValue),
			step: ElementKeyValue(NewValue(String, UnknownValue)),
			err:  ErrInvalidStep,
		},
		"null_attr": {
			val:  NewValue(List{ElementType: Number}, nil),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"null_eki": {
			val:  NewValue(List{ElementType: Number}, nil),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"null_eks": {
			val:  NewValue(List{ElementType: Number}, nil),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"null_ekv": {
			val:  NewValue(List{ElementType: Number}, nil),
			step: ElementKeyValue(NewValue(List{ElementType: Number}, nil)),
			err:  ErrInvalidStep,
		},
		"list_attr": {
			val: NewValue(List{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"list_eki": {
			val: NewValue(List{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyInt(1),
			res:  NewValue(Number, big.NewFloat(2)),
		},
		"list_eki_invalid_index": {
			val: NewValue(List{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyInt(3),
			err:  ErrInvalidStep,
		},
		"list_eks": {
			val: NewValue(List{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"list_ekv": {
			val: NewValue(List{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyValue(NewValue(Number, big.NewFloat(1))),
			err:  ErrInvalidStep,
		},
		"set_eki": {
			val: NewValue(Set{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyInt(1),
			err:  ErrInvalidStep,
		},
		"set_eks": {
			val: NewValue(Set{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"set_ekv": {
			val: NewValue(Set{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyValue(NewValue(Number, big.NewFloat(2))),
			res:  NewValue(Number, big.NewFloat(2)),
		},
		"set_ekv_invalid_value": {
			val: NewValue(Set{ElementType: Number}, []Value{
				NewValue(Number, big.NewFloat(1)),
				NewValue(Number, big.NewFloat(2)),
				NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyValue(NewValue(Number, big.NewFloat(4))),
			err:  ErrInvalidStep,
		},
		"object_attr": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}}, map[string]Value{
				"a": NewValue(String, "hello, world"),
				"b": NewValue(Number, big.NewFloat(1)),
				"c": NewValue(Bool, true),
			}),
			step: AttributeName("a"),
			res:  NewValue(String, "hello, world"),
		},
		"object_attr_notfound": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}}, map[string]Value{
				"a": NewValue(String, "hello, world"),
				"b": NewValue(Number, big.NewFloat(1)),
				"c": NewValue(Bool, true),
			}),
			step: AttributeName("foo"),
			err:  ErrInvalidStep,
		},
		"object_eki": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}}, map[string]Value{
				"a": NewValue(String, "hello, world"),
				"b": NewValue(Number, big.NewFloat(1)),
				"c": NewValue(Bool, true),
			}),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"object_eks": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}}, map[string]Value{
				"a": NewValue(String, "hello, world"),
				"b": NewValue(Number, big.NewFloat(1)),
				"c": NewValue(Bool, true),
			}),
			step: ElementKeyString("foo"),
			err:  ErrInvalidStep,
		},
		"object_ekv": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
				"c": Bool,
			}}, map[string]Value{
				"a": NewValue(String, "hello, world"),
				"b": NewValue(Number, big.NewFloat(1)),
				"c": NewValue(Bool, true),
			}),
			step: ElementKeyValue(NewValue(String, "a")),
			err:  ErrInvalidStep,
		},
		"tuple_attr": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: AttributeName("test"),
			err:  ErrInvalidStep,
		},
		"tuple_eki": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: ElementKeyInt(0),
			res:  NewValue(String, "hello"),
		},
		"tuple_eki_invalid_index": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: ElementKeyInt(3),
			err:  ErrInvalidStep,
		},
		"tuple_eki_negative_index": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: ElementKeyInt(-1),
			err:  ErrInvalidStep,
		},
		"tuple_eks": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: ElementKeyString("test"),
			err:  ErrInvalidStep,
		},
		"tuple_ekv": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			step: ElementKeyValue(NewValue(String, "test")),
			err:  ErrInvalidStep,
		},
		"map_attr": {
			val: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			step: AttributeName("one"),
			err:  ErrInvalidStep,
		},
		"map_eki": {
			val: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyInt(123),
			err:  ErrInvalidStep,
		},
		"map_eks": {
			val: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyString("two"),
			res:  NewValue(Number, big.NewFloat(2)),
		},
		"map_eks_invalid_key": {
			val: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyString("four"),
			err:  ErrInvalidStep,
		},
		"map_ekv": {
			val: NewValue(Map{ElementType: Number}, map[string]Value{
				"one":   NewValue(Number, big.NewFloat(1)),
				"two":   NewValue(Number, big.NewFloat(2)),
				"three": NewValue(Number, big.NewFloat(3)),
			}),
			step: ElementKeyValue(NewValue(String, "four")),
			err:  ErrInvalidStep,
		},
	}
	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			res, err := test.val.ApplyTerraform5AttributePathStep(test.step)
			if !errors.Is(err, test.err) {
				t.Errorf("Expected %v, got %v", test.err, err)
			}
			if !cmp.Equal(res, test.res) {
				t.Errorf("Expected %v, got %v", test.res, res)
			}
		})
	}
}

func TestValueWalkAttributePath(t *testing.T) {
	t.Parallel()
	type testCase struct {
		val      Value
		path     *AttributePath
		expected Value
	}
	tests := map[string]testCase{
		"primitive": {
			val:      NewValue(String, "hello"),
			path:     NewAttributePath(),
			expected: NewValue(String, "hello"),
		},
		"list": {
			val: NewValue(List{ElementType: String}, []Value{
				NewValue(String, "foo"), NewValue(String, "bar"),
			}),
			path:     NewAttributePath().WithElementKeyInt(1),
			expected: NewValue(String, "bar"),
		},
		"set": {
			val: NewValue(Set{ElementType: String}, []Value{
				NewValue(String, "foo"), NewValue(String, "bar"),
			}),
			path:     NewAttributePath().WithElementKeyValue(NewValue(String, "bar")),
			expected: NewValue(String, "bar"),
		},
		"map": {
			val: NewValue(Map{ElementType: String}, map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(String, "bar"),
			}),
			path:     NewAttributePath().WithElementKeyString("b"),
			expected: NewValue(String, "bar"),
		},
		"object": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": String,
				"b": Number,
			}}, map[string]Value{
				"a": NewValue(String, "foo"),
				"b": NewValue(Number, big.NewFloat(123)),
			}),
			path:     NewAttributePath().WithAttributeName("b"),
			expected: NewValue(Number, big.NewFloat(123)),
		},
		"tuple": {
			val: NewValue(Tuple{ElementTypes: []Type{
				String, Number,
			}}, []Value{
				NewValue(String, "foo"),
				NewValue(Number, big.NewFloat(123)),
			}),
			path:     NewAttributePath().WithElementKeyInt(1),
			expected: NewValue(Number, big.NewFloat(123)),
		},
		"complex": {
			val: NewValue(Object{AttributeTypes: map[string]Type{
				"a": List{ElementType: Tuple{ElementTypes: []Type{
					Number, Bool, List{ElementType: String},
				}}},
			}}, map[string]Value{
				"a": NewValue(List{ElementType: Tuple{ElementTypes: []Type{
					Number, Bool, List{ElementType: String},
				}}}, []Value{
					NewValue(Tuple{ElementTypes: []Type{
						Number, Bool, List{ElementType: String},
					}}, []Value{
						NewValue(Number, big.NewFloat(1)),
						NewValue(Bool, true),
						NewValue(List{ElementType: String}, []Value{
							NewValue(String, "foo"),
							NewValue(String, "bar"),
						}),
					}),
					NewValue(Tuple{ElementTypes: []Type{
						Number, Bool, List{ElementType: String},
					}}, []Value{
						NewValue(Number, big.NewFloat(456)),
						NewValue(Bool, false),
						NewValue(List{ElementType: String}, []Value{
							NewValue(String, "hello"),
							NewValue(String, "world"),
						}),
					}),
				}),
			}),
			path:     NewAttributePath().WithAttributeName("a").WithElementKeyInt(1).WithElementKeyInt(2).WithElementKeyInt(1),
			expected: NewValue(String, "world"),
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, remaining, err := WalkAttributePath(test.val, test.path)
			if err != nil {
				t.Fatalf("error walking attribute path, %v still remains in the path: %s", remaining, err)
			}
			if res, ok := result.(Value); !ok {
				t.Errorf("got non-Value result %v", result)
			} else if !res.Equal(test.expected) {
				t.Errorf("expected %s, got %s", test.expected, res)
			}
		})
	}
}

func TestValueString(t *testing.T) {
	t.Parallel()
	type testCase struct {
		in       Value
		expected string
	}

	tests := map[string]testCase{
		"nil": {
			in:       Value{},
			expected: "invalid typeless tftypes.Value<>",
		},
		"string": {
			in:       NewValue(String, "hello"),
			expected: "tftypes.String<\"hello\">",
		},
		"string-null": {
			in:       NewValue(String, nil),
			expected: "tftypes.String<null>",
		},
		"number": {
			in:       NewValue(Number, big.NewFloat(123)),
			expected: "tftypes.Number<\"123\">",
		},
		"number-null": {
			in:       NewValue(Number, nil),
			expected: "tftypes.Number<null>",
		},
		"bool": {
			in:       NewValue(Bool, true),
			expected: "tftypes.Bool<\"true\">",
		},
		"bool-null": {
			in:       NewValue(Bool, nil),
			expected: "tftypes.Bool<null>",
		},
		"dynamic-null": {
			in:       NewValue(DynamicPseudoType, nil),
			expected: "tftypes.DynamicPseudoType<null>",
		},
		"map": {
			in: NewValue(Map{ElementType: String}, map[string]Value{
				"hello": NewValue(String, "world"),
			}),
			expected: `tftypes.Map[tftypes.String]<"hello":tftypes.String<"world">>`,
		},
		"map-null": {
			in:       NewValue(Map{ElementType: String}, nil),
			expected: "tftypes.Map[tftypes.String]<null>",
		},
		"list": {
			in:       NewValue(List{ElementType: String}, []Value{NewValue(String, "hello")}),
			expected: `tftypes.List[tftypes.String]<tftypes.String<"hello">>`,
		},
		"list-dynamic": {
			in:       NewValue(List{ElementType: DynamicPseudoType}, []Value{NewValue(String, "hello")}),
			expected: `tftypes.List[tftypes.DynamicPseudoType]<tftypes.String<"hello">>`,
		},
		"list-null": {
			in:       NewValue(List{ElementType: String}, nil),
			expected: "tftypes.List[tftypes.String]<null>",
		},
		"list-object": {
			in: NewValue(List{ElementType: Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}}, []Value{NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			})}),
			expected: `tftypes.List[tftypes.Object["bar":tftypes.Number, "baz":tftypes.Bool, "foo":tftypes.String]]<tftypes.Object["bar":tftypes.Number, "baz":tftypes.Bool, "foo":tftypes.String]<"bar":tftypes.Number<"123">, "baz":tftypes.Bool<"true">, "foo":tftypes.String<"hello">>>`,
		},
		"set": {
			in:       NewValue(Set{ElementType: String}, []Value{NewValue(String, "hello")}),
			expected: `tftypes.Set[tftypes.String]<tftypes.String<"hello">>`,
		},
		"set-dynamic": {
			in:       NewValue(Set{ElementType: DynamicPseudoType}, []Value{NewValue(String, "hello")}),
			expected: `tftypes.Set[tftypes.DynamicPseudoType]<tftypes.String<"hello">>`,
		},
		"set-null": {
			in:       NewValue(Set{ElementType: String}, nil),
			expected: "tftypes.Set[tftypes.String]<null>",
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
			expected: `tftypes.Object["bar":tftypes.Number, "baz":tftypes.Bool, "foo":tftypes.String]<"bar":tftypes.Number<"123">, "baz":tftypes.Bool<"true">, "foo":tftypes.String<"hello">>`,
		},
		"object-dynamic": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": DynamicPseudoType,
				"bar": DynamicPseudoType,
				"baz": DynamicPseudoType,
			}}, map[string]Value{
				"foo": NewValue(String, "hello"),
				"bar": NewValue(Number, big.NewFloat(123)),
				"baz": NewValue(Bool, true),
			}),
			expected: `tftypes.Object["bar":tftypes.DynamicPseudoType, "baz":tftypes.DynamicPseudoType, "foo":tftypes.DynamicPseudoType]<"bar":tftypes.Number<"123">, "baz":tftypes.Bool<"true">, "foo":tftypes.String<"hello">>`,
		},
		"object-null": {
			in: NewValue(Object{AttributeTypes: map[string]Type{
				"foo": String,
				"bar": Number,
				"baz": Bool,
			}}, nil),
			expected: `tftypes.Object["bar":tftypes.Number, "baz":tftypes.Bool, "foo":tftypes.String]<null>`,
		},
		"tuple": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, []Value{
				NewValue(String, "hello"),
				NewValue(Number, big.NewFloat(123)),
				NewValue(Bool, true),
			}),
			expected: `tftypes.Tuple[tftypes.String, tftypes.Number, tftypes.Bool]<tftypes.String<"hello">, tftypes.Number<"123">, tftypes.Bool<"true">>`,
		},
		"tuple-null": {
			in: NewValue(Tuple{ElementTypes: []Type{
				String, Number, Bool,
			}}, nil),
			expected: "tftypes.Tuple[tftypes.String, tftypes.Number, tftypes.Bool]<null>",
		},
		"object-list-dynamic": {
			in: NewValue(Object{
				AttributeTypes: map[string]Type{
					"result": List{ElementType: DynamicPseudoType},
				},
			}, map[string]Value{
				"result": NewValue(List{ElementType: Object{
					AttributeTypes: map[string]Type{
						"testcol": String,
					},
				}}, []Value{}),
			}),
			expected: `tftypes.Object["result":tftypes.List[tftypes.DynamicPseudoType]]<"result":tftypes.List[tftypes.Object["testcol":tftypes.String]]<>>`,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			str := test.in.String()
			if diff := cmp.Diff(test.expected, str); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}
