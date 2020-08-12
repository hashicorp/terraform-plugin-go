package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func GetProviderSchemaRequest(in tfplugin5.GetProviderSchema_Request) tfprotov5.GetProviderSchemaRequest {
	return tfprotov5.GetProviderSchemaRequest{}
}

func GetProviderSchemaResponse(in tfplugin5.GetProviderSchema_Response) tfprotov5.GetProviderSchemaResponse {
	var resp tfprotov5.GetProviderSchemaResponse
	if in.Provider != nil {
		schema := Schema(*in.Provider)
		resp.Provider = &schema
	}
	if in.ProviderMeta != nil {
		schema := Schema(*in.ProviderMeta)
		resp.ProviderMeta = &schema
	}
	resp.ResourceSchemas = make(map[string]*tfprotov5.Schema, len(in.ResourceSchemas))
	for k, v := range in.ResourceSchemas {
		if v == nil {
			resp.ResourceSchemas[k] = nil
			continue
		}
		schema := Schema(*v)
		resp.ResourceSchemas[k] = &schema
	}
	resp.DataSourceSchemas = make(map[string]*tfprotov5.Schema, len(in.DataSourceSchemas))
	for k, v := range in.DataSourceSchemas {
		if v == nil {
			resp.DataSourceSchemas[k] = nil
			continue
		}
		schema := Schema(*v)
		resp.DataSourceSchemas[k] = &schema
	}
	resp.Diagnostics = Diagnostics(in.Diagnostics)
	return resp
}

func PrepareProviderConfigRequest(in tfplugin5.PrepareProviderConfig_Request) tfprotov5.PrepareProviderConfigRequest {
	resp := tfprotov5.PrepareProviderConfigRequest{}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	return resp
}

func PrepareProviderConfigResponse(in tfplugin5.PrepareProviderConfig_Response) tfprotov5.PrepareProviderConfigResponse {
	resp := tfprotov5.PrepareProviderConfigResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.PreparedConfig != nil {
		config := TerraformTypesRawValue(*in.PreparedConfig)
		resp.PreparedConfig = &config
	}
	return resp
}

func ConfigureProviderRequest(in tfplugin5.Configure_Request) tfprotov5.ConfigureProviderRequest {
	resp := tfprotov5.ConfigureProviderRequest{
		TerraformVersion: in.TerraformVersion,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	return resp
}

func ConfigureProviderResponse(in tfplugin5.Configure_Response) tfprotov5.ConfigureProviderResponse {
	return tfprotov5.ConfigureProviderResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func StopProviderRequest(in tfplugin5.Stop_Request) tfprotov5.StopProviderRequest {
	return tfprotov5.StopProviderRequest{}
}

func StopProviderResponse(in tfplugin5.Stop_Response) tfprotov5.StopProviderResponse {
	return tfprotov5.StopProviderResponse{
		Error: in.Error,
	}
}
