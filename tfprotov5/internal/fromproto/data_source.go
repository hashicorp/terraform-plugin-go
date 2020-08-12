package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateDataSourceConfigRequest(in tfplugin5.ValidateDataSourceConfig_Request) tfprotov5.ValidateDataSourceConfigRequest {
	resp := tfprotov5.ValidateDataSourceConfigRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	return resp
}

func ValidateDataSourceConfigResponse(in tfplugin5.ValidateDataSourceConfig_Response) tfprotov5.ValidateDataSourceConfigResponse {
	return tfprotov5.ValidateDataSourceConfigResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ReadDataSourceRequest(in tfplugin5.ReadDataSource_Request) tfprotov5.ReadDataSourceRequest {
	resp := tfprotov5.ReadDataSourceRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	if in.ProviderMeta != nil {
		meta := TerraformTypesRawValue(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ReadDataSourceResponse(in tfplugin5.ReadDataSource_Response) tfprotov5.ReadDataSourceResponse {
	resp := tfprotov5.ReadDataSourceResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.State != nil {
		state := TerraformTypesRawValue(*in.State)
		resp.State = &state
	}
	return resp
}
