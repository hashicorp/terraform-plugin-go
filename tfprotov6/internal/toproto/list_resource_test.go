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

func TestGetMetadata_ListResourceMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ListResourceMetadata
		expected *tfplugin6.GetMetadata_ListResourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ListResourceMetadata{},
			expected: &tfplugin6.GetMetadata_ListResourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov6.ListResourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin6.GetMetadata_ListResourceMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_ListResourceMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.GetMetadata_ListResourceMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateListResourceConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ValidateListResourceConfigResponse
		expected *tfplugin6.ValidateListResourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ValidateListResourceConfigResponse{},
			expected: &tfplugin6.ValidateListResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ValidateListResourceConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ValidateListResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ValidateListResourceConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ValidateListResourceConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
