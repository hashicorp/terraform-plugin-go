package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func GetProviderSchema_Request(in tfprotov5.GetProviderSchemaRequest) tfplugin5.GetProviderSchema_Request {
	return tfplugin5.GetProviderSchema_Request{}
}

func GetProviderSchema_Response(in tfprotov5.GetProviderSchemaResponse) tfplugin5.GetProviderSchema_Response {
	var resp tfplugin5.GetProviderSchema_Response
	if in.Provider != nil {
		schema := Schema(*in.Provider)
		resp.Provider = &schema
	}
	if in.ProviderMeta != nil {
		schema := Schema(*in.ProviderMeta)
		resp.ProviderMeta = &schema
	}
	resp.ResourceSchemas = make(map[string]*tfplugin5.Schema, len(in.ResourceSchemas))
	for k, v := range in.ResourceSchemas {
		if v == nil {
			resp.ResourceSchemas[k] = nil
			continue
		}
		schema := Schema(*v)
		resp.ResourceSchemas[k] = &schema
	}
	resp.DataSourceSchemas = make(map[string]*tfplugin5.Schema, len(in.DataSourceSchemas))
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

func PrepareProviderConfig_Request(in tfprotov5.PrepareProviderConfigRequest) tfplugin5.PrepareProviderConfig_Request {
	resp := tfplugin5.PrepareProviderConfig_Request{}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	return resp
}

func PrepareProviderConfig_Response(in tfprotov5.PrepareProviderConfigResponse) tfplugin5.PrepareProviderConfig_Response {
	resp := tfplugin5.PrepareProviderConfig_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.PreparedConfig != nil {
		config := Cty(*in.PreparedConfig)
		resp.PreparedConfig = &config
	}
	return resp
}

func Configure_Request(in tfprotov5.ConfigureProviderRequest) tfplugin5.Configure_Request {
	resp := tfplugin5.Configure_Request{
		TerraformVersion: in.TerraformVersion,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	return resp
}

func Configure_Response(in tfprotov5.ConfigureProviderResponse) tfplugin5.Configure_Response {
	return tfplugin5.Configure_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func Stop_Request(in tfprotov5.StopProviderRequest) tfplugin5.Stop_Request {
	return tfplugin5.Stop_Request{}
}

func Stop_Response(in tfprotov5.StopProviderResponse) tfplugin5.Stop_Response {
	return tfplugin5.Stop_Response{
		Error: in.Error,
	}
}
