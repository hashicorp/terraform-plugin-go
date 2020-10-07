package tfprotov5

import "context"

type ErrUnsupportedResource string

func (e ErrUnsupportedResource) Error() string {
	return "unsupported resource: " + string(e)
}

type ResourceRouter map[string]ResourceServer

func (r ResourceRouter) ValidateResourceTypeConfig(ctx context.Context, req *ValidateResourceTypeConfigRequest) (*ValidateResourceTypeConfigResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ValidateResourceTypeConfig(ctx, req)
}

func (r ResourceRouter) UpgradeResourceState(ctx context.Context, req *UpgradeResourceStateRequest) (*UpgradeResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.UpgradeResourceState(ctx, req)
}

func (r ResourceRouter) ReadResource(ctx context.Context, req *ReadResourceRequest) (*ReadResourceResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ReadResource(ctx, req)
}

func (r ResourceRouter) PlanResourceChange(ctx context.Context, req *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.PlanResourceChange(ctx, req)
}

func (r ResourceRouter) ApplyResourceChange(ctx context.Context, req *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ApplyResourceChange(ctx, req)
}

func (r ResourceRouter) ImportResourceState(ctx context.Context, req *ImportResourceStateRequest) (*ImportResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ImportResourceState(ctx, req)
}
