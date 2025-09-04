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
