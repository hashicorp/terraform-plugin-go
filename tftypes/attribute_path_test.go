package tftypes

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type attributePathStepperTestStruct struct {
	Name   string
	Colors attributePathStepperTestSlice
}

func (a attributePathStepperTestStruct) ApplyTerraform5AttributePathStep(step AttributePathStep) (interface{}, error) {
	attributeName, ok := step.(AttributeName)
	if !ok {
		return nil, fmt.Errorf("unsupported attribute path step type: %T", step)
	}
	switch attributeName {
	case "Name":
		return a.Name, nil
	case "Colors":
		return a.Colors, nil
	}
	return nil, fmt.Errorf("unsupported attribute path step attribute name: %q", attributeName)
}

type attributePathStepperTestSlice []string

func (a attributePathStepperTestSlice) ApplyTerraform5AttributePathStep(step AttributePathStep) (interface{}, error) {
	element, ok := step.(ElementKeyInt)
	if !ok {
		return nil, fmt.Errorf("unsupported attribute path step type: %T", step)
	}
	if element >= 0 && int(element) < len(a) {
		return a[element], nil
	}
	return nil, fmt.Errorf("unsupported attribute path step element key: %d", element)
}

func TestWalkAttributePath(t *testing.T) {
	t.Parallel()
	type testCase struct {
		in       interface{}
		path     *AttributePath
		expected interface{}
	}
	tests := map[string]testCase{
		"msi-root": {
			in: map[string]interface{}{
				"a": map[string]interface{}{
					"red":  true,
					"blue": 123,
				},
				"b": map[string]interface{}{
					"red":  false,
					"blue": 234,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("a"),
				},
			},
			expected: map[string]interface{}{
				"red":  true,
				"blue": 123,
			},
		},
		"msi-full": {
			in: map[string]interface{}{
				"a": map[string]interface{}{
					"red":  true,
					"blue": 123,
				},
				"b": map[string]interface{}{
					"red":  false,
					"blue": 234,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("a"),
					AttributeName("red"),
				},
			},
			expected: true,
		},
		"Object-AttributeName-Bool": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": String,
					"test":  Bool,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: Bool,
		},
		"Object-AttributeName-DynamicPseudoType": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  DynamicPseudoType,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: DynamicPseudoType,
		},
		"Object-AttributeName-List": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  List{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: List{ElementType: String},
		},
		"Object-AttributeName-List-ElementKeyInt": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  List{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
					ElementKeyInt(0),
				},
			},
			expected: String,
		},
		"Object-AttributeName-Map": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Map{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: Map{ElementType: String},
		},
		"Object-AttributeName-Map-ElementKeyString": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Map{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
					ElementKeyString("sub-test"),
				},
			},
			expected: String,
		},
		"Object-AttributeName-Number": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Number,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: Number,
		},
		"Object-AttributeName-Set": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Set{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: Set{ElementType: String},
		},
		"Object-AttributeName-Set-ElementKeyValue": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Set{ElementType: String},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
					ElementKeyValue(NewValue(String, "sub-test")),
				},
			},
			expected: String,
		},
		"Object-AttributeName-String": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  String,
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: String,
		},
		"Object-AttributeName-Tuple": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Tuple{ElementTypes: []Type{Bool, String}},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
				},
			},
			expected: Tuple{ElementTypes: []Type{Bool, String}},
		},
		"Object-AttributeName-Tuple-ElementKeyInt": {
			in: Object{
				AttributeTypes: map[string]Type{
					"other": Bool,
					"test":  Tuple{ElementTypes: []Type{Bool, String}},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					AttributeName("test"),
					ElementKeyInt(1),
				},
			},
			expected: String,
		},
		"slice-interface-root": {
			in: []interface{}{
				map[string]interface{}{
					"a": true,
					"b": 123,
					"c": "hello",
				},
				map[string]interface{}{
					"a": false,
					"b": 1234,
					"c": []interface{}{
						"hello world",
						"happy terraforming",
					},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					ElementKeyInt(1),
				},
			},
			expected: map[string]interface{}{
				"a": false,
				"b": 1234,
				"c": []interface{}{
					"hello world",
					"happy terraforming",
				},
			},
		},
		"slice-interface-full": {
			in: []interface{}{
				map[string]interface{}{
					"a": true,
					"b": 123,
					"c": "hello",
				},
				map[string]interface{}{
					"a": false,
					"b": 1234,
					"c": []interface{}{
						"hello world",
						"happy terraforming",
					},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					ElementKeyInt(1),
					AttributeName("c"),
					ElementKeyInt(0),
				},
			},
			expected: "hello world",
		},
		"attributepathstepper": {
			in: []interface{}{
				attributePathStepperTestStruct{
					Name: "terraform",
					Colors: []string{
						"purple", "white",
					},
				},
				attributePathStepperTestStruct{
					Name: "nomad",
					Colors: []string{
						"green",
					},
				},
			},
			path: &AttributePath{
				steps: []AttributePathStep{
					ElementKeyInt(1),
					AttributeName("Colors"),
					ElementKeyInt(0),
				},
			},
			expected: "green",
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			result, remaining, err := WalkAttributePath(test.in, test.path)
			if err != nil {
				t.Fatalf("error walking attribute path, %v still remains in the path: %s", remaining, err)
			}
			if diff := cmp.Diff(test.expected, result, cmp.Comparer(numberComparer)); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}

func TestAttributePathEqual(t *testing.T) {
	t.Parallel()
	type testCase struct {
		path1 *AttributePath
		path2 *AttributePath
		equal bool
	}

	tests := map[string]testCase{
		"empty": {
			path1: NewAttributePath(),
			path2: NewAttributePath(),
			equal: true,
		},
		"nil": {
			equal: true,
		},
		"empty-and-nil": {
			path1: NewAttributePath(),
			equal: true,
		},
		"an-different-types": {
			path1: NewAttributePath().WithAttributeName("testing"),
			path2: NewAttributePath().WithElementKeyString("testing"),
			equal: false,
		},
		"eks-different-types": {
			path1: NewAttributePath().WithElementKeyString("testing"),
			path2: NewAttributePath().WithAttributeName("testing"),
			equal: false,
		},
		"eki-different-types": {
			path1: NewAttributePath().WithElementKeyInt(1234),
			path2: NewAttributePath().WithAttributeName("testing"),
			equal: false,
		},
		"ekv-different-types": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(String, "testing")),
			path2: NewAttributePath().WithAttributeName("testing"),
			equal: false,
		},
		"an": {
			path1: NewAttributePath().WithAttributeName("testing"),
			path2: NewAttributePath().WithAttributeName("testing"),
			equal: true,
		},
		"an-an": {
			path1: NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			path2: NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			equal: true,
		},
		"eks": {
			path1: NewAttributePath().WithElementKeyString("testing"),
			path2: NewAttributePath().WithElementKeyString("testing"),
			equal: true,
		},
		"eks-eks": {
			path1: NewAttributePath().WithElementKeyString("testing").WithElementKeyString("testing2"),
			path2: NewAttributePath().WithElementKeyString("testing").WithElementKeyString("testing2"),
			equal: true,
		},
		"eki": {
			path1: NewAttributePath().WithElementKeyInt(123),
			path2: NewAttributePath().WithElementKeyInt(123),
			equal: true,
		},
		"eki-eki": {
			path1: NewAttributePath().WithElementKeyInt(123).WithElementKeyInt(456),
			path2: NewAttributePath().WithElementKeyInt(123).WithElementKeyInt(456),
			equal: true,
		},
		"ekv": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})),
			path2: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})),
			equal: true,
		},
		"ekv-ekv": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})).WithElementKeyValue(NewValue(Bool, true)),
			path2: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})).WithElementKeyValue(NewValue(Bool, true)),
			equal: true,
		},
		"an-eks-eki-ekv": {
			path1: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			path2: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			equal: true,
		},
		"ekv-eki-eks-an": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			path2: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			equal: true,
		},
		"an-diff": {
			path1: NewAttributePath().WithAttributeName("testing"),
			path2: NewAttributePath().WithAttributeName("testing2"),
			equal: false,
		},
		"an-an-diff": {
			path1: NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			path2: NewAttributePath().WithAttributeName("testing2").WithAttributeName("testing2"),
			equal: false,
		},
		"an-an-diff-2": {
			path1: NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			path2: NewAttributePath().WithAttributeName("testing").WithAttributeName("testing3"),
			equal: false,
		},
		"eks-diff": {
			path1: NewAttributePath().WithElementKeyString("testing"),
			path2: NewAttributePath().WithElementKeyString("testing2"),
			equal: false,
		},
		"eks-eks-diff": {
			path1: NewAttributePath().WithElementKeyString("testing").WithElementKeyString("testing2"),
			path2: NewAttributePath().WithElementKeyString("testing2").WithElementKeyString("testing2"),
			equal: false,
		},
		"eks-eks-diff-2": {
			path1: NewAttributePath().WithElementKeyString("testing").WithElementKeyString("testing2"),
			path2: NewAttributePath().WithElementKeyString("testing").WithElementKeyString("testing3"),
			equal: false,
		},
		"eki-diff": {
			path1: NewAttributePath().WithElementKeyInt(123),
			path2: NewAttributePath().WithElementKeyInt(1234),
			equal: false,
		},
		"eki-eki-diff": {
			path1: NewAttributePath().WithElementKeyInt(123).WithElementKeyInt(456),
			path2: NewAttributePath().WithElementKeyInt(1234).WithElementKeyInt(456),
			equal: false,
		},
		"eki-eki-diff-2": {
			path1: NewAttributePath().WithElementKeyInt(123).WithElementKeyInt(456),
			path2: NewAttributePath().WithElementKeyInt(123).WithElementKeyInt(4567),
			equal: false,
		},
		"ekv-diff": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})),
			path2: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "fren"),
			})),
			equal: false,
		},
		"ekv-ekv-diff": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})).WithElementKeyValue(NewValue(Bool, true)),
			path2: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "fren"),
			})).WithElementKeyValue(NewValue(Bool, true)),
			equal: false,
		},
		"ekv-ekv-diff-2": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})).WithElementKeyValue(NewValue(Bool, true)),
			path2: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})).WithElementKeyValue(NewValue(Bool, false)),
			equal: false,
		},
		"an-eks-eki-ekv-diff": {
			path1: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			path2: NewAttributePath().WithAttributeName("testing2").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			equal: false,
		},
		"an-eks-eki-ekv-diff-2": {
			path1: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			path2: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing3").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			equal: false,
		},
		"an-eks-eki-ekv-diff-3": {
			path1: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			path2: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(1234).WithElementKeyValue(NewValue(String, "hello, world")),
			equal: false,
		},
		"an-eks-eki-ekv-diff-4": {
			path1: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, world")),
			path2: NewAttributePath().WithAttributeName("testing").WithElementKeyString("testing2").WithElementKeyInt(123).WithElementKeyValue(NewValue(String, "hello, friend")),
			equal: false,
		},
		"ekv-eki-eks-an-diff": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			path2: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(12345)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			equal: false,
		},
		"ekv-eki-eks-an-diff-2": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			path2: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(1234).WithElementKeyString("testing").WithAttributeName("othertesting"),
			equal: false,
		},
		"ekv-eki-eks-an-diff-3": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			path2: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing2").WithAttributeName("othertesting"),
			equal: false,
		},
		"ekv-eki-eks-an-diff-4": {
			path1: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting"),
			path2: NewAttributePath().WithElementKeyValue(NewValue(Object{
				AttributeTypes: map[string]Type{
					"foo": Bool,
					"bar": Number,
				},
			}, map[string]Value{
				"foo": NewValue(Bool, true),
				"bar": NewValue(Number, big.NewFloat(1234)),
			})).WithElementKeyInt(123).WithElementKeyString("testing").WithAttributeName("othertesting2"),
			equal: false,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			isEqual := test.path1.Equal(test.path2)
			if isEqual != test.equal {
				t.Fatalf("expected %v, got %v", test.equal, isEqual)
			}
			isEqual = test.path2.Equal(test.path1)
			if isEqual != test.equal {
				t.Fatalf("expected %v, got %v", test.equal, isEqual)
			}
		})
	}
}

func TestAttributePathLastStep(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		path     *AttributePath
		expected AttributePathStep
	}{
		"empty": {
			path:     NewAttributePath(),
			expected: nil,
		},
		"nil": {
			path:     nil,
			expected: nil,
		},
		"AttributeName": {
			path:     NewAttributePath().WithAttributeName("testing"),
			expected: AttributeName("testing"),
		},
		"AttributeName-AttributeName": {
			path:     NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			expected: AttributeName("testing2"),
		},
		"AttributeName-AttributeName-AttributeName": {
			path:     NewAttributePath().WithElementKeyString("testing").WithAttributeName("testing2").WithAttributeName("testing3"),
			expected: AttributeName("testing3"),
		},
		"ElementKeyInt": {
			path:     NewAttributePath().WithElementKeyInt(1234),
			expected: ElementKeyInt(1234),
		},
		"ElementKeyString": {
			path:     NewAttributePath().WithElementKeyString("testing"),
			expected: ElementKeyString("testing"),
		},
		"ElementKeyValue": {
			path:     NewAttributePath().WithElementKeyValue(NewValue(String, "testing")),
			expected: ElementKeyValue(NewValue(String, "testing")),
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.path.LastStep()

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestAttributePathString(t *testing.T) {
	t.Parallel()
	type testCase struct {
		path     *AttributePath
		expected string
	}

	tests := map[string]testCase{
		"empty": {
			path:     NewAttributePath(),
			expected: "",
		},
		"nil": {
			path:     nil,
			expected: "",
		},
		"attribute-name": {
			path:     NewAttributePath().WithAttributeName("testing"),
			expected: `AttributeName("testing")`,
		},
		"element-key-string": {
			path:     NewAttributePath().WithElementKeyString("testing"),
			expected: `ElementKeyString("testing")`,
		},
		"element-key-int": {
			path:     NewAttributePath().WithElementKeyInt(1234),
			expected: `ElementKeyInt(1234)`,
		},
		"element-key-value": {
			path:     NewAttributePath().WithElementKeyValue(NewValue(String, "testing")),
			expected: `ElementKeyValue(tftypes.String<"testing">)`,
		},
		"an-an": {
			path:     NewAttributePath().WithAttributeName("testing").WithAttributeName("testing2"),
			expected: `AttributeName("testing").AttributeName("testing2")`,
		},
		"long": {
			path:     NewAttributePath().WithElementKeyString("testing").WithElementKeyInt(20).WithAttributeName("testing2"),
			expected: `ElementKeyString("testing").ElementKeyInt(20).AttributeName("testing2")`,
		},
		"ekv-complex": {
			path: NewAttributePath().WithElementKeyValue(NewValue(List{
				ElementType: String,
			}, []Value{
				NewValue(String, "hello"),
				NewValue(String, "world"),
			})),
			expected: `ElementKeyValue(tftypes.List[tftypes.String]<tftypes.String<"hello">, tftypes.String<"world">>)`,
		},
	}

	for name, test := range tests {
		name, test := name, test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			str := test.path.String()
			if diff := cmp.Diff(test.expected, str); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestAttributeNameEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		attributeName AttributeName
		other         AttributePathStep
		expected      bool
	}{
		"AttributeName-different": {
			attributeName: AttributeName("test"),
			other:         AttributeName("other"),
			expected:      false,
		},
		"AttributeName-equal": {
			attributeName: AttributeName("test"),
			other:         AttributeName("test"),
			expected:      true,
		},
		"ElementKeyInt": {
			attributeName: AttributeName("test"),
			other:         ElementKeyInt(1),
			expected:      false,
		},
		"ElementKeyString": {
			attributeName: AttributeName("test"),
			other:         ElementKeyString("test"),
			expected:      false,
		},
		"ElementKeyValue": {
			attributeName: AttributeName("test"),
			other:         ElementKeyValue(NewValue(String, "test")),
			expected:      false,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.attributeName.Equal(tc.other)

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestElementKeyIntEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		elementKeyInt ElementKeyInt
		other         AttributePathStep
		expected      bool
	}{
		"AttributeName": {
			elementKeyInt: ElementKeyInt(1),
			other:         AttributeName("test"),
			expected:      false,
		},
		"ElementKeyInt-different": {
			elementKeyInt: ElementKeyInt(1),
			other:         ElementKeyInt(2),
			expected:      false,
		},
		"ElementKeyInt-equal": {
			elementKeyInt: ElementKeyInt(1),
			other:         ElementKeyInt(1),
			expected:      true,
		},
		"ElementKeyString": {
			elementKeyInt: ElementKeyInt(1),
			other:         ElementKeyString("test"),
			expected:      false,
		},
		"ElementKeyValue": {
			elementKeyInt: ElementKeyInt(1),
			other:         ElementKeyValue(NewValue(String, "test")),
			expected:      false,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.elementKeyInt.Equal(tc.other)

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestElementKeyStringEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		elementKeyString ElementKeyString
		other            AttributePathStep
		expected         bool
	}{
		"AttributeName": {
			elementKeyString: ElementKeyString("test"),
			other:            AttributeName("test"),
			expected:         false,
		},
		"ElementKeyInt": {
			elementKeyString: ElementKeyString("test"),
			other:            ElementKeyInt(1),
			expected:         false,
		},
		"ElementKeyString-different": {
			elementKeyString: ElementKeyString("test"),
			other:            ElementKeyString("other"),
			expected:         false,
		},
		"ElementKeyString-equal": {
			elementKeyString: ElementKeyString("test"),
			other:            ElementKeyString("test"),
			expected:         true,
		},
		"ElementKeyValue": {
			elementKeyString: ElementKeyString("test"),
			other:            ElementKeyValue(NewValue(String, "test")),
			expected:         false,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.elementKeyString.Equal(tc.other)

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}

func TestElementKeyValueEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		elementKeyValue ElementKeyValue
		other           AttributePathStep
		expected        bool
	}{
		"AttributeName-different": {
			elementKeyValue: ElementKeyValue(NewValue(String, "test")),
			other:           AttributeName("test"),
			expected:        false,
		},
		"ElementKeyInt": {
			elementKeyValue: ElementKeyValue(NewValue(String, "test")),
			other:           ElementKeyInt(1),
			expected:        false,
		},
		"ElementKeyString": {
			elementKeyValue: ElementKeyValue(NewValue(String, "test")),
			other:           ElementKeyString("test"),
			expected:        false,
		},
		"ElementKeyValue-different": {
			elementKeyValue: ElementKeyValue(NewValue(String, "test")),
			other:           ElementKeyValue(NewValue(String, "other")),
			expected:        false,
		},
		"ElementKeyValue-equal": {
			elementKeyValue: ElementKeyValue(NewValue(String, "test")),
			other:           ElementKeyValue(NewValue(String, "test")),
			expected:        true,
		},
	}

	for name, tc := range testCases {
		name, tc := name, tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := tc.elementKeyValue.Equal(tc.other)

			if diff := cmp.Diff(tc.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted, +got): %s", diff)
			}
		})
	}
}
