package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func GetMetadata_Response(in *tfprotov5.GetMetadataResponse) (*tfplugin5.GetMetadata_Response, error) {
	if in == nil {
		return nil, nil
	}

	resp := &tfplugin5.GetMetadata_Response{
		DataSources:        make([]*tfplugin5.GetMetadata_DataSourceMetadata, 0, len(in.DataSources)),
		Functions:          make([]*tfplugin5.GetMetadata_FunctionMetadata, 0, len(in.Functions)),
		Resources:          make([]*tfplugin5.GetMetadata_ResourceMetadata, 0, len(in.Resources)),
		ServerCapabilities: ServerCapabilities(in.ServerCapabilities),
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return resp, err
	}

	resp.Diagnostics = diags

	for _, datasource := range in.DataSources {
		resp.DataSources = append(resp.DataSources, GetMetadata_DataSourceMetadata(&datasource))
	}

	for _, function := range in.Functions {
		resp.Functions = append(resp.Functions, GetMetadata_FunctionMetadata(&function))
	}

	for _, resource := range in.Resources {
		resp.Resources = append(resp.Resources, GetMetadata_ResourceMetadata(&resource))
	}

	return resp, nil
}

func GetProviderSchema_Response(in *tfprotov5.GetProviderSchemaResponse) (*tfplugin5.GetProviderSchema_Response, error) {
	if in == nil {
		return nil, nil
	}

	diagnostics, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	provider, err := Schema(in.Provider)

	if err != nil {
		return nil, fmt.Errorf("error marshaling provider schema: %w", err)
	}

	providerMeta, err := Schema(in.ProviderMeta)

	if err != nil {
		return nil, fmt.Errorf("error marshaling provider_meta schema: %w", err)
	}

	resp := &tfplugin5.GetProviderSchema_Response{
		DataSourceSchemas:  make(map[string]*tfplugin5.Schema, len(in.DataSourceSchemas)),
		Diagnostics:        diagnostics,
		Functions:          make(map[string]*tfplugin5.Function, len(in.Functions)),
		Provider:           provider,
		ProviderMeta:       providerMeta,
		ResourceSchemas:    make(map[string]*tfplugin5.Schema, len(in.ResourceSchemas)),
		ServerCapabilities: ServerCapabilities(in.ServerCapabilities),
	}

	for k, v := range in.ResourceSchemas {
		schema, err := Schema(v)

		if err != nil {
			return nil, fmt.Errorf("error marshaling resource schema for %q: %w", k, err)
		}

		resp.ResourceSchemas[k] = schema
	}

	for k, v := range in.DataSourceSchemas {
		schema, err := Schema(v)

		if err != nil {
			return nil, fmt.Errorf("error marshaling data source schema for %q: %w", k, err)
		}

		resp.DataSourceSchemas[k] = schema
	}

	for name, functionPtr := range in.Functions {
		function, err := Function(functionPtr)

		if err != nil {
			return nil, fmt.Errorf("error marshaling function definition for %q: %w", name, err)
		}

		resp.Functions[name] = function
	}

	return resp, nil
}

func PrepareProviderConfig_Response(in *tfprotov5.PrepareProviderConfigResponse) (*tfplugin5.PrepareProviderConfig_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.PrepareProviderConfig_Response{
		Diagnostics:    diags,
		PreparedConfig: DynamicValue(in.PreparedConfig),
	}

	return resp, nil
}

func Configure_Response(in *tfprotov5.ConfigureProviderResponse) (*tfplugin5.Configure_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.Configure_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func Stop_Response(in *tfprotov5.StopProviderResponse) *tfplugin5.Stop_Response {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.Stop_Response{
		Error: in.Error,
	}

	return resp
}
