package tfprotov5

import (
	"context"

	tfproto "github.com/hashicorp/terraform-protocol-go"
)

var _ ResourceServer = &BaseResourceServer{}

type BaseResourceServer struct{}

func (b *BaseResourceServer) ValidateResourceTypeConfig(_ context.Context, _ *ValidateResourceTypeConfigRequest) (*ValidateResourceTypeConfigResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseResourceServer) UpgradeResourceState(_ context.Context, _ *UpgradeResourceStateRequest) (*UpgradeResourceStateResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseResourceServer) ReadResource(_ context.Context, _ *ReadResourceRequest) (*ReadResourceResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseResourceServer) PlanResourceChange(_ context.Context, _ *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseResourceServer) ApplyResourceChange(_ context.Context, _ *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error) {
	return nil, tfproto.ErrUnimplemented
}

func (b *BaseResourceServer) ImportResourceState(_ context.Context, _ *ImportResourceStateRequest) (*ImportResourceStateResponse, error) {
	return nil, tfproto.ErrUnimplemented
}
