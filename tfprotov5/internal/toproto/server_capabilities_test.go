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

func TestServerCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ServerCapabilities
		expected *tfplugin5.ServerCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ServerCapabilities{},
			expected: &tfplugin5.ServerCapabilities{},
		},
		"GetProviderSchemaOptional": {
			in: &tfprotov5.ServerCapabilities{
				GetProviderSchemaOptional: true,
			},
			expected: &tfplugin5.ServerCapabilities{
				GetProviderSchemaOptional: true,
			},
		},
		"MoveResourceState": {
			in: &tfprotov5.ServerCapabilities{
				MoveResourceState: true,
			},
			expected: &tfplugin5.ServerCapabilities{
				MoveResourceState: true,
			},
		},
		"PlanDestroy": {
			in: &tfprotov5.ServerCapabilities{
				PlanDestroy: true,
			},
			expected: &tfplugin5.ServerCapabilities{
				PlanDestroy: true,
			},
		},
		"GenerateResourceConfig": {
			in: &tfprotov5.ServerCapabilities{
				GenerateResourceConfig: true,
			},
			expected: &tfplugin5.ServerCapabilities{
				GenerateResourceConfig: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ServerCapabilities(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ServerCapabilities{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
