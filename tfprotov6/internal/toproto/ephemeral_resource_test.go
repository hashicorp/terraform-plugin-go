// Copyright (c) HashiCorp, Inc.
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

func TestGetMetadata_EphemeralResourceMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.EphemeralResourceMetadata
		expected *tfplugin6.GetMetadata_EphemeralResourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.EphemeralResourceMetadata{},
			expected: &tfplugin6.GetMetadata_EphemeralResourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov6.EphemeralResourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin6.GetMetadata_EphemeralResourceMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_EphemeralResourceMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.GetMetadata_EphemeralResourceMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateEphemeralResourceConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ValidateEphemeralResourceConfigResponse
		expected *tfplugin6.ValidateEphemeralResourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ValidateEphemeralResourceConfigResponse{},
			expected: &tfplugin6.ValidateEphemeralResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ValidateEphemeralResourceConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ValidateEphemeralResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ValidateEphemeralResourceConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ValidateEphemeralResourceConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestOpenEphemeralResource_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.OpenEphemeralResourceResponse
		expected *tfplugin6.OpenEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.OpenEphemeralResourceResponse{},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.OpenEphemeralResourceResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Result": {
			in: &tfprotov6.OpenEphemeralResourceResponse{
				Result: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Result:      testTfplugin6DynamicValue(),
			},
		},
		"Private": {
			in: &tfprotov6.OpenEphemeralResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"RenewAt": {
			in: &tfprotov6.OpenEphemeralResourceResponse{
				RenewAt: testGoTime(),
			},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				RenewAt:     testPbTimestamp(),
			},
		},
		"Deferred": {
			in: &tfprotov6.OpenEphemeralResourceResponse{
				Deferred: &tfprotov6.Deferred{
					Reason: tfprotov6.DeferredReasonResourceConfigUnknown,
				},
			},
			expected: &tfplugin6.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Deferred: &tfplugin6.Deferred{
					Reason: tfplugin6.Deferred_RESOURCE_CONFIG_UNKNOWN,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.OpenEphemeralResource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.OpenEphemeralResource_Response{},
				tfplugin6.Deferred{},
				timestamppb.Timestamp{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestRenewEphemeralResource_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.RenewEphemeralResourceResponse
		expected *tfplugin6.RenewEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.RenewEphemeralResourceResponse{},
			expected: &tfplugin6.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.RenewEphemeralResourceResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Private": {
			in: &tfprotov6.RenewEphemeralResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin6.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"RenewAt": {
			in: &tfprotov6.RenewEphemeralResourceResponse{
				RenewAt: testGoTime(),
			},
			expected: &tfplugin6.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				RenewAt:     testPbTimestamp(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.RenewEphemeralResource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.RenewEphemeralResource_Response{},
				timestamppb.Timestamp{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCloseEphemeralResource_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.CloseEphemeralResourceResponse
		expected *tfplugin6.CloseEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.CloseEphemeralResourceResponse{},
			expected: &tfplugin6.CloseEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.CloseEphemeralResourceResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.CloseEphemeralResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.CloseEphemeralResource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.CloseEphemeralResource_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
