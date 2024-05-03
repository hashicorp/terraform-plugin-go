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

func TestReadDataSourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ReadDataSource_Request
		expected *tfprotov6.ReadDataSourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ReadDataSource_Request{},
			expected: &tfprotov6.ReadDataSourceRequest{},
		},
		"Config": {
			in: &tfplugin6.ReadDataSource_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin6.ReadDataSource_Request{
				ProviderMeta: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				ProviderMeta: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ReadDataSource_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.ReadDataSource_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				ClientCapabilities: &tfprotov6.ReadDataSourceClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ReadDataSourceRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateDataResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateDataResourceConfig_Request
		expected *tfprotov6.ValidateDataResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateDataResourceConfig_Request{},
			expected: &tfprotov6.ValidateDataResourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateDataResourceConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateDataResourceConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ValidateDataResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ValidateDataResourceConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateDataResourceConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
