package tftypes

// these tests are based heavily on github.com/zclconf/go-cty
// used under the MIT License
//
// Copyright (c) 2017-2018 Martin Atkins
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWalk(t *testing.T) {
	valType := Object{
		AttributeTypes: map[string]Type{
			"string":       String,
			"number":       Number,
			"bool":         Bool,
			"list":         List{ElementType: Bool},
			"list_empty":   List{ElementType: Bool},
			"set":          Set{ElementType: Bool},
			"set_empty":    Set{ElementType: Bool},
			"tuple":        Tuple{ElementTypes: []Type{Bool}},
			"tuple_empty":  Tuple{ElementTypes: []Type{}},
			"map":          Map{ElementType: Bool},
			"map_empty":    Map{ElementType: Bool},
			"object":       Object{AttributeTypes: map[string]Type{"true": Bool}},
			"object_empty": Object{AttributeTypes: map[string]Type{}},
			"null":         List{ElementType: String},
			"unknown":      Map{ElementType: Bool},
		},
	}
	listContent := []Value{NewValue(Bool, true)}
	setContent := []Value{NewValue(Bool, true)}
	tupleContent := []Value{NewValue(Bool, true)}
	mapContent := map[string]Value{"true": NewValue(Bool, true)}
	objectContent := map[string]Value{"true": NewValue(Bool, true)}
	valContent := map[string]Value{
		"string":       NewValue(valType.AttributeTypes["string"], "hello"),
		"number":       NewValue(valType.AttributeTypes["number"], big.NewFloat(10)),
		"bool":         NewValue(valType.AttributeTypes["bool"], true),
		"list":         NewValue(valType.AttributeTypes["list"], listContent),
		"list_empty":   NewValue(valType.AttributeTypes["list_empty"], []Value{}),
		"set":          NewValue(valType.AttributeTypes["set"], setContent),
		"set_empty":    NewValue(valType.AttributeTypes["set_empty"], []Value{}),
		"tuple":        NewValue(valType.AttributeTypes["tuple"], tupleContent),
		"tuple_empty":  NewValue(valType.AttributeTypes["tuple_empty"], []Value{}),
		"map":          NewValue(valType.AttributeTypes["map"], mapContent),
		"map_empty":    NewValue(valType.AttributeTypes["map_empty"], map[string]Value{}),
		"object":       NewValue(valType.AttributeTypes["object"], objectContent),
		"object_empty": NewValue(valType.AttributeTypes["object_empty"], map[string]Value{}),
		"null":         NewValue(valType.AttributeTypes["null"], nil),
		"unknown":      NewValue(valType.AttributeTypes["unknown"], UnknownValue),
	}
	val := NewValue(valType, valContent)

	gotCalls := map[string]Value{}
	wantCalls := map[string]Value{
		``:                                       val,
		`AttributeName("string")`:                valContent["string"],
		`AttributeName("number")`:                valContent["number"],
		`AttributeName("bool")`:                  valContent["bool"],
		`AttributeName("list")`:                  valContent["list"],
		`AttributeName("list").ElementKeyInt(0)`: listContent[0],
		`AttributeName("list_empty")`:            valContent["list_empty"],
		`AttributeName("set")`:                   valContent["set"],
		`AttributeName("set").ElementKeyValue(` + setContent[0].String() + `)`: setContent[0],
		`AttributeName("set_empty")`:                    valContent["set_empty"],
		`AttributeName("tuple")`:                        valContent["tuple"],
		`AttributeName("tuple").ElementKeyInt(0)`:       tupleContent[0],
		`AttributeName("tuple_empty")`:                  valContent["tuple_empty"],
		`AttributeName("map")`:                          valContent["map"],
		`AttributeName("map").ElementKeyString("true")`: mapContent["true"],
		`AttributeName("map_empty")`:                    valContent["map_empty"],
		`AttributeName("object")`:                       valContent["object"],
		`AttributeName("object").AttributeName("true")`: objectContent["true"],
		`AttributeName("object_empty")`:                 valContent["object_empty"],
		`AttributeName("null")`:                         valContent["null"],
		`AttributeName("unknown")`:                      valContent["unknown"],
	}

	err := Walk(val, func(path *AttributePath, val Value) (bool, error) {
		gotCalls[path.String()] = val
		return true, nil
	})
	if err != nil {
		t.Fatal(err)
	}

	for path, expected := range wantCalls {
		got, ok := gotCalls[path]
		if !ok {
			t.Errorf("no value at %q, expected %s", path, expected.String())
		} else {
			if diff := cmp.Diff(expected, got, cmp.Comparer(numberComparer)); diff != "" {
				t.Errorf("wrong value at %q. (-wanted, +got): %s", path, diff)
			}
		}
	}

	for path, unexpected := range gotCalls {
		if _, ok := wantCalls[path]; ok {
			continue
		}
		t.Errorf("unexpected call to %q, has value %s", path, unexpected.String())
	}
}

func TestTransform(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val         Value
		f           func(*AttributePath, Value) (Value, error)
		diffs       []ValueDiff
		expectedErr error
	}

	newValuePointer := func(typ Type, v interface{}) *Value {
		val := NewValue(typ, v)
		return &val
	}

	tests := map[string]testCase{
		"string": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"string": String,
					},
				},
				map[string]Value{
					"string": NewValue(String, "hello"),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, "hello, world"), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, "hello, world"),
				},
			},
		},
		"string:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"string": String,
					},
				},
				map[string]Value{
					"string": NewValue(String, "hello"),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, nil),
				},
			},
		},
		"string:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"string": String,
					},
				},
				map[string]Value{
					"string": NewValue(String, "hello"),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, UnknownValue),
				},
			},
		},
		"number": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"number": Number,
					},
				},
				map[string]Value{
					"number": NewValue(Number, big.NewFloat(10)),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, big.NewFloat(123)), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, big.NewFloat(123)),
				},
			},
		},
		"number:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"number": Number,
					},
				},
				map[string]Value{
					"number": NewValue(Number, big.NewFloat(10)),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, nil),
				},
			},
		},
		"number:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"number": Number,
					},
				},
				map[string]Value{
					"number": NewValue(Number, big.NewFloat(10)),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, UnknownValue),
				},
			},
		},
		"bool": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"bool": Bool,
					},
				},
				map[string]Value{
					"bool": NewValue(Bool, true),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"bool:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"bool": Bool,
					},
				},
				map[string]Value{
					"bool": NewValue(Bool, true),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"bool:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"bool": Bool,
					},
				},
				map[string]Value{
					"bool": NewValue(Bool, true),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"list:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(List{ElementType: Bool}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("list"),
					Value1: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(List{ElementType: Bool}, nil),
				},
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(List{ElementType: Bool}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("list"),
					Value1: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(List{ElementType: Bool}, UnknownValue),
				},
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list:nullElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"list:unknownElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"list:replaceElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"list:addElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(List{ElementType: Bool}, []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("list"),
					Value1: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(1),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"list:removeElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list": NewValue(List{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(List{ElementType: Bool}, []Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("list"),
					Value1: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(List{ElementType: Bool}, []Value{}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list_empty": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"list_empty": List{ElementType: Bool},
					},
				},
				map[string]Value{
					"list_empty": NewValue(List{ElementType: Bool}, []Value{}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("list_empty")
				if path.Equal(target) {
					return NewValue(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("list_empty"),
					Value1: newValuePointer(List{ElementType: Bool}, []Value{}),
					Value2: newValuePointer(List{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("list_empty").WithElementKeyInt(0),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"set:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(Set{ElementType: Bool}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("set"),
					Value1: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Set{ElementType: Bool}, nil),
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(Set{ElementType: Bool}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("set"),
					Value1: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Set{ElementType: Bool}, UnknownValue),
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set:nullElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, nil)),
					Value1: nil,
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"set:unknownElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, UnknownValue)),
					Value1: nil,
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"set:replaceElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, false)),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"set:addElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("set"),
					Value1: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, false)),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"set:removeElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set": NewValue(Set{ElementType: Bool}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(Set{ElementType: Bool}, []Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("set"),
					Value1: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Set{ElementType: Bool}, []Value{}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set_empty": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"set_empty": Set{ElementType: Bool},
					},
				},
				map[string]Value{
					"set_empty": NewValue(Set{ElementType: Bool}, []Value{}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("set_empty")
				if path.Equal(target) {
					return NewValue(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("set_empty"),
					Value1: newValuePointer(Set{ElementType: Bool}, []Value{}),
					Value2: newValuePointer(Set{ElementType: Bool}, []Value{
						NewValue(Bool, true),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("set_empty").WithElementKeyValue(NewValue(Bool, true)),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"tuple:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"tuple": Tuple{ElementTypes: []Type{Bool}},
					},
				},
				map[string]Value{
					"tuple": NewValue(Tuple{ElementTypes: []Type{Bool}}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("tuple")
				if path.Equal(target) {
					return NewValue(Tuple{ElementTypes: []Type{Bool}}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("tuple"),
					Value1: newValuePointer(Tuple{ElementTypes: []Type{Bool}}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Tuple{ElementTypes: []Type{Bool}}, nil),
				},
				{
					Path:   NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"tuple:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"tuple": Tuple{ElementTypes: []Type{Bool}},
					},
				},
				map[string]Value{
					"tuple": NewValue(Tuple{ElementTypes: []Type{Bool}}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("tuple")
				if path.Equal(target) {
					return NewValue(Tuple{ElementTypes: []Type{Bool}}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("tuple"),
					Value1: newValuePointer(Tuple{ElementTypes: []Type{Bool}}, []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(Tuple{ElementTypes: []Type{Bool}}, UnknownValue),
				},
				{
					Path:   NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"tuple:nullElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"tuple": Tuple{ElementTypes: []Type{Bool}},
					},
				},
				map[string]Value{
					"tuple": NewValue(Tuple{ElementTypes: []Type{Bool}}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"tuple:unknownElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"tuple": Tuple{ElementTypes: []Type{Bool}},
					},
				},
				map[string]Value{
					"tuple": NewValue(Tuple{ElementTypes: []Type{Bool}}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"tuple:replaceElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"tuple": Tuple{ElementTypes: []Type{Bool}},
					},
				},
				map[string]Value{
					"tuple": NewValue(Tuple{ElementTypes: []Type{Bool}}, []Value{NewValue(Bool, true)}),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("map"),
					Value1: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Map{ElementType: Bool}, nil),
				},
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("map"),
					Value1: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Map{ElementType: Bool}, UnknownValue),
				},
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map:nullElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"map:unknownElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"map:replaceElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:addElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, map[string]Value{
						"true":  NewValue(Bool, true),
						"false": NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("map"),
					Value1: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"true":  NewValue(Bool, true),
						"false": NewValue(Bool, false),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("false"),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:removeElement": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map": NewValue(
						Map{ElementType: Bool},
						map[string]Value{
							"true": NewValue(Bool, true),
						},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, map[string]Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: NewAttributePath().WithAttributeName("map"),
					Value1: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Map{ElementType: Bool}, map[string]Value{}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map_empty": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"map_empty": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"map_empty": NewValue(
						Map{ElementType: Bool},
						map[string]Value{},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("map_empty")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, map[string]Value{
						"foo": NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("map_empty"),
					Value1: newValuePointer(Map{ElementType: Bool}, map[string]Value{}),
					Value2: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"foo": NewValue(Bool, true),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("map_empty").WithElementKeyString("foo"),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"object:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"object": Object{AttributeTypes: map[string]Type{"true": Bool}},
					},
				},
				map[string]Value{
					"object": NewValue(
						Object{AttributeTypes: map[string]Type{"true": Bool}},
						map[string]Value{"true": NewValue(Bool, true)},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("object")
				if path.Equal(target) {
					return NewValue(Object{AttributeTypes: map[string]Type{"true": Bool}}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path: NewAttributePath().WithAttributeName("object"),
					Value1: newValuePointer(Object{AttributeTypes: map[string]Type{"true": Bool}}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Object{AttributeTypes: map[string]Type{"true": Bool}}, UnknownValue),
				},
			},
		},
		"object:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"object": Object{AttributeTypes: map[string]Type{"true": Bool}},
					},
				},
				map[string]Value{
					"object": NewValue(
						Object{AttributeTypes: map[string]Type{"true": Bool}},
						map[string]Value{"true": NewValue(Bool, true)},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("object")
				if path.Equal(target) {
					return NewValue(Object{AttributeTypes: map[string]Type{"true": Bool}}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path: NewAttributePath().WithAttributeName("object"),
					Value1: newValuePointer(Object{AttributeTypes: map[string]Type{"true": Bool}}, map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(Object{AttributeTypes: map[string]Type{"true": Bool}}, nil),
				},
			},
		},
		"object:attributeNull": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"object": Object{AttributeTypes: map[string]Type{"true": Bool}},
					},
				},
				map[string]Value{
					"object": NewValue(
						Object{AttributeTypes: map[string]Type{"true": Bool}},
						map[string]Value{"true": NewValue(Bool, true)},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"object:attributeUnknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"object": Object{AttributeTypes: map[string]Type{"true": Bool}},
					},
				},
				map[string]Value{
					"object": NewValue(
						Object{AttributeTypes: map[string]Type{"true": Bool}},
						map[string]Value{"true": NewValue(Bool, true)},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"object:replaceAttribute": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"object": Object{AttributeTypes: map[string]Type{"true": Bool}},
					},
				},
				map[string]Value{
					"object": NewValue(
						Object{AttributeTypes: map[string]Type{"true": Bool}},
						map[string]Value{"true": NewValue(Bool, true)},
					),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"null:unknown": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"null": List{ElementType: String},
					},
				},
				map[string]Value{
					"null": NewValue(List{ElementType: String}, nil),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("null")
				if path.Equal(target) {
					return NewValue(List{ElementType: String}, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("null"),
					Value1: newValuePointer(List{ElementType: String}, nil),
					Value2: newValuePointer(List{ElementType: String}, UnknownValue),
				},
			},
		},
		"null:set": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"null": List{ElementType: String},
					},
				},
				map[string]Value{
					"null": NewValue(List{ElementType: String}, nil),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("null")
				if path.Equal(target) {
					return NewValue(List{ElementType: String}, []Value{
						NewValue(String, "testing"),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("null"),
					Value1: newValuePointer(List{ElementType: String}, nil),
					Value2: newValuePointer(List{ElementType: String}, []Value{
						NewValue(String, "testing"),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("null").WithElementKeyInt(0),
					Value1: nil,
					Value2: newValuePointer(String, "testing"),
				},
			},
		},
		"unknown:null": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"unknown": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"unknown": NewValue(Map{ElementType: Bool}, UnknownValue),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("unknown")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("unknown"),
					Value1: newValuePointer(Map{ElementType: Bool}, UnknownValue),
					Value2: newValuePointer(Map{ElementType: Bool}, nil),
				},
			},
		},
		"unknown:set": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"unknown": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"unknown": NewValue(Map{ElementType: Bool}, UnknownValue),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("unknown")
				if path.Equal(target) {
					return NewValue(Map{ElementType: Bool}, map[string]Value{
						"testing": NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("unknown"),
					Value1: newValuePointer(Map{ElementType: Bool}, UnknownValue),
					Value2: newValuePointer(Map{ElementType: Bool}, map[string]Value{
						"testing": NewValue(Bool, true),
					}),
				},
				{
					Path:   NewAttributePath().WithAttributeName("unknown").WithElementKeyString("testing"),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"null:byValue": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"null":    List{ElementType: String},
						"unknown": Map{ElementType: Bool},
					},
				},
				map[string]Value{
					"null":    NewValue(List{ElementType: String}, nil),
					"unknown": NewValue(Map{ElementType: Bool}, UnknownValue),
				},
			),
			f: func(_ *AttributePath, v Value) (Value, error) {
				if v.IsNull() {
					return NewValue(v.Type(), UnknownValue), nil
				}
				if !v.IsKnown() {
					return NewValue(v.Type(), nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   NewAttributePath().WithAttributeName("unknown"),
					Value1: newValuePointer(Map{ElementType: Bool}, UnknownValue),
					Value2: newValuePointer(Map{ElementType: Bool}, nil),
				},
				{
					Path:   NewAttributePath().WithAttributeName("null"),
					Value1: newValuePointer(List{ElementType: String}, nil),
					Value2: newValuePointer(List{ElementType: String}, UnknownValue),
				},
			},
		},
		"empty:valueMissingTypeError": {
			val: Value{},
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{},
			expectedErr: AttributePathError{
				Path: NewAttributePath(),
				err:  fmt.Errorf("invalid transform: value missing type"),
			},
		},
		"string:missingValueTypeError": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"test": String,
					},
				},
				map[string]Value{
					"test": NewValue(String, "hello"),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("test")
				if path.Equal(target) {
					return Value{}, nil
				}
				return v, nil
			},
			diffs: []ValueDiff{},
			expectedErr: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  fmt.Errorf("missing value type"),
			},
		},
		"string:wrongTypeError": {
			val: NewValue(
				Object{
					AttributeTypes: map[string]Type{
						"test": String,
					},
				},
				map[string]Value{
					"test": NewValue(String, "hello"),
				},
			),
			f: func(path *AttributePath, v Value) (Value, error) {
				target := NewAttributePath().WithAttributeName("test")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{},
			expectedErr: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  fmt.Errorf("can't use tftypes.Bool as tftypes.String"),
			},
		},
	}

	for name, testCase := range tests {
		name, testCase := name, testCase
		t.Run(fmt.Sprintf("testCase=%s", name), func(t *testing.T) {
			t.Parallel()

			gotVal, err := Transform(testCase.val.Copy(), testCase.f)

			if diff := cmp.Diff(testCase.expectedErr, err); diff != "" {
				t.Fatalf("Unexpected error (-wanted, +got): %s", diff)
			}

			diffs, err := testCase.val.Diff(gotVal)

			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			wantedDiffs := map[string]ValueDiff{}
			for _, diff := range testCase.diffs {
				wantedDiffs[diff.Path.String()] = diff
			}
			gotDiffs := map[string]ValueDiff{}
			for _, diff := range diffs {
				gotDiffs[diff.Path.String()] = diff
			}
			for k, diff := range wantedDiffs {
				gotDiff, ok := gotDiffs[k]
				if !ok {
					t.Errorf("Missing diff %s", diff)
				} else if !diff.Equal(gotDiff) {
					t.Errorf("Unexpected diff, wanted %s, got %s", diff, gotDiff)
				}
			}
			for k, diff := range gotDiffs {
				if _, ok := wantedDiffs[k]; ok {
					continue
				}
				t.Errorf("Unexpected diff: %s", diff)
			}
		})
	}
}
