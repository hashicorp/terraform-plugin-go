package tfprotov5

import (
	"context"

	tfproto "github.com/hashicorp/terraform-protocol-go"
)

var _ DataSourceServer = &BaseDataSourceServer{}

type BaseDataSourceServer struct{}

func (b *BaseDataSourceServer) ValidateDataSourceConfig(_ context.Context, _ *ValidateDataSourceConfigRequest) (*ValidateDataSourceConfigResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseDataSourceServer) ReadDataSource(_ context.Context, _ *ReadDataSourceRequest) (*ReadDataSourceResponse, error) {
	return nil, tfproto.ErrUnimplemented
}
