package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateDataSourceConfig_Request(in tfprotov5.ValidateDataSourceConfigRequest) (tfplugin5.ValidateDataSourceConfig_Request, error) {
	resp := tfplugin5.ValidateDataSourceConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func ValidateDataSourceConfig_Response(in tfprotov5.ValidateDataSourceConfigResponse) (tfplugin5.ValidateDataSourceConfig_Response, error) {
	var resp tfplugin5.ValidateDataSourceConfig_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ReadDataSource_Request(in tfprotov5.ReadDataSourceRequest) (tfplugin5.ReadDataSource_Request, error) {
	resp := tfplugin5.ReadDataSource_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	if in.ProviderMeta != nil {
		meta, err := Cty(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ReadDataSource_Response(in tfprotov5.ReadDataSourceResponse) (tfplugin5.ReadDataSource_Response, error) {
	var resp tfplugin5.ReadDataSource_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.State != nil {
		state, err := Cty(*in.State)
		if err != nil {
			return resp, fmt.Errorf("Error converting state: %w", err)
		}
		resp.State = &state
	}
	return resp, nil
}
