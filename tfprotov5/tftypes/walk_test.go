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
			"map":          Map{AttributeType: Bool},
			"map_empty":    Map{AttributeType: Bool},
			"object":       Object{AttributeTypes: map[string]Type{"true": Bool}},
			"object_empty": Object{AttributeTypes: map[string]Type{}},
			"null":         List{ElementType: String},
			"unknown":      Map{AttributeType: Bool},
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

	err := Walk(val, func(path AttributePath, val Value) (bool, error) {
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
			if diff := cmp.Diff(expected, got, ValueComparer()); diff != "" {
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
			"map":          Map{AttributeType: Bool},
			"map_empty":    Map{AttributeType: Bool},
			"object":       Object{AttributeTypes: map[string]Type{"true": Bool}},
			"object_empty": Object{AttributeTypes: map[string]Type{}},
			"null":         List{ElementType: String},
			"unknown":      Map{AttributeType: Bool},
		},
	}
	valContent := map[string]Value{
		"string":     NewValue(valType.AttributeTypes["string"], "hello"),
		"number":     NewValue(valType.AttributeTypes["number"], big.NewFloat(10)),
		"bool":       NewValue(valType.AttributeTypes["bool"], true),
		"list":       NewValue(valType.AttributeTypes["list"], []Value{NewValue(Bool, true)}),
		"list_empty": NewValue(valType.AttributeTypes["list_empty"], []Value{}),
		"set":        NewValue(valType.AttributeTypes["set"], []Value{NewValue(Bool, true)}),
		"set_empty":  NewValue(valType.AttributeTypes["set_empty"], []Value{}),
		"tuple":      NewValue(valType.AttributeTypes["tuple"], []Value{NewValue(Bool, true)}),
		"map":        NewValue(valType.AttributeTypes["map"], map[string]Value{"true": NewValue(Bool, true)}),
		"map_empty":  NewValue(valType.AttributeTypes["map_empty"], map[string]Value{}),
		"object":     NewValue(valType.AttributeTypes["object"], map[string]Value{"true": NewValue(Bool, true)}),
		"null":       NewValue(valType.AttributeTypes["null"], nil),
		"unknown":    NewValue(valType.AttributeTypes["unknown"], UnknownValue),
	}
	val := NewValue(valType, valContent)

	type testCase struct {
		f     func(AttributePath, Value) (Value, error)
		diffs []ValueDiff
	}

	newValuePointer := func(typ Type, v interface{}) *Value {
		val := NewValue(typ, v)
		return &val
	}

	tests := map[string]testCase{
		"string": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, "hello, world"), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, "hello, world"),
				},
			},
		},
		"string:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, nil),
				},
			},
		},
		"string:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("string")
				if path.Equal(target) {
					return NewValue(String, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("string"),
					Value1: newValuePointer(String, "hello"),
					Value2: newValuePointer(String, UnknownValue),
				},
			},
		},
		"number": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, big.NewFloat(123)), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, big.NewFloat(123)),
				},
			},
		},
		"number:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, nil),
				},
			},
		},
		"number:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("number")
				if path.Equal(target) {
					return NewValue(Number, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("number"),
					Value1: newValuePointer(Number, big.NewFloat(10)),
					Value2: newValuePointer(Number, UnknownValue),
				},
			},
		},
		"bool": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"bool:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"bool:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("bool")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("bool"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"list:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["list"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("list"),
					Value1: newValuePointer(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["list"], nil),
				},
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["list"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("list"),
					Value1: newValuePointer(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["list"], UnknownValue),
				},
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list:nullElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"list:unknownElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"list:replaceElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"list:addElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("list"),
					Value1: newValuePointer(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(1),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"list:removeElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["list"], []Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("list"),
					Value1: newValuePointer(valType.AttributeTypes["list"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["list"], []Value{}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("list").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"list_empty": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("list_empty")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["list_empty"], []Value{
						NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("list_empty"),
					Value1: newValuePointer(valType.AttributeTypes["list_empty"], []Value{}),
					Value2: newValuePointer(valType.AttributeTypes["list_empty"], []Value{
						NewValue(Bool, true),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("list_empty").WithElementKeyInt(0),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"set:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["set"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("set"),
					Value1: newValuePointer(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["set"], nil),
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["set"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("set"),
					Value1: newValuePointer(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["set"], UnknownValue),
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set:nullElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, nil)),
					Value1: nil,
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"set:unknownElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, UnknownValue)),
					Value1: nil,
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"set:replaceElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true))
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, false)),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"set:addElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("set"),
					Value1: newValuePointer(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true), NewValue(Bool, false),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, false)),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"set:removeElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["set"], []Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("set"),
					Value1: newValuePointer(valType.AttributeTypes["set"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["set"], []Value{}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("set").WithElementKeyValue(NewValue(Bool, true)),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"set_empty": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("set_empty")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["set_empty"], []Value{
						NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("set_empty"),
					Value1: newValuePointer(valType.AttributeTypes["set_empty"], []Value{}),
					Value2: newValuePointer(valType.AttributeTypes["set_empty"], []Value{
						NewValue(Bool, true),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("set_empty").WithElementKeyValue(NewValue(Bool, true)),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"tuple:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("tuple")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["tuple"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("tuple"),
					Value1: newValuePointer(valType.AttributeTypes["tuple"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["tuple"], nil),
				},
				{
					Path:   AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"tuple:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("tuple")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["tuple"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("tuple"),
					Value1: newValuePointer(valType.AttributeTypes["tuple"], []Value{
						NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["tuple"], UnknownValue),
				},
				{
					Path:   AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"tuple:nullElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"tuple:unknownElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"tuple:replaceElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0)
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("tuple").WithElementKeyInt(0),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["map"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("map"),
					Value1: newValuePointer(valType.AttributeTypes["map"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["map"], nil),
				},
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["map"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("map"),
					Value1: newValuePointer(valType.AttributeTypes["map"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["map"], UnknownValue),
				},
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map:nullElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"map:unknownElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"map:replaceElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map").WithElementKeyString("true")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:addElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["map"], map[string]Value{
						"true":  NewValue(Bool, true),
						"false": NewValue(Bool, false),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("map"),
					Value1: newValuePointer(valType.AttributeTypes["map"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["map"], map[string]Value{
						"true":  NewValue(Bool, true),
						"false": NewValue(Bool, false),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("false"),
					Value1: nil,
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"map:removeElement": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["map"], map[string]Value{}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path: AttributePath{}.WithAttributeName("map"),
					Value1: newValuePointer(valType.AttributeTypes["map"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["map"], map[string]Value{}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("map").WithElementKeyString("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
			},
		},
		"map_empty": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("map_empty")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["map_empty"], map[string]Value{
						"foo": NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("map_empty"),
					Value1: newValuePointer(valType.AttributeTypes["map_empty"], map[string]Value{}),
					Value2: newValuePointer(valType.AttributeTypes["map_empty"], map[string]Value{
						"foo": NewValue(Bool, true),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("map_empty").WithElementKeyString("foo"),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"object:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("object")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["object"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path: AttributePath{}.WithAttributeName("object"),
					Value1: newValuePointer(valType.AttributeTypes["object"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["object"], UnknownValue),
				},
			},
		},
		"object:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("object")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["object"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: nil,
				},
				{
					Path: AttributePath{}.WithAttributeName("object"),
					Value1: newValuePointer(valType.AttributeTypes["object"], map[string]Value{
						"true": NewValue(Bool, true),
					}),
					Value2: newValuePointer(valType.AttributeTypes["object"], nil),
				},
			},
		},
		"object:attributeNull": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, nil),
				},
			},
		},
		"object:attributeUnknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, UnknownValue),
				},
			},
		},
		"object:replaceAttribute": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("object").WithAttributeName("true")
				if path.Equal(target) {
					return NewValue(Bool, false), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("object").WithAttributeName("true"),
					Value1: newValuePointer(Bool, true),
					Value2: newValuePointer(Bool, false),
				},
			},
		},
		"null:unknown": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("null")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["null"], UnknownValue), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("null"),
					Value1: newValuePointer(valType.AttributeTypes["null"], nil),
					Value2: newValuePointer(valType.AttributeTypes["null"], UnknownValue),
				},
			},
		},
		"null:set": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("null")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["null"], []Value{
						NewValue(String, "testing"),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("null"),
					Value1: newValuePointer(valType.AttributeTypes["null"], nil),
					Value2: newValuePointer(valType.AttributeTypes["null"], []Value{
						NewValue(String, "testing"),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("null").WithElementKeyInt(0),
					Value1: nil,
					Value2: newValuePointer(String, "testing"),
				},
			},
		},
		"unknown:null": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("unknown")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["unknown"], nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("unknown"),
					Value1: newValuePointer(valType.AttributeTypes["unknown"], UnknownValue),
					Value2: newValuePointer(valType.AttributeTypes["unknown"], nil),
				},
			},
		},
		"unknown:set": {
			f: func(path AttributePath, v Value) (Value, error) {
				target := AttributePath{}.WithAttributeName("unknown")
				if path.Equal(target) {
					return NewValue(valType.AttributeTypes["unknown"], map[string]Value{
						"testing": NewValue(Bool, true),
					}), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("unknown"),
					Value1: newValuePointer(valType.AttributeTypes["unknown"], UnknownValue),
					Value2: newValuePointer(valType.AttributeTypes["unknown"], map[string]Value{
						"testing": NewValue(Bool, true),
					}),
				},
				{
					Path:   AttributePath{}.WithAttributeName("unknown").WithElementKeyString("testing"),
					Value1: nil,
					Value2: newValuePointer(Bool, true),
				},
			},
		},
		"null:byValue": {
			f: func(_ AttributePath, v Value) (Value, error) {
				if v.IsNull() {
					// TODO: replace with v.Type() when #58 lands
					return NewValue(v.typ, UnknownValue), nil
				}
				if !v.IsKnown() {
					return NewValue(v.typ, nil), nil
				}
				return v, nil
			},
			diffs: []ValueDiff{
				{
					Path:   AttributePath{}.WithAttributeName("unknown"),
					Value1: newValuePointer(valType.AttributeTypes["unknown"], UnknownValue),
					Value2: newValuePointer(valType.AttributeTypes["unknown"], nil),
				},
				{
					Path:   AttributePath{}.WithAttributeName("null"),
					Value1: newValuePointer(valType.AttributeTypes["null"], nil),
					Value2: newValuePointer(valType.AttributeTypes["null"], UnknownValue),
				},
			},
		},
	}

	for name, testCase := range tests {
		name, testCase, val := name, testCase, val.Copy()
		t.Run(fmt.Sprintf("testCase=%s", name), func(t *testing.T) {
			gotVal, err := Transform(val, testCase.f)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			diffs, err := val.Diff(gotVal)
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
