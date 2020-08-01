package tfprotov5

import (
	"context"

	tfproto "github.com/hashicorp/terraform-protocol-go"
)

var _ ProviderServer = &BaseProviderServer{}
var _ ResourceServer = &BaseProviderServer{}
var _ DataSourceServer = &BaseProviderServer{}

type BaseProviderServer struct {
	*BaseResourceServer
	*BaseDataSourceServer
}

func (b *BaseProviderServer) GetProviderSchema(_ context.Context, _ *GetProviderSchemaRequest) (*GetProviderSchemaResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseProviderServer) PrepareProviderConfig(_ context.Context, _ *PrepareProviderConfigRequest) (*PrepareProviderConfigResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseProviderServer) ConfigureProvider(_ context.Context, _ *ConfigureProviderRequest) (*ConfigureProviderResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseProviderServer) StopProvider(_ context.Context, _ *StopProviderRequest) (*StopProviderResponse, error) {
	return nil, tfproto.ErrUnimplemented
}
