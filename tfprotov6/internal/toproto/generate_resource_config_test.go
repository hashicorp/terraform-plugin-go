// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
)

func TestGenerateResourceConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.GenerateResourceConfigResponse
		expected *tfplugin6.GenerateResourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.GenerateResourceConfigResponse{},
			expected: &tfplugin6.GenerateResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.GenerateResourceConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.GenerateResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Config": {
			in: &tfprotov6.GenerateResourceConfigResponse{
				Config: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.GenerateResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Config:      testTfplugin6DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GenerateResourceConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.GenerateResourceConfig_Response{},
				tfplugin6.Deferred{},
				timestamppb.Timestamp{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
