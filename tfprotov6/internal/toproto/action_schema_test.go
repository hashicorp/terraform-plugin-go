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
)

func TestActionSchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ActionSchema
		expected *tfplugin6.ActionSchema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"Schema": {
			in: &tfprotov6.ActionSchema{
				Schema: &tfprotov6.Schema{
					Block: &tfprotov6.SchemaBlock{
						Attributes: []*tfprotov6.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
				Type: tfprotov6.UnlinkedActionSchemaType{},
			},
			expected: &tfplugin6.ActionSchema{
				Schema: &tfplugin6.Schema{
					Block: &tfplugin6.Schema_Block{
						Attributes: []*tfplugin6.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin6.Schema_NestedBlock{},
					},
				},
				Type: &tfplugin6.ActionSchema_Unlinked_{
					Unlinked: &tfplugin6.ActionSchema_Unlinked{},
				},
			},
		},
		"Type - UnlinkedActionSchemaType": {
			in: &tfprotov6.ActionSchema{
				Type: tfprotov6.UnlinkedActionSchemaType{},
			},
			expected: &tfplugin6.ActionSchema{
				Type: &tfplugin6.ActionSchema_Unlinked_{
					Unlinked: &tfplugin6.ActionSchema_Unlinked{},
				},
			},
		},
		"Type - LifecycleActionSchemaType - Executes - Before": {
			in: &tfprotov6.ActionSchema{
				Type: tfprotov6.LifecycleActionSchemaType{
					Executes: tfprotov6.LifecycleExecutionOrderBefore,
				},
			},
			expected: &tfplugin6.ActionSchema{
				Type: &tfplugin6.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin6.ActionSchema_Lifecycle{
						Executes: tfplugin6.ActionSchema_Lifecycle_BEFORE,
					},
				},
			},
		},
		"Type - LifecycleActionSchemaType - Executes - After": {
			in: &tfprotov6.ActionSchema{
				Type: tfprotov6.LifecycleActionSchemaType{
					Executes: tfprotov6.LifecycleExecutionOrderAfter,
				},
			},
			expected: &tfplugin6.ActionSchema{
				Type: &tfplugin6.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin6.ActionSchema_Lifecycle{
						Executes: tfplugin6.ActionSchema_Lifecycle_AFTER,
					},
				},
			},
		},
		"Type - LifecycleActionSchemaType - LinkedResource": {
			in: &tfprotov6.ActionSchema{
				Type: tfprotov6.LifecycleActionSchemaType{
					LinkedResource: &tfprotov6.LinkedResourceSchema{
						TypeName:    "test",
						Description: "This is a test linked resource.",
					},
				},
			},
			expected: &tfplugin6.ActionSchema{
				Type: &tfplugin6.ActionSchema_Lifecycle_{
					Lifecycle: &tfplugin6.ActionSchema_Lifecycle{
						Executes: tfplugin6.ActionSchema_Lifecycle_INVALID,
						LinkedResource: &tfplugin6.ActionSchema_LinkedResource{
							TypeName:    "test",
							Description: "This is a test linked resource.",
						},
					},
				},
			},
		},
		"Type - LinkedActionSchemaType - LinkedResources": {
			in: &tfprotov6.ActionSchema{
				Type: tfprotov6.LinkedActionSchemaType{
					LinkedResources: []*tfprotov6.LinkedResourceSchema{
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
			expected: &tfplugin6.ActionSchema{
				Type: &tfplugin6.ActionSchema_Linked_{
					Linked: &tfplugin6.ActionSchema_Linked{
						LinkedResources: []*tfplugin6.ActionSchema_LinkedResource{
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
				tfplugin6.ActionSchema{},
				tfplugin6.ActionSchema_Unlinked{},
				tfplugin6.ActionSchema_Unlinked_{},
				tfplugin6.ActionSchema_Lifecycle{},
				tfplugin6.ActionSchema_Lifecycle_{},
				tfplugin6.ActionSchema_Linked{},
				tfplugin6.ActionSchema_Linked_{},
				tfplugin6.ActionSchema_LinkedResource{},
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
