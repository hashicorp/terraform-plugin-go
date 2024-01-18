// Copyright (c) HashiCorp, Inc.

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.Schema
		expected *tfplugin5.Schema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.Schema{},
			expected: &tfplugin5.Schema{},
		},
		"Block": {
			in: &tfprotov5.Schema{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test",
						},
					},
				},
			},
			expected: &tfplugin5.Schema{
				Block: &tfplugin5.Schema_Block{
					Attributes: []*tfplugin5.Schema_Attribute{
						{
							Name: "test",
						},
					},
					BlockTypes: []*tfplugin5.Schema_NestedBlock{},
				},
			},
		},
		"Version": {
			in: &tfprotov5.Schema{
				Version: 123,
			},
			expected: &tfplugin5.Schema{
				Version: 123,
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
			got, _ := toproto.Schema(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema{},
				tfplugin5.Schema_Attribute{},
				tfplugin5.Schema_Block{},
				tfplugin5.Schema_NestedBlock{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_Attribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.SchemaAttribute
		expected *tfplugin5.Schema_Attribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.SchemaAttribute{},
			expected: &tfplugin5.Schema_Attribute{},
		},
		"Computed": {
			in: &tfprotov5.SchemaAttribute{
				Computed: true,
			},
			expected: &tfplugin5.Schema_Attribute{
				Computed: true,
			},
		},
		"Deprecated": {
			in: &tfprotov5.SchemaAttribute{
				Deprecated: true,
			},
			expected: &tfplugin5.Schema_Attribute{
				Deprecated: true,
			},
		},
		"Description": {
			in: &tfprotov5.SchemaAttribute{
				Description: "test",
			},
			expected: &tfplugin5.Schema_Attribute{
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.SchemaAttribute{
				DescriptionKind: tfprotov5.StringKindMarkdown,
			},
			expected: &tfplugin5.Schema_Attribute{
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
			},
		},
		"Name": {
			in: &tfprotov5.SchemaAttribute{
				Name: "test",
			},
			expected: &tfplugin5.Schema_Attribute{
				Name: "test",
			},
		},
		"Optional": {
			in: &tfprotov5.SchemaAttribute{
				Optional: true,
			},
			expected: &tfplugin5.Schema_Attribute{
				Optional: true,
			},
		},
		"Required": {
			in: &tfprotov5.SchemaAttribute{
				Required: true,
			},
			expected: &tfplugin5.Schema_Attribute{
				Required: true,
			},
		},
		"Sensitive": {
			in: &tfprotov5.SchemaAttribute{
				Sensitive: true,
			},
			expected: &tfplugin5.Schema_Attribute{
				Sensitive: true,
			},
		},
		"Type": {
			in: &tfprotov5.SchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Schema_Attribute{
				Type: []byte(`"bool"`),
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
			got, _ := toproto.Schema_Attribute(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema_Attribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_Attributes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tfprotov5.SchemaAttribute
		expected []*tfplugin5.Schema_Attribute
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin5.Schema_Attribute{},
		},
		"zero": {
			in:       []*tfprotov5.SchemaAttribute{},
			expected: []*tfplugin5.Schema_Attribute{},
		},
		"one": {
			in: []*tfprotov5.SchemaAttribute{
				{
					Name: "test",
				},
			},
			expected: []*tfplugin5.Schema_Attribute{
				{
					Name: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov5.SchemaAttribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
			expected: []*tfplugin5.Schema_Attribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
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
			got, _ := toproto.Schema_Attributes(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema_Attribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_Block(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.SchemaBlock
		expected *tfplugin5.Schema_Block
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.SchemaBlock{},
			expected: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{},
				BlockTypes: []*tfplugin5.Schema_NestedBlock{},
			},
		},
		"Attributes": {
			in: &tfprotov5.SchemaBlock{
				Attributes: []*tfprotov5.SchemaAttribute{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{
					{
						Name: "test",
					},
				},
				BlockTypes: []*tfplugin5.Schema_NestedBlock{},
			},
		},
		"BlockTypes": {
			in: &tfprotov5.SchemaBlock{
				BlockTypes: []*tfprotov5.SchemaNestedBlock{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{},
				BlockTypes: []*tfplugin5.Schema_NestedBlock{
					{
						TypeName: "test",
					},
				},
			},
		},
		"Deprecated": {
			in: &tfprotov5.SchemaBlock{
				Deprecated: true,
			},
			expected: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{},
				BlockTypes: []*tfplugin5.Schema_NestedBlock{},
				Deprecated: true,
			},
		},
		"Description": {
			in: &tfprotov5.SchemaBlock{
				Description: "test",
			},
			expected: &tfplugin5.Schema_Block{
				Attributes:  []*tfplugin5.Schema_Attribute{},
				BlockTypes:  []*tfplugin5.Schema_NestedBlock{},
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.SchemaBlock{
				DescriptionKind: tfprotov5.StringKindMarkdown,
			},
			expected: &tfplugin5.Schema_Block{
				Attributes:      []*tfplugin5.Schema_Attribute{},
				BlockTypes:      []*tfplugin5.Schema_NestedBlock{},
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
			},
		},
		"Version": {
			in: &tfprotov5.SchemaBlock{
				Version: 123,
			},
			expected: &tfplugin5.Schema_Block{
				Attributes: []*tfplugin5.Schema_Attribute{},
				BlockTypes: []*tfplugin5.Schema_NestedBlock{},
				Version:    123,
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
			got, _ := toproto.Schema_Block(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema_Attribute{},
				tfplugin5.Schema_Block{},
				tfplugin5.Schema_NestedBlock{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_NestedBlock(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.SchemaNestedBlock
		expected *tfplugin5.Schema_NestedBlock
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.SchemaNestedBlock{},
			expected: &tfplugin5.Schema_NestedBlock{},
		},
		"Block": {
			in: &tfprotov5.SchemaNestedBlock{
				Block: &tfprotov5.SchemaBlock{
					Attributes: []*tfprotov5.SchemaAttribute{
						{
							Name: "test",
						},
					},
				},
			},
			expected: &tfplugin5.Schema_NestedBlock{
				Block: &tfplugin5.Schema_Block{
					Attributes: []*tfplugin5.Schema_Attribute{
						{
							Name: "test",
						},
					},
					BlockTypes: []*tfplugin5.Schema_NestedBlock{},
				},
			},
		},
		"MaxItems": {
			in: &tfprotov5.SchemaNestedBlock{
				MaxItems: 123,
			},
			expected: &tfplugin5.Schema_NestedBlock{
				MaxItems: 123,
			},
		},
		"MinItems": {
			in: &tfprotov5.SchemaNestedBlock{
				MinItems: 123,
			},
			expected: &tfplugin5.Schema_NestedBlock{
				MinItems: 123,
			},
		},
		"Nesting": {
			in: &tfprotov5.SchemaNestedBlock{
				Nesting: tfprotov5.SchemaNestedBlockNestingModeList,
			},
			expected: &tfplugin5.Schema_NestedBlock{
				Nesting: tfplugin5.Schema_NestedBlock_LIST,
			},
		},
		"TypeName": {
			in: &tfprotov5.SchemaNestedBlock{
				TypeName: "test",
			},
			expected: &tfplugin5.Schema_NestedBlock{
				TypeName: "test",
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
			got, _ := toproto.Schema_NestedBlock(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema_Attribute{},
				tfplugin5.Schema_Block{},
				tfplugin5.Schema_NestedBlock{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_NestedBlocks(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tfprotov5.SchemaNestedBlock
		expected []*tfplugin5.Schema_NestedBlock
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin5.Schema_NestedBlock{},
		},
		"zero": {
			in:       []*tfprotov5.SchemaNestedBlock{},
			expected: []*tfplugin5.Schema_NestedBlock{},
		},
		"one": {
			in: []*tfprotov5.SchemaNestedBlock{
				{
					TypeName: "test",
				},
			},
			expected: []*tfplugin5.Schema_NestedBlock{
				{
					TypeName: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov5.SchemaNestedBlock{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
				},
			},
			expected: []*tfplugin5.Schema_NestedBlock{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
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
			got, _ := toproto.Schema_NestedBlocks(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Schema_NestedBlock{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
