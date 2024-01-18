package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestAttributePath(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tftypes.AttributePath
		expected *tfplugin6.AttributePath
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: tftypes.NewAttributePath(),
			expected: &tfplugin6.AttributePath{
				Steps: []*tfplugin6.AttributePath_Step{},
			},
		},
		"Steps": {
			in: tftypes.NewAttributePath().WithAttributeName("test"),
			expected: &tfplugin6.AttributePath{
				Steps: []*tfplugin6.AttributePath_Step{
					{
						Selector: &tfplugin6.AttributePath_Step_AttributeName{
							AttributeName: "test",
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.AttributePath(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestAttributePaths(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tftypes.AttributePath
		expected []*tfplugin6.AttributePath
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.AttributePath{},
		},
		"zero": {
			in:       []*tftypes.AttributePath{},
			expected: []*tfplugin6.AttributePath{},
		},
		"one": {
			in: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("test"),
			},
			expected: []*tfplugin6.AttributePath{
				{
					Steps: []*tfplugin6.AttributePath_Step{
						{
							Selector: &tfplugin6.AttributePath_Step_AttributeName{
								AttributeName: "test",
							},
						},
					},
				},
			},
		},
		"two": {
			in: []*tftypes.AttributePath{
				tftypes.NewAttributePath().WithAttributeName("test1"),
				tftypes.NewAttributePath().WithAttributeName("test2"),
			},
			expected: []*tfplugin6.AttributePath{
				{
					Steps: []*tfplugin6.AttributePath_Step{
						{
							Selector: &tfplugin6.AttributePath_Step_AttributeName{
								AttributeName: "test1",
							},
						},
					},
				},
				{
					Steps: []*tfplugin6.AttributePath_Step{
						{
							Selector: &tfplugin6.AttributePath_Step_AttributeName{
								AttributeName: "test2",
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.AttributePaths(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestAttributePath_Step(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       tftypes.AttributePathStep
		expected *tfplugin6.AttributePath_Step
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"AttributeName": {
			in: tftypes.AttributeName("test"),
			expected: &tfplugin6.AttributePath_Step{
				Selector: &tfplugin6.AttributePath_Step_AttributeName{
					AttributeName: "test",
				},
			},
		},
		"ElementKeyInt": {
			in: tftypes.ElementKeyInt(123),
			expected: &tfplugin6.AttributePath_Step{
				Selector: &tfplugin6.AttributePath_Step_ElementKeyInt{
					ElementKeyInt: 123,
				},
			},
		},
		"ElementKeyString": {
			in: tftypes.ElementKeyString("test"),
			expected: &tfplugin6.AttributePath_Step{
				Selector: &tfplugin6.AttributePath_Step_ElementKeyString{
					ElementKeyString: "test",
				},
			},
		},
		"ElementKeyValue": {
			in:       tftypes.ElementKeyValue(tftypes.NewValue(tftypes.String, "test")),
			expected: nil,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.AttributePath_Step(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath_Step{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestAttributePath_Steps(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []tftypes.AttributePathStep
		expected []*tfplugin6.AttributePath_Step
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.AttributePath_Step{},
		},
		"zero": {
			in:       []tftypes.AttributePathStep{},
			expected: []*tfplugin6.AttributePath_Step{},
		},
		"one": {
			in: []tftypes.AttributePathStep{
				tftypes.AttributeName("test"),
			},
			expected: []*tfplugin6.AttributePath_Step{
				{
					Selector: &tfplugin6.AttributePath_Step_AttributeName{
						AttributeName: "test",
					},
				},
			},
		},
		"two": {
			in: []tftypes.AttributePathStep{
				tftypes.AttributeName("test1"),
				tftypes.AttributeName("test2"),
			},
			expected: []*tfplugin6.AttributePath_Step{
				{
					Selector: &tfplugin6.AttributePath_Step_AttributeName{
						AttributeName: "test1",
					},
				},
				{
					Selector: &tfplugin6.AttributePath_Step_AttributeName{
						AttributeName: "test2",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.AttributePath_Steps(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath_Step{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
