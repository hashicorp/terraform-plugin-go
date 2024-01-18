// Copyright (c) HashiCorp, Inc.

package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func GetMetadata_Response(in *tfprotov6.GetMetadataResponse) (*tfplugin6.GetMetadata_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.GetMetadata_Response{
		DataSources:        make([]*tfplugin6.GetMetadata_DataSourceMetadata, 0, len(in.DataSources)),
		Diagnostics:        diags,
		Functions:          make([]*tfplugin6.GetMetadata_FunctionMetadata, 0, len(in.Functions)),
		Resources:          make([]*tfplugin6.GetMetadata_ResourceMetadata, 0, len(in.Resources)),
		ServerCapabilities: ServerCapabilities(in.ServerCapabilities),
	}

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

func GetProviderSchema_Response(in *tfprotov6.GetProviderSchemaResponse) (*tfplugin6.GetProviderSchema_Response, error) {
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

	resp := &tfplugin6.GetProviderSchema_Response{
		DataSourceSchemas:  make(map[string]*tfplugin6.Schema, len(in.DataSourceSchemas)),
		Diagnostics:        diagnostics,
		Functions:          make(map[string]*tfplugin6.Function, len(in.Functions)),
		Provider:           provider,
		ProviderMeta:       providerMeta,
		ResourceSchemas:    make(map[string]*tfplugin6.Schema, len(in.ResourceSchemas)),
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

func ValidateProviderConfig_Response(in *tfprotov6.ValidateProviderConfigResponse) (*tfplugin6.ValidateProviderConfig_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ValidateProviderConfig_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func ConfigureProvider_Response(in *tfprotov6.ConfigureProviderResponse) (*tfplugin6.ConfigureProvider_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ConfigureProvider_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func StopProvider_Response(in *tfprotov6.StopProviderResponse) *tfplugin6.StopProvider_Response {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.StopProvider_Response{
		Error: in.Error,
	}

	return resp
}
