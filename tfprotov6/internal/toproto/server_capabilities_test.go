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

func TestServerCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ServerCapabilities
		expected *tfplugin6.ServerCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ServerCapabilities{},
			expected: &tfplugin6.ServerCapabilities{},
		},
		"GetProviderSchemaOptional": {
			in: &tfprotov6.ServerCapabilities{
				GetProviderSchemaOptional: true,
			},
			expected: &tfplugin6.ServerCapabilities{
				GetProviderSchemaOptional: true,
			},
		},
		"MoveResourceState": {
			in: &tfprotov6.ServerCapabilities{
				MoveResourceState: true,
			},
			expected: &tfplugin6.ServerCapabilities{
				MoveResourceState: true,
			},
		},
		"PlanDestroy": {
			in: &tfprotov6.ServerCapabilities{
				PlanDestroy: true,
			},
			expected: &tfplugin6.ServerCapabilities{
				PlanDestroy: true,
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
				tfplugin6.ServerCapabilities{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
