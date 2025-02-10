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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfigureProvider_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ConfigureProviderResponse
		expected *tfplugin6.ConfigureProvider_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ConfigureProviderResponse{},
			expected: &tfplugin6.ConfigureProvider_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ConfigureProviderResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ConfigureProvider_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ConfigureProvider_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ConfigureProvider_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetMetadata_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.GetMetadataResponse
		expected *tfplugin6.GetMetadata_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.GetMetadataResponse{},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources:        []*tfplugin6.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin6.Diagnostic{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Functions:          []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin6.GetMetadata_ResourceMetadata{},
			},
		},
		"DataSources": {
			in: &tfprotov6.GetMetadataResponse{
				DataSources: []tfprotov6.DataSourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources: []*tfplugin6.GetMetadata_DataSourceMetadata{
					{
						TypeName: "test",
					},
				},
				Diagnostics:        []*tfplugin6.Diagnostic{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Functions:          []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin6.GetMetadata_ResourceMetadata{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.GetMetadataResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources: []*tfplugin6.GetMetadata_DataSourceMetadata{},
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Functions:          []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin6.GetMetadata_ResourceMetadata{},
			},
		},
		"EphemeralResources": {
			in: &tfprotov6.GetMetadataResponse{
				EphemeralResources: []tfprotov6.EphemeralResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources: []*tfplugin6.GetMetadata_DataSourceMetadata{},
				Diagnostics: []*tfplugin6.Diagnostic{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{
					{
						TypeName: "test",
					},
				},
				Functions: []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources: []*tfplugin6.GetMetadata_ResourceMetadata{},
			},
		},
		"Functions": {
			in: &tfprotov6.GetMetadataResponse{
				Functions: []tfprotov6.FunctionMetadata{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources:        []*tfplugin6.GetMetadata_DataSourceMetadata{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Diagnostics:        []*tfplugin6.Diagnostic{},
				Functions: []*tfplugin6.GetMetadata_FunctionMetadata{
					{
						Name: "test",
					},
				},
				Resources: []*tfplugin6.GetMetadata_ResourceMetadata{},
			},
		},
		"Resources": {
			in: &tfprotov6.GetMetadataResponse{
				Resources: []tfprotov6.ResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources:        []*tfplugin6.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin6.Diagnostic{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Functions:          []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources: []*tfplugin6.GetMetadata_ResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
		},
		"ServerCapabilities": {
			in: &tfprotov6.GetMetadataResponse{
				ServerCapabilities: &tfprotov6.ServerCapabilities{
					PlanDestroy: true,
				},
			},
			expected: &tfplugin6.GetMetadata_Response{
				DataSources:        []*tfplugin6.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin6.Diagnostic{},
				EphemeralResources: []*tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				Functions:          []*tfplugin6.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin6.GetMetadata_ResourceMetadata{},
				ServerCapabilities: &tfplugin6.ServerCapabilities{
					PlanDestroy: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.GetMetadata_DataSourceMetadata{},
				tfplugin6.GetMetadata_EphemeralResourceMetadata{},
				tfplugin6.GetMetadata_FunctionMetadata{},
				tfplugin6.GetMetadata_Response{},
				tfplugin6.GetMetadata_ResourceMetadata{},
				tfplugin6.ServerCapabilities{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetProviderSchema_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.GetProviderSchemaResponse
		expected *tfplugin6.GetProviderSchema_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.GetProviderSchemaResponse{},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ResourceSchemas:          map[string]*tfplugin6.Schema{},
			},
		},
		"DataSources": {
			in: &tfprotov6.GetProviderSchemaResponse{
				DataSourceSchemas: map[string]*tfprotov6.Schema{
					"test": {
						Block: &tfprotov6.SchemaBlock{
							Attributes: []*tfprotov6.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas: map[string]*tfplugin6.Schema{
					"test": {
						Block: &tfplugin6.Schema_Block{
							Attributes: []*tfplugin6.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin6.Schema_NestedBlock{},
						},
					},
				},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ResourceSchemas:          map[string]*tfplugin6.Schema{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.GetProviderSchemaResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas: map[string]*tfplugin6.Schema{},
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ResourceSchemas:          map[string]*tfplugin6.Schema{},
			},
		},
		"EphemeralResources": {
			in: &tfprotov6.GetProviderSchemaResponse{
				EphemeralResourceSchemas: map[string]*tfprotov6.Schema{
					"test": {
						Block: &tfprotov6.SchemaBlock{
							Attributes: []*tfprotov6.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas: map[string]*tfplugin6.Schema{},
				Diagnostics:       []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{
					"test": {
						Block: &tfplugin6.Schema_Block{
							Attributes: []*tfplugin6.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin6.Schema_NestedBlock{},
						},
					},
				},
				Functions:       map[string]*tfplugin6.Function{},
				ResourceSchemas: map[string]*tfplugin6.Schema{},
			},
		},
		"Functions": {
			in: &tfprotov6.GetProviderSchemaResponse{
				Functions: map[string]*tfprotov6.Function{
					"test": {
						Return: &tfprotov6.FunctionReturn{
							Type: tftypes.Bool,
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions: map[string]*tfplugin6.Function{
					"test": {
						Parameters: []*tfplugin6.Function_Parameter{},
						Return: &tfplugin6.Function_Return{
							Type: []byte(`"bool"`),
						},
					},
				},
				ResourceSchemas: map[string]*tfplugin6.Schema{},
			},
		},
		"Provider": {
			in: &tfprotov6.GetProviderSchemaResponse{
				Provider: &tfprotov6.Schema{
					Block: &tfprotov6.SchemaBlock{
						Attributes: []*tfprotov6.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				Functions:                map[string]*tfplugin6.Function{},
				Provider: &tfplugin6.Schema{
					Block: &tfplugin6.Schema_Block{
						Attributes: []*tfplugin6.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin6.Schema_NestedBlock{},
					},
				},
				ResourceSchemas: map[string]*tfplugin6.Schema{},
			},
		},
		"ProviderMeta": {
			in: &tfprotov6.GetProviderSchemaResponse{
				ProviderMeta: &tfprotov6.Schema{
					Block: &tfprotov6.SchemaBlock{
						Attributes: []*tfprotov6.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ProviderMeta: &tfplugin6.Schema{
					Block: &tfplugin6.Schema_Block{
						Attributes: []*tfplugin6.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin6.Schema_NestedBlock{},
					},
				},
				ResourceSchemas: map[string]*tfplugin6.Schema{},
			},
		},
		"Resources": {
			in: &tfprotov6.GetProviderSchemaResponse{
				ResourceSchemas: map[string]*tfprotov6.Schema{
					"test": {
						Block: &tfprotov6.SchemaBlock{
							Attributes: []*tfprotov6.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ResourceSchemas: map[string]*tfplugin6.Schema{
					"test": {
						Block: &tfplugin6.Schema_Block{
							Attributes: []*tfplugin6.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin6.Schema_NestedBlock{},
						},
					},
				},
			},
		},
		"ServerCapabilities": {
			in: &tfprotov6.GetProviderSchemaResponse{
				ServerCapabilities: &tfprotov6.ServerCapabilities{
					PlanDestroy: true,
				},
			},
			expected: &tfplugin6.GetProviderSchema_Response{
				DataSourceSchemas:        map[string]*tfplugin6.Schema{},
				Diagnostics:              []*tfplugin6.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin6.Schema{},
				Functions:                map[string]*tfplugin6.Function{},
				ResourceSchemas:          map[string]*tfplugin6.Schema{},
				ServerCapabilities: &tfplugin6.ServerCapabilities{
					PlanDestroy: true,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetProviderSchema_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.Function{},
				tfplugin6.Function_Return{},
				tfplugin6.GetProviderSchema_Response{},
				tfplugin6.Schema{},
				tfplugin6.Schema_Attribute{},
				tfplugin6.Schema_Block{},
				tfplugin6.ServerCapabilities{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetResourceIdentitySchemas_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.GetResourceIdentitySchemasResponse
		expected *tfplugin6.GetResourceIdentitySchemas_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.GetResourceIdentitySchemasResponse{},
			expected: &tfplugin6.GetResourceIdentitySchemas_Response{
				Diagnostics:     []*tfplugin6.Diagnostic{},
				IdentitySchemas: map[string]*tfplugin6.ResourceIdentitySchema{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.GetResourceIdentitySchemasResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.GetResourceIdentitySchemas_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				IdentitySchemas: map[string]*tfplugin6.ResourceIdentitySchema{},
			},
		},
		"IdentitySchemas": {
			in: &tfprotov6.GetResourceIdentitySchemasResponse{
				IdentitySchemas: map[string]*tfprotov6.ResourceIdentitySchema{
					"test": {
						Version: 1,
						IdentityAttributes: []*tfprotov6.ResourceIdentitySchemaAttribute{
							{
								Name:              "req",
								RequiredForImport: true,
								Description:       "this one's required",
							},
							{
								Name:              "opt",
								OptionalForImport: true,
								Description:       "this one's optional",
							},
						},
					},
				},
			},
			expected: &tfplugin6.GetResourceIdentitySchemas_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				IdentitySchemas: map[string]*tfplugin6.ResourceIdentitySchema{
					"test": {
						Version: 1,
						IdentityAttributes: []*tfplugin6.ResourceIdentitySchema_IdentityAttribute{
							{
								Name:              "req",
								RequiredForImport: true,
								Description:       "this one's required",
							},
							{
								Name:              "opt",
								OptionalForImport: true,
								Description:       "this one's optional",
							},
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetResourceIdentitySchemas_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.GetResourceIdentitySchemas_Response{},
				tfplugin6.ResourceIdentitySchema{},
				tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateProviderConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ValidateProviderConfigResponse
		expected *tfplugin6.ValidateProviderConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ValidateProviderConfigResponse{},
			expected: &tfplugin6.ValidateProviderConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ValidateProviderConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ValidateProviderConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ValidateProviderConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ValidateProviderConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestStopProvider_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.StopProviderResponse
		expected *tfplugin6.StopProvider_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.StopProviderResponse{},
			expected: &tfplugin6.StopProvider_Response{},
		},
		"Error": {
			in: &tfprotov6.StopProviderResponse{
				Error: "test",
			},
			expected: &tfplugin6.StopProvider_Response{
				Error: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.StopProvider_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.StopProvider_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
