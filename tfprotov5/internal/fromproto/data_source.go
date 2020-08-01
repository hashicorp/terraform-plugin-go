package fromproto

import (
	"github.com/hashicorp/terraform-protocol-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-protocol-go/tfprotov5"
)

func ValidateDataSourceConfigRequest(in tfplugin5.ValidateDataSourceConfig_Request) tfprotov5.ValidateDataSourceConfigRequest {
	return tfprotov5.ValidateDataSourceConfigRequest{
		TypeName: in.TypeName,
		Config:   nil, // TODO: unmarshal config from DynamicValue
	}
}

func ValidateDataSourceConfigResponse(in tfplugin5.ValidateDataSourceConfig_Response) tfprotov5.ValidateDataSourceConfigResponse {
	return tfprotov5.ValidateDataSourceConfigResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ReadDataSourceRequest(in tfplugin5.ReadDataSource_Request) tfprotov5.ReadDataSourceRequest {
	return tfprotov5.ReadDataSourceRequest{
		TypeName:     in.TypeName,
		Config:       nil, // TODO: unmarshal config from DynamicValue
		ProviderMeta: nil, // TODO: unmarshal provider_meta from DynamicValue
	}
}

func ReadDataSourceResponse(in tfplugin5.ReadDataSource_Response) tfprotov5.ReadDataSourceResponse {
	return tfprotov5.ReadDataSourceResponse{
		State:       nil, // TODO: figure out how to convert state appropriately
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}
