// Copyright (c) HashiCorp, Inc.

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestReadDataSourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ReadDataSource_Request
		expected *tfprotov5.ReadDataSourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ReadDataSource_Request{},
			expected: &tfprotov5.ReadDataSourceRequest{},
		},
		"Config": {
			in: &tfplugin5.ReadDataSource_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ReadDataSourceRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin5.ReadDataSource_Request{
				ProviderMeta: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ReadDataSourceRequest{
				ProviderMeta: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ReadDataSource_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ReadDataSourceRequest{
				TypeName: "test",
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

func TestValidateDataSourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ValidateDataSourceConfig_Request
		expected *tfprotov5.ValidateDataSourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ValidateDataSourceConfig_Request{},
			expected: &tfprotov5.ValidateDataSourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin5.ValidateDataSourceConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateDataSourceConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ValidateDataSourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ValidateDataSourceConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateDataSourceConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
