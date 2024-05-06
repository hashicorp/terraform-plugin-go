// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestGetMetadataRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.GetMetadata_Request
		expected *tfprotov5.GetMetadataRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.GetMetadata_Request{},
			expected: &tfprotov5.GetMetadataRequest{},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.GetMetadataRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetProviderSchemaRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.GetProviderSchema_Request
		expected *tfprotov5.GetProviderSchemaRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.GetProviderSchema_Request{},
			expected: &tfprotov5.GetProviderSchemaRequest{},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.GetProviderSchemaRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestConfigureProviderRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.Configure_Request
		expected *tfprotov5.ConfigureProviderRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.Configure_Request{},
			expected: &tfprotov5.ConfigureProviderRequest{},
		},
		"Config": {
			in: &tfplugin5.Configure_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ConfigureProviderRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TerraformVersion": {
			in: &tfplugin5.Configure_Request{
				TerraformVersion: "0.0.1",
			},
			expected: &tfprotov5.ConfigureProviderRequest{
				TerraformVersion: "0.0.1",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.Configure_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov5.ConfigureProviderRequest{
				ClientCapabilities: &tfprotov5.ConfigureProviderClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ConfigureProviderRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPrepareProviderConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.PrepareProviderConfig_Request
		expected *tfprotov5.PrepareProviderConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.PrepareProviderConfig_Request{},
			expected: &tfprotov5.PrepareProviderConfigRequest{},
		},
		"Config": {
			in: &tfplugin5.PrepareProviderConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PrepareProviderConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.PrepareProviderConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStopProviderRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.Stop_Request
		expected *tfprotov5.StopProviderRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.Stop_Request{},
			expected: &tfprotov5.StopProviderRequest{},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.StopProviderRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
