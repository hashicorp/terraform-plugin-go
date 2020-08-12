package tfprotov5

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

type DataSourceServer interface {
	ValidateDataSourceConfig(context.Context, *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error)
	ReadDataSource(context.Context, *ReadDataSourceRequest) (*ReadDataSourceResponse, error)
}

type ValidateDataSourceConfigRequest struct {
	TypeName string
	Config   *tftypes.RawValue
}

type ValidateDataSourceConfigResponse struct {
	Diagnostics []*Diagnostic
}

type ReadDataSourceRequest struct {
	TypeName     string
	Config       *tftypes.RawValue
	ProviderMeta *tftypes.RawValue
}

type ReadDataSourceResponse struct {
	State       *RawState
	Diagnostics []*Diagnostic
}
