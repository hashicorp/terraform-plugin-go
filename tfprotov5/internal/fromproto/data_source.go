package fromproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateDataSourceConfigRequest(in tfplugin5.ValidateDataSourceConfig_Request) (tfprotov5.ValidateDataSourceConfigRequest, error) {
	resp := tfprotov5.ValidateDataSourceConfigRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := TerraformTypesRawValue(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func ValidateDataSourceConfigResponse(in tfplugin5.ValidateDataSourceConfig_Response) (tfprotov5.ValidateDataSourceConfigResponse, error) {
	var resp tfprotov5.ValidateDataSourceConfigResponse
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return tfprotov5.ValidateDataSourceConfigResponse{}, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ReadDataSourceRequest(in tfplugin5.ReadDataSource_Request) (tfprotov5.ReadDataSourceRequest, error) {
	resp := tfprotov5.ReadDataSourceRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := TerraformTypesRawValue(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	if in.ProviderMeta != nil {
		meta, err := TerraformTypesRawValue(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ReadDataSourceResponse(in tfplugin5.ReadDataSource_Response) (tfprotov5.ReadDataSourceResponse, error) {
	var resp tfprotov5.ReadDataSourceResponse
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.State != nil {
		state, err := TerraformTypesRawValue(*in.State)
		if err != nil {
			return resp, fmt.Errorf("Error converting state: %w", err)
		}
		resp.State = &state
	}
	return resp, nil
}
