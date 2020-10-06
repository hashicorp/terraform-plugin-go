package tfprotov5

import (
	"context"
)

type DataSourceServer interface {
	ValidateDataSourceConfig(context.Context, *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error)
	ReadDataSource(context.Context, *ReadDataSourceRequest) (*ReadDataSourceResponse, error)
}

type ValidateDataSourceConfigRequest struct {
	TypeName string
	Config   *DynamicValue
}

type ValidateDataSourceConfigResponse struct {
	Diagnostics []*Diagnostic
}

type ReadDataSourceRequest struct {
	TypeName     string
	Config       *DynamicValue
	ProviderMeta *DynamicValue
}

type ReadDataSourceResponse struct {
	State       *DynamicValue
	Diagnostics []*Diagnostic
}
