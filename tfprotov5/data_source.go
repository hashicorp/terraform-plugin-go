package tfprotov5

import "context"

type DataSourceServer interface {
	ValidateDataSourceConfig(context.Context, *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error)
	ReadDataSource(context.Context, *ReadDataSourceRequest) (*ReadDataSourceResponse, error)
}

type ValidateDataSourceConfigRequest struct {
	TypeName string
	Config   *CtyBlock
}

type ValidateDataSourceConfigResponse struct {
	Diagnostics []*Diagnostic
}

type ReadDataSourceRequest struct {
	TypeName     string
	Config       *CtyBlock
	ProviderMeta *CtyBlock
}

type ReadDataSourceResponse struct {
	State       *CtyBlock // TODO: figure out how to represent state
	Diagnostics []*Diagnostic
}
