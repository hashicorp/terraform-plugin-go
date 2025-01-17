// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.Schema
		expected *tfplugin6.Schema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.Schema{},
			expected: &tfplugin6.Schema{},
		},
		"Block": {
			in: &tfprotov6.Schema{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test",
						},
					},
				},
			},
			expected: &tfplugin6.Schema{
				Block: &tfplugin6.Schema_Block{
					Attributes: []*tfplugin6.Schema_Attribute{
						{
							Name: "test",
						},
					},
					BlockTypes: []*tfplugin6.Schema_NestedBlock{},
				},
			},
		},
		"Version": {
			in: &tfprotov6.Schema{
				Version: 123,
			},
			expected: &tfplugin6.Schema{
				Version: 123,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Schema(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema{},
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Block{},
				tfplugin6.Schema_NestedBlock{},
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
		in       *tfprotov6.SchemaAttribute
		expected *tfplugin6.Schema_Attribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.SchemaAttribute{},
			expected: &tfplugin6.Schema_Attribute{},
		},
		"Computed": {
			in: &tfprotov6.SchemaAttribute{
				Computed: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				Computed: true,
			},
		},
		"Deprecated": {
			in: &tfprotov6.SchemaAttribute{
				Deprecated: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				Deprecated: true,
			},
		},
		"Description": {
			in: &tfprotov6.SchemaAttribute{
				Description: "test",
			},
			expected: &tfplugin6.Schema_Attribute{
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.SchemaAttribute{
				DescriptionKind: tfprotov6.StringKindMarkdown,
			},
			expected: &tfplugin6.Schema_Attribute{
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
			},
		},
		"Name": {
			in: &tfprotov6.SchemaAttribute{
				Name: "test",
			},
			expected: &tfplugin6.Schema_Attribute{
				Name: "test",
			},
		},
		"NestedType": {
			in: &tfprotov6.SchemaAttribute{
				NestedType: &tfprotov6.SchemaObject{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test",
						},
					},
					Nesting: tfprotov6.SchemaObjectNestingModeList,
				},
			},
			expected: &tfplugin6.Schema_Attribute{
				NestedType: &tfplugin6.Schema_Object{
					Attributes: []*tfplugin6.Schema_Attribute{
						{
							Name: "test",
						},
					},
					Nesting: tfplugin6.Schema_Object_LIST,
				},
			},
		},
		"Optional": {
			in: &tfprotov6.SchemaAttribute{
				Optional: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				Optional: true,
			},
		},
		"Required": {
			in: &tfprotov6.SchemaAttribute{
				Required: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				Required: true,
			},
		},
		"Sensitive": {
			in: &tfprotov6.SchemaAttribute{
				Sensitive: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				Sensitive: true,
			},
		},
		"Type": {
			in: &tfprotov6.SchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Schema_Attribute{
				Type: []byte(`"bool"`),
			},
		},
		"WriteOnly": {
			in: &tfprotov6.SchemaAttribute{
				WriteOnly: true,
			},
			expected: &tfplugin6.Schema_Attribute{
				WriteOnly: true,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Schema_Attribute(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Object{},
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
		in       []*tfprotov6.SchemaAttribute
		expected []*tfplugin6.Schema_Attribute
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.Schema_Attribute{},
		},
		"zero": {
			in:       []*tfprotov6.SchemaAttribute{},
			expected: []*tfplugin6.Schema_Attribute{},
		},
		"one": {
			in: []*tfprotov6.SchemaAttribute{
				{
					Name: "test",
				},
			},
			expected: []*tfplugin6.Schema_Attribute{
				{
					Name: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov6.SchemaAttribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
			expected: []*tfplugin6.Schema_Attribute{
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

			got := toproto.Schema_Attributes(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_Attribute{},
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
		in       *tfprotov6.SchemaBlock
		expected *tfplugin6.Schema_Block
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.SchemaBlock{},
			expected: &tfplugin6.Schema_Block{
				Attributes: []*tfplugin6.Schema_Attribute{},
				BlockTypes: []*tfplugin6.Schema_NestedBlock{},
			},
		},
		"Attributes": {
			in: &tfprotov6.SchemaBlock{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin6.Schema_Block{
				Attributes: []*tfplugin6.Schema_Attribute{
					{
						Name: "test",
					},
				},
				BlockTypes: []*tfplugin6.Schema_NestedBlock{},
			},
		},
		"BlockTypes": {
			in: &tfprotov6.SchemaBlock{
				BlockTypes: []*tfprotov6.SchemaNestedBlock{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin6.Schema_Block{
				Attributes: []*tfplugin6.Schema_Attribute{},
				BlockTypes: []*tfplugin6.Schema_NestedBlock{
					{
						TypeName: "test",
					},
				},
			},
		},
		"Deprecated": {
			in: &tfprotov6.SchemaBlock{
				Deprecated: true,
			},
			expected: &tfplugin6.Schema_Block{
				Attributes: []*tfplugin6.Schema_Attribute{},
				BlockTypes: []*tfplugin6.Schema_NestedBlock{},
				Deprecated: true,
			},
		},
		"Description": {
			in: &tfprotov6.SchemaBlock{
				Description: "test",
			},
			expected: &tfplugin6.Schema_Block{
				Attributes:  []*tfplugin6.Schema_Attribute{},
				BlockTypes:  []*tfplugin6.Schema_NestedBlock{},
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.SchemaBlock{
				DescriptionKind: tfprotov6.StringKindMarkdown,
			},
			expected: &tfplugin6.Schema_Block{
				Attributes:      []*tfplugin6.Schema_Attribute{},
				BlockTypes:      []*tfplugin6.Schema_NestedBlock{},
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
			},
		},
		"Version": {
			in: &tfprotov6.SchemaBlock{
				Version: 123,
			},
			expected: &tfplugin6.Schema_Block{
				Attributes: []*tfplugin6.Schema_Attribute{},
				BlockTypes: []*tfplugin6.Schema_NestedBlock{},
				Version:    123,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Schema_Block(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Block{},
				tfplugin6.Schema_NestedBlock{},
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
		in       *tfprotov6.SchemaNestedBlock
		expected *tfplugin6.Schema_NestedBlock
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.SchemaNestedBlock{},
			expected: &tfplugin6.Schema_NestedBlock{},
		},
		"Block": {
			in: &tfprotov6.SchemaNestedBlock{
				Block: &tfprotov6.SchemaBlock{
					Attributes: []*tfprotov6.SchemaAttribute{
						{
							Name: "test",
						},
					},
				},
			},
			expected: &tfplugin6.Schema_NestedBlock{
				Block: &tfplugin6.Schema_Block{
					Attributes: []*tfplugin6.Schema_Attribute{
						{
							Name: "test",
						},
					},
					BlockTypes: []*tfplugin6.Schema_NestedBlock{},
				},
			},
		},
		"MaxItems": {
			in: &tfprotov6.SchemaNestedBlock{
				MaxItems: 123,
			},
			expected: &tfplugin6.Schema_NestedBlock{
				MaxItems: 123,
			},
		},
		"MinItems": {
			in: &tfprotov6.SchemaNestedBlock{
				MinItems: 123,
			},
			expected: &tfplugin6.Schema_NestedBlock{
				MinItems: 123,
			},
		},
		"Nesting": {
			in: &tfprotov6.SchemaNestedBlock{
				Nesting: tfprotov6.SchemaNestedBlockNestingModeList,
			},
			expected: &tfplugin6.Schema_NestedBlock{
				Nesting: tfplugin6.Schema_NestedBlock_LIST,
			},
		},
		"TypeName": {
			in: &tfprotov6.SchemaNestedBlock{
				TypeName: "test",
			},
			expected: &tfplugin6.Schema_NestedBlock{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Schema_NestedBlock(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Block{},
				tfplugin6.Schema_NestedBlock{},
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
		in       []*tfprotov6.SchemaNestedBlock
		expected []*tfplugin6.Schema_NestedBlock
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.Schema_NestedBlock{},
		},
		"zero": {
			in:       []*tfprotov6.SchemaNestedBlock{},
			expected: []*tfplugin6.Schema_NestedBlock{},
		},
		"one": {
			in: []*tfprotov6.SchemaNestedBlock{
				{
					TypeName: "test",
				},
			},
			expected: []*tfplugin6.Schema_NestedBlock{
				{
					TypeName: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov6.SchemaNestedBlock{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
				},
			},
			expected: []*tfplugin6.Schema_NestedBlock{
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

			got := toproto.Schema_NestedBlocks(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_NestedBlock{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestSchema_Object(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.SchemaObject
		expected *tfplugin6.Schema_Object
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.SchemaObject{},
			expected: &tfplugin6.Schema_Object{
				Attributes: []*tfplugin6.Schema_Attribute{},
			},
		},
		"Attributes": {
			in: &tfprotov6.SchemaObject{
				Attributes: []*tfprotov6.SchemaAttribute{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin6.Schema_Object{
				Attributes: []*tfplugin6.Schema_Attribute{
					{
						Name: "test",
					},
				},
			},
		},
		"Nesting": {
			in: &tfprotov6.SchemaObject{
				Nesting: tfprotov6.SchemaObjectNestingModeList,
			},
			expected: &tfplugin6.Schema_Object{
				Attributes: []*tfplugin6.Schema_Attribute{},
				Nesting:    tfplugin6.Schema_Object_LIST,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Schema_Object(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Object{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
