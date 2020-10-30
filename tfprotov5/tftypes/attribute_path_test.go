package tftypes

import (
	"fmt"
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
		value    interface{}
		path     AttributePath
		expected interface{}
	}
	tests := map[string]testCase{
		"msi-root": {
			value: map[string]interface{}{
				"a": map[string]interface{}{
					"red":  true,
					"blue": 123,
				},
				"b": map[string]interface{}{
					"red":  false,
					"blue": 234,
				},
			},
			path: AttributePath{
				Steps: []AttributePathStep{
					AttributeName("a"),
				},
			},
			expected: map[string]interface{}{
				"red":  true,
				"blue": 123,
			},
		},
		"msi-full": {
			value: map[string]interface{}{
				"a": map[string]interface{}{
					"red":  true,
					"blue": 123,
				},
				"b": map[string]interface{}{
					"red":  false,
					"blue": 234,
				},
			},
			path: AttributePath{
				Steps: []AttributePathStep{
					AttributeName("a"),
					AttributeName("red"),
				},
			},
			expected: true,
		},
		"slice-interface-root": {
			value: []interface{}{
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
			path: AttributePath{
				Steps: []AttributePathStep{
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
			value: []interface{}{
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
			path: AttributePath{
				Steps: []AttributePathStep{
					ElementKeyInt(1),
					AttributeName("c"),
					ElementKeyInt(0),
				},
			},
			expected: "hello world",
		},
		"attributepathstepper": {
			value: []interface{}{
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
			path: AttributePath{
				Steps: []AttributePathStep{
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
			result, remaining, err := WalkAttributePath(test.value, test.path)
			if err != nil {
				t.Fatalf("error walking attribute path, %v still remains in the path: %s", remaining, err)
			}
			if diff := cmp.Diff(test.expected, result, cmp.Comparer(numberComparer), ValueComparer()); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
