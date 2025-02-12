// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestGetMetadataRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.GetMetadata_Request
		expected *tfprotov6.GetMetadataRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.GetMetadata_Request{},
			expected: &tfprotov6.GetMetadataRequest{},
		},
	}

	for name, testCase := range testCases {
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
		in       *tfplugin6.GetProviderSchema_Request
		expected *tfprotov6.GetProviderSchemaRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.GetProviderSchema_Request{},
			expected: &tfprotov6.GetProviderSchemaRequest{},
		},
	}

	for name, testCase := range testCases {
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
		in       *tfplugin6.ConfigureProvider_Request
		expected *tfprotov6.ConfigureProviderRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ConfigureProvider_Request{},
			expected: &tfprotov6.ConfigureProviderRequest{},
		},
		"Config": {
			in: &tfplugin6.ConfigureProvider_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ConfigureProviderRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TerraformVersion": {
			in: &tfplugin6.ConfigureProvider_Request{
				TerraformVersion: "0.0.1",
			},
			expected: &tfprotov6.ConfigureProviderRequest{
				TerraformVersion: "0.0.1",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.ConfigureProvider_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.ConfigureProviderRequest{
				ClientCapabilities: &tfprotov6.ConfigureProviderClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ConfigureProviderRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStopProviderRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.StopProvider_Request
		expected *tfprotov6.StopProviderRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.StopProvider_Request{},
			expected: &tfprotov6.StopProviderRequest{},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.StopProviderRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateProviderConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateProviderConfig_Request
		expected *tfprotov6.ValidateProviderConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateProviderConfig_Request{},
			expected: &tfprotov6.ValidateProviderConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateProviderConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateProviderConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateProviderConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
