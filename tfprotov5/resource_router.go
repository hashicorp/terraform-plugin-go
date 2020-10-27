package tfprotov5

import "context"

// ErrUnsupportedResource is returned when a `ResourceRouter` receives a
// request for a resource that is not found in its map. It will be set to the
// resource that was not found.
type ErrUnsupportedResource string

func (e ErrUnsupportedResource) Error() string {
	return "unsupported resource: " + string(e)
}

// ResourceRouter is an implementation of the `ResourceServer` interface.
// Requests will have their TypeName property matched to the keys in the
// ResourceRouter, and the equivalent method of the `ResourceServer` for that
// key will be called, passing in the request.
//
// This is used for giving each type of resource its own RPC methods, rather
// than needing to implement all resources in the same methods.
type ResourceRouter map[string]ResourceServer

// ValidateResourceTypeConfig is called when Terraform is attempting to
// validate the configuration for a resource. ValidateResourceTypeConfig uses
// the TypeName property of `req` to look up a `ResourceServer` in `r`, then
// calls that `ResourceServer`'s `ValidateResourceTypeConfig` method, passing
// in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) ValidateResourceTypeConfig(ctx context.Context, req *ValidateResourceTypeConfigRequest) (*ValidateResourceTypeConfigResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ValidateResourceTypeConfig(ctx, req)
}

// UpgradeResourceState is called when Terraform is attempting to upgrade a
// resource's state to the latest state format supported by the provider.
// UpgradeResourceState uses the TypeName property of `req` to look up a
// `ResourceServer` in `r`, then calls that `ResourceServer`'s
// `UpgradeResourceState` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) UpgradeResourceState(ctx context.Context, req *UpgradeResourceStateRequest) (*UpgradeResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.UpgradeResourceState(ctx, req)
}

// ReadResource is called when Terraform is attempting to retrieve the latest
// state for a resource. ReadResource uses the TypeName property of `req` to
// look up a `ResourceServer` in `r`, then calls that `ResourceServer`'s
// `ReadResource` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) ReadResource(ctx context.Context, req *ReadResourceRequest) (*ReadResourceResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ReadResource(ctx, req)
}

// PlanResourceChange is called when Terraform wants a provider's input on the
// planned changes to a resource. PlanResourceChange uses the TypeName property
// of `req` to look up a `ResourceServer` in `r`, then calls that
// `ResourceServer`'s `PlanResourceChange` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) PlanResourceChange(ctx context.Context, req *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.PlanResourceChange(ctx, req)
}

// ApplyResourceChange is called when Terraform wants a provider to make
// certain changes to a resource. ApplyResourceChange uses the TypeName
// property of `req` to look up a `ResourceServer` in `r`, then calls that
// `ResourceServer`'s `ApplyResourceChange` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) ApplyResourceChange(ctx context.Context, req *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ApplyResourceChange(ctx, req)
}

// ImportResourceState is called when Terraform wants a provider to import one
// or more resources so that Terraform can manage their state.
// ImportResourceState uses the TypeName property of `req` to look up a
// `ResourceServer` in `r`, then calls that `ResourceServer`'s
// `ImportResourceState` method, passing in `req`.
//
// If no key matches `req`'s TypeName property, an `ErrUnsupportedResource`
// error is returned, with its value set to `req`'s TypeName.
func (r ResourceRouter) ImportResourceState(ctx context.Context, req *ImportResourceStateRequest) (*ImportResourceStateResponse, error) {
	res, ok := r[req.TypeName]
	if !ok {
		return nil, ErrUnsupportedResource(req.TypeName)
	}
	return res.ImportResourceState(ctx, req)
}
