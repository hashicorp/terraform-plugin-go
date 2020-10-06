package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateDataSourceConfig_Request(in *tfprotov5.ValidateDataSourceConfigRequest) (*tfplugin5.ValidateDataSourceConfig_Request, error) {
	resp := &tfplugin5.ValidateDataSourceConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	return resp, nil
}

func ValidateDataSourceConfig_Response(in *tfprotov5.ValidateDataSourceConfigResponse) (*tfplugin5.ValidateDataSourceConfig_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	return &tfplugin5.ValidateDataSourceConfig_Response{
		Diagnostics: diags,
	}, nil
}

func ReadDataSource_Request(in *tfprotov5.ReadDataSourceRequest) (*tfplugin5.ReadDataSource_Request, error) {
	resp := &tfplugin5.ReadDataSource_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		resp.Config = DynamicValue(in.Config)
	}
	if in.ProviderMeta != nil {
		resp.ProviderMeta = DynamicValue(in.ProviderMeta)
	}
	return resp, nil
}

func ReadDataSource_Response(in *tfprotov5.ReadDataSourceResponse) (*tfplugin5.ReadDataSource_Response, error) {
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return nil, err
	}
	resp := &tfplugin5.ReadDataSource_Response{
		Diagnostics: diags,
	}
	if in.State != nil {
		resp.State = DynamicValue(in.State)
	}
	return resp, nil
}
