// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestConfigureProvider_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ConfigureProviderResponse
		expected *tfplugin5.Configure_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ConfigureProviderResponse{},
			expected: &tfplugin5.Configure_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ConfigureProviderResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.Configure_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Configure_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.Configure_Response{},
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
		in       *tfprotov5.GetMetadataResponse
		expected *tfplugin5.GetMetadata_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.GetMetadataResponse{},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:            []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"Actions": {
			in: &tfprotov5.GetMetadataResponse{
				Actions: []tfprotov5.ActionMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions: []*tfplugin5.GetMetadata_ActionMetadata{
					{
						TypeName: "test",
					},
				},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"DataSources": {
			in: &tfprotov5.GetMetadataResponse{
				DataSources: []tfprotov5.DataSourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions: []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources: []*tfplugin5.GetMetadata_DataSourceMetadata{
					{
						TypeName: "test",
					},
				},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.GetMetadataResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:     []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources: []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"EphemeralResources": {
			in: &tfprotov5.GetMetadataResponse{
				EphemeralResources: []tfprotov5.EphemeralResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:     []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources: []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics: []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{
					{
						TypeName: "test",
					},
				},
				ListResources: []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:     []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:     []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"Functions": {
			in: &tfprotov5.GetMetadataResponse{
				Functions: []tfprotov5.FunctionMetadata{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:            []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions: []*tfplugin5.GetMetadata_FunctionMetadata{
					{
						Name: "test",
					},
				},
				Resources: []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"ListResources": {
			in: &tfprotov5.GetMetadataResponse{
				ListResources: []tfprotov5.ListResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:            []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources: []*tfplugin5.GetMetadata_ListResourceMetadata{
					{
						TypeName: "test",
					},
				},
				Functions: []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources: []*tfplugin5.GetMetadata_ResourceMetadata{},
			},
		},
		"Resources": {
			in: &tfprotov5.GetMetadataResponse{
				Resources: []tfprotov5.ResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:            []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources: []*tfplugin5.GetMetadata_ResourceMetadata{
					{
						TypeName: "test",
					},
				},
			},
		},
		"ServerCapabilities": {
			in: &tfprotov5.GetMetadataResponse{
				ServerCapabilities: &tfprotov5.ServerCapabilities{
					PlanDestroy: true,
				},
			},
			expected: &tfplugin5.GetMetadata_Response{
				Actions:            []*tfplugin5.GetMetadata_ActionMetadata{},
				DataSources:        []*tfplugin5.GetMetadata_DataSourceMetadata{},
				Diagnostics:        []*tfplugin5.Diagnostic{},
				EphemeralResources: []*tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				ListResources:      []*tfplugin5.GetMetadata_ListResourceMetadata{},
				Functions:          []*tfplugin5.GetMetadata_FunctionMetadata{},
				Resources:          []*tfplugin5.GetMetadata_ResourceMetadata{},
				ServerCapabilities: &tfplugin5.ServerCapabilities{
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
				tfplugin5.Diagnostic{},
				tfplugin5.GetMetadata_ActionMetadata{},
				tfplugin5.GetMetadata_DataSourceMetadata{},
				tfplugin5.GetMetadata_EphemeralResourceMetadata{},
				tfplugin5.GetMetadata_FunctionMetadata{},
				tfplugin5.GetMetadata_ListResourceMetadata{},
				tfplugin5.GetMetadata_Response{},
				tfplugin5.GetMetadata_ResourceMetadata{},
				tfplugin5.ServerCapabilities{},
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
		in       *tfprotov5.GetProviderSchemaResponse
		expected *tfplugin5.GetProviderSchema_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.GetProviderSchemaResponse{},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas:          map[string]*tfplugin5.Schema{},
			},
		},
		"Actions": {
			in: &tfprotov5.GetProviderSchemaResponse{
				ActionSchemas: map[string]*tfprotov5.ActionSchema{
					"test": {
						Schema: &tfprotov5.Schema{
							Block: &tfprotov5.SchemaBlock{
								Attributes: []*tfprotov5.SchemaAttribute{
									{
										Name: "test",
									},
								},
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas: map[string]*tfplugin5.ActionSchema{
					"test": {
						Schema: &tfplugin5.Schema{
							Block: &tfplugin5.Schema_Block{
								Attributes: []*tfplugin5.Schema_Attribute{
									{
										Name: "test",
									},
								},
								BlockTypes: []*tfplugin5.Schema_NestedBlock{},
							},
						},
					},
				},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas:          map[string]*tfplugin5.Schema{},
			},
		},
		"DataSources": {
			in: &tfprotov5.GetProviderSchemaResponse{
				DataSourceSchemas: map[string]*tfprotov5.Schema{
					"test": {
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas: map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas: map[string]*tfplugin5.Schema{
					"test": {
						Block: &tfplugin5.Schema_Block{
							Attributes: []*tfplugin5.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin5.Schema_NestedBlock{},
						},
					},
				},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas:          map[string]*tfplugin5.Schema{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.GetProviderSchemaResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:     map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas: map[string]*tfplugin5.Schema{},
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas:          map[string]*tfplugin5.Schema{},
			},
		},
		"EphemeralResources": {
			in: &tfprotov5.GetProviderSchemaResponse{
				EphemeralResourceSchemas: map[string]*tfprotov5.Schema{
					"test": {
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:     map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas: map[string]*tfplugin5.Schema{},
				Diagnostics:       []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{
					"test": {
						Block: &tfplugin5.Schema_Block{
							Attributes: []*tfplugin5.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin5.Schema_NestedBlock{},
						},
					},
				},
				ListResourceSchemas: map[string]*tfplugin5.Schema{},
				Functions:           map[string]*tfplugin5.Function{},
				ResourceSchemas:     map[string]*tfplugin5.Schema{},
			},
		},
		"Functions": {
			in: &tfprotov5.GetProviderSchemaResponse{
				Functions: map[string]*tfprotov5.Function{
					"test": {
						Return: &tfprotov5.FunctionReturn{
							Type: tftypes.Bool,
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions: map[string]*tfplugin5.Function{
					"test": {
						Parameters: []*tfplugin5.Function_Parameter{},
						Return: &tfplugin5.Function_Return{
							Type: []byte(`"bool"`),
						},
					},
				},
				ResourceSchemas: map[string]*tfplugin5.Schema{},
			},
		},
		"ListResources": {
			in: &tfprotov5.GetProviderSchemaResponse{
				ListResourceSchemas: map[string]*tfprotov5.Schema{
					"test": {
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas: map[string]*tfplugin5.Schema{
					"test": {
						Block: &tfplugin5.Schema_Block{
							Attributes: []*tfplugin5.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin5.Schema_NestedBlock{},
						},
					},
				},
				Functions:       map[string]*tfplugin5.Function{},
				ResourceSchemas: map[string]*tfplugin5.Schema{},
			},
		},
		"Provider": {
			in: &tfprotov5.GetProviderSchemaResponse{
				Provider: &tfprotov5.Schema{
					Block: &tfprotov5.SchemaBlock{
						Attributes: []*tfprotov5.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				Provider: &tfplugin5.Schema{
					Block: &tfplugin5.Schema_Block{
						Attributes: []*tfplugin5.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin5.Schema_NestedBlock{},
					},
				},
				ResourceSchemas: map[string]*tfplugin5.Schema{},
			},
		},
		"ProviderMeta": {
			in: &tfprotov5.GetProviderSchemaResponse{
				ProviderMeta: &tfprotov5.Schema{
					Block: &tfprotov5.SchemaBlock{
						Attributes: []*tfprotov5.SchemaAttribute{
							{
								Name: "test",
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ProviderMeta: &tfplugin5.Schema{
					Block: &tfplugin5.Schema_Block{
						Attributes: []*tfplugin5.Schema_Attribute{
							{
								Name: "test",
							},
						},
						BlockTypes: []*tfplugin5.Schema_NestedBlock{},
					},
				},
				ResourceSchemas: map[string]*tfplugin5.Schema{},
			},
		},
		"Resources": {
			in: &tfprotov5.GetProviderSchemaResponse{
				ResourceSchemas: map[string]*tfprotov5.Schema{
					"test": {
						Block: &tfprotov5.SchemaBlock{
							Attributes: []*tfprotov5.SchemaAttribute{
								{
									Name: "test",
								},
							},
						},
					},
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas: map[string]*tfplugin5.Schema{
					"test": {
						Block: &tfplugin5.Schema_Block{
							Attributes: []*tfplugin5.Schema_Attribute{
								{
									Name: "test",
								},
							},
							BlockTypes: []*tfplugin5.Schema_NestedBlock{},
						},
					},
				},
			},
		},
		"ServerCapabilities": {
			in: &tfprotov5.GetProviderSchemaResponse{
				ServerCapabilities: &tfprotov5.ServerCapabilities{
					PlanDestroy: true,
				},
			},
			expected: &tfplugin5.GetProviderSchema_Response{
				ActionSchemas:            map[string]*tfplugin5.ActionSchema{},
				DataSourceSchemas:        map[string]*tfplugin5.Schema{},
				Diagnostics:              []*tfplugin5.Diagnostic{},
				EphemeralResourceSchemas: map[string]*tfplugin5.Schema{},
				ListResourceSchemas:      map[string]*tfplugin5.Schema{},
				Functions:                map[string]*tfplugin5.Function{},
				ResourceSchemas:          map[string]*tfplugin5.Schema{},
				ServerCapabilities: &tfplugin5.ServerCapabilities{
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
				tfplugin5.Diagnostic{},
				tfplugin5.ActionSchema{},
				tfplugin5.Function{},
				tfplugin5.Function_Return{},
				tfplugin5.GetProviderSchema_Response{},
				tfplugin5.Schema{},
				tfplugin5.Schema_Attribute{},
				tfplugin5.Schema_Block{},
				tfplugin5.ServerCapabilities{},
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
		in       *tfprotov5.GetResourceIdentitySchemasResponse
		expected *tfplugin5.GetResourceIdentitySchemas_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.GetResourceIdentitySchemasResponse{},
			expected: &tfplugin5.GetResourceIdentitySchemas_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				IdentitySchemas: map[string]*tfplugin5.ResourceIdentitySchema{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.GetResourceIdentitySchemasResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.GetResourceIdentitySchemas_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				IdentitySchemas: map[string]*tfplugin5.ResourceIdentitySchema{},
			},
		},
		"IdentitySchemas": {
			in: &tfprotov5.GetResourceIdentitySchemasResponse{
				IdentitySchemas: map[string]*tfprotov5.ResourceIdentitySchema{
					"test": {
						Version: 1,
						IdentityAttributes: []*tfprotov5.ResourceIdentitySchemaAttribute{
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
			expected: &tfplugin5.GetResourceIdentitySchemas_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				IdentitySchemas: map[string]*tfplugin5.ResourceIdentitySchema{
					"test": {
						Version: 1,
						IdentityAttributes: []*tfplugin5.ResourceIdentitySchema_IdentityAttribute{
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
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetResourceIdentitySchemas_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.GetResourceIdentitySchemas_Response{},
				tfplugin5.ResourceIdentitySchema{},
				tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPrepareProviderConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.PrepareProviderConfigResponse
		expected *tfplugin5.PrepareProviderConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.PrepareProviderConfigResponse{},
			expected: &tfplugin5.PrepareProviderConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.PrepareProviderConfigResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.PrepareProviderConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"PreparedConfig": {
			in: &tfprotov5.PrepareProviderConfigResponse{
				PreparedConfig: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.PrepareProviderConfig_Response{
				Diagnostics:    []*tfplugin5.Diagnostic{},
				PreparedConfig: testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.PrepareProviderConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.PrepareProviderConfig_Response{},
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
		in       *tfprotov5.StopProviderResponse
		expected *tfplugin5.Stop_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.StopProviderResponse{},
			expected: &tfplugin5.Stop_Response{},
		},
		"Error": {
			in: &tfprotov5.StopProviderResponse{
				Error: "test",
			},
			expected: &tfplugin5.Stop_Response{
				Error: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Stop_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.Stop_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
