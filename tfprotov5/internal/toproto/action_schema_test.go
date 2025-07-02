// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func TestActionSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ActionSchema
		expected *tfplugin5.ActionSchema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"Schema": {
			in: &tfprotov5.ActionSchema{
				Schema: &tfprotov5.Schema{
					Block: &tfprotov5.SchemaBlock{
						Attributes: []*tfprotov5.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
				Type: tfprotov5.UnlinkedActionSchemaType{},
			},
			expected: &tfplugin5.ActionSchema{
				Schema: &tfplugin5.Schema{
					Block: &tfplugin5.Schema_Block{
						Attributes: []*tfplugin5.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin5.Schema_NestedBlock{},
					},
				},
				Type: &tfplugin5.ActionSchema_Unlinked_{
					Unlinked: &tfplugin5.ActionSchema_Unlinked{},
				},
			},
		},
		"Type - UnlinkedActionSchemaType": {
			in: &tfprotov5.ActionSchema{
				Type: tfprotov5.UnlinkedActionSchemaType{},
			},
			expected: &tfplugin5.ActionSchema{
				Type: &tfplugin5.ActionSchema_Unlinked_{
					Unlinked: &tfplugin5.ActionSchema_Unlinked{},
				},
			},
		},
		"Type - LifecycleActionSchemaType - Executes - Before": {
			in: &tfprotov5.ActionSchema{
				Type: tfprotov5.LifecycleActionSchemaType{
					Executes: tfprotov5.LifecycleExecutionOrderBefore,
				},
			},
			expected: &tfplugin5.ActionSchema{
				Type: &tfplugin5.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin5.ActionSchema_Lifecycle{
						Executes: tfplugin5.ActionSchema_Lifecycle_BEFORE,
					},
				},
			},
		},
		"Type - LifecycleActionSchemaType - Executes - After": {
			in: &tfprotov5.ActionSchema{
				Type: tfprotov5.LifecycleActionSchemaType{
					Executes: tfprotov5.LifecycleExecutionOrderAfter,
				},
			},
			expected: &tfplugin5.ActionSchema{
				Type: &tfplugin5.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin5.ActionSchema_Lifecycle{
						Executes: tfplugin5.ActionSchema_Lifecycle_AFTER,
					},
				},
			},
		},
		"Type - LifecycleActionSchemaType - LinkedResource": {
			in: &tfprotov5.ActionSchema{
				Type: tfprotov5.LifecycleActionSchemaType{
					LinkedResource: &tfprotov5.LinkedResourceSchema{
						TypeName:    "test",
						Description: "This is a test linked resource.",
					},
				},
			},
			expected: &tfplugin5.ActionSchema{
				Type: &tfplugin5.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin5.ActionSchema_Lifecycle{
						Executes: tfplugin5.ActionSchema_Lifecycle_INVALID,
						LinkedResource: &tfplugin5.ActionSchema_LinkedResource{
							TypeName:    "test",
							Description: "This is a test linked resource.",
						},
					},
				},
			},
		},
		"Type - LinkedActionSchemaType - LinkedResources": {
			in: &tfprotov5.ActionSchema{
				Type: tfprotov5.LinkedActionSchemaType{
					LinkedResources: []*tfprotov5.LinkedResourceSchema{
						{
							TypeName:    "test 1",
							Description: "This is a test linked resource.",
						},
						{
							TypeName:    "test 2",
							Description: "This is also a test linked resource.",
						},
					},
				},
			},
			expected: &tfplugin5.ActionSchema{
				Type: &tfplugin5.ActionSchema_Linked_{
					Linked: &tfplugin5.ActionSchema_Linked{
						LinkedResources: []*tfplugin5.ActionSchema_LinkedResource{
							{
								TypeName:    "test 1",
								Description: "This is a test linked resource.",
							},
							{
								TypeName:    "test 2",
								Description: "This is also a test linked resource.",
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ActionSchema(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ActionSchema{},
				tfplugin5.ActionSchema_Unlinked{},
				tfplugin5.ActionSchema_Unlinked_{},
				tfplugin5.ActionSchema_Lifecycle{},
				tfplugin5.ActionSchema_Lifecycle_{},
				tfplugin5.ActionSchema_Linked{},
				tfplugin5.ActionSchema_Linked_{},
				tfplugin5.ActionSchema_LinkedResource{},
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
