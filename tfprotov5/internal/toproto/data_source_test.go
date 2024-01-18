// Copyright (c) HashiCorp, Inc.

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func TestGetMetadata_DataSourceMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.DataSourceMetadata
		expected *tfplugin5.GetMetadata_DataSourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.DataSourceMetadata{},
			expected: &tfplugin5.GetMetadata_DataSourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov5.DataSourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin5.GetMetadata_DataSourceMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_DataSourceMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.GetMetadata_DataSourceMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadDataSource_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ReadDataSourceResponse
		expected *tfplugin5.ReadDataSource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ReadDataSourceResponse{},
			expected: &tfplugin5.ReadDataSource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ReadDataSourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ReadDataSource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"State": {
			in: &tfprotov5.ReadDataSourceResponse{
				State: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.ReadDataSource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				State:       testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ReadDataSource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.ReadDataSource_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateDataSourceConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ValidateDataSourceConfigResponse
		expected *tfplugin5.ValidateDataSourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ValidateDataSourceConfigResponse{},
			expected: &tfplugin5.ValidateDataSourceConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ValidateDataSourceConfigResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ValidateDataSourceConfig_Response{
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

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ValidateDataSourceConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.ValidateDataSourceConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
