package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateDataSourceConfig_Request(in tfprotov5.ValidateDataSourceConfigRequest) tfplugin5.ValidateDataSourceConfig_Request {
	resp := tfplugin5.ValidateDataSourceConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	return resp
}

func ValidateDataSourceConfig_Response(in tfprotov5.ValidateDataSourceConfigResponse) tfplugin5.ValidateDataSourceConfig_Response {
	return tfplugin5.ValidateDataSourceConfig_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ReadDataSource_Request(in tfprotov5.ReadDataSourceRequest) tfplugin5.ReadDataSource_Request {
	resp := tfplugin5.ReadDataSource_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	if in.ProviderMeta != nil {
		meta := Cty(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ReadDataSource_Response(in tfprotov5.ReadDataSourceResponse) tfplugin5.ReadDataSource_Response {
	resp := tfplugin5.ReadDataSource_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.State != nil {
		state := Cty(*in.State)
		resp.State = &state
	}
	return resp
}
