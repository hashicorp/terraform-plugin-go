// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func TestGetMetadata_EphemeralResourceMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.EphemeralResourceMetadata
		expected *tfplugin5.GetMetadata_EphemeralResourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.EphemeralResourceMetadata{},
			expected: &tfplugin5.GetMetadata_EphemeralResourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov5.EphemeralResourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin5.GetMetadata_EphemeralResourceMetadata{
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
				tfplugin5.GetMetadata_EphemeralResourceMetadata{},
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
		in       *tfprotov5.ValidateEphemeralResourceConfigResponse
		expected *tfplugin5.ValidateEphemeralResourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ValidateEphemeralResourceConfigResponse{},
			expected: &tfplugin5.ValidateEphemeralResourceConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ValidateEphemeralResourceConfigResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ValidateEphemeralResourceConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
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
				tfplugin5.Diagnostic{},
				tfplugin5.ValidateEphemeralResourceConfig_Response{},
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
		in       *tfprotov5.OpenEphemeralResourceResponse
		expected *tfplugin5.OpenEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.OpenEphemeralResourceResponse{},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.OpenEphemeralResourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"Result": {
			in: &tfprotov5.OpenEphemeralResourceResponse{
				Result: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Result:      testTfplugin5DynamicValue(),
			},
		},
		"Private": {
			in: &tfprotov5.OpenEphemeralResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"RenewAt": {
			in: &tfprotov5.OpenEphemeralResourceResponse{
				RenewAt: testGoTime(),
			},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				RenewAt:     testPbTimestamp(),
			},
		},
		"Deferred": {
			in: &tfprotov5.OpenEphemeralResourceResponse{
				Deferred: &tfprotov5.Deferred{
					Reason: tfprotov5.DeferredReasonResourceConfigUnknown,
				},
			},
			expected: &tfplugin5.OpenEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Deferred: &tfplugin5.Deferred{
					Reason: tfplugin5.Deferred_RESOURCE_CONFIG_UNKNOWN,
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
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.OpenEphemeralResource_Response{},
				tfplugin5.Deferred{},
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
		in       *tfprotov5.RenewEphemeralResourceResponse
		expected *tfplugin5.RenewEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.RenewEphemeralResourceResponse{},
			expected: &tfplugin5.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.RenewEphemeralResourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"Private": {
			in: &tfprotov5.RenewEphemeralResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin5.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"RenewAt": {
			in: &tfprotov5.RenewEphemeralResourceResponse{
				RenewAt: testGoTime(),
			},
			expected: &tfplugin5.RenewEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
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
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.RenewEphemeralResource_Response{},
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
		in       *tfprotov5.CloseEphemeralResourceResponse
		expected *tfplugin5.CloseEphemeralResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.CloseEphemeralResourceResponse{},
			expected: &tfplugin5.CloseEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.CloseEphemeralResourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.CloseEphemeralResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
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
				tfplugin5.Diagnostic{},
				tfplugin5.CloseEphemeralResource_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
