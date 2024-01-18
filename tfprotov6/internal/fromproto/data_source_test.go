// Copyright (c) HashiCorp, Inc.

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
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
				Config: testTfplugin6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				Config: testTfprotov6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
			},
		},
		"ProviderMeta": {
			in: &tfplugin6.ReadDataSource_Request{
				ProviderMeta: testTfplugin6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
			},
			expected: &tfprotov6.ReadDataSourceRequest{
				ProviderMeta: testTfprotov6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
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
				Config: testTfplugin6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
			},
			expected: &tfprotov6.ValidateDataResourceConfigRequest{
				Config: testTfprotov6DynamicValue(t,
					tftypes.Object{},
					tftypes.NewValue(tftypes.Object{}, nil),
				),
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
