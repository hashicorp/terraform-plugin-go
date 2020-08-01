package tfprotov5

import "context"

type ResourceServer interface {
	ValidateResourceTypeConfig(context.Context, *ValidateResourceTypeConfigRequest) (*ValidateResourceTypeConfigResponse, error)
	UpgradeResourceState(context.Context, *UpgradeResourceStateRequest) (*UpgradeResourceStateResponse, error)
	ReadResource(context.Context, *ReadResourceRequest) (*ReadResourceResponse, error)
	PlanResourceChange(context.Context, *PlanResourceChangeRequest) (*PlanResourceChangeResponse, error)
	ApplyResourceChange(context.Context, *ApplyResourceChangeRequest) (*ApplyResourceChangeResponse, error)
	ImportResourceState(context.Context, *ImportResourceStateRequest) (*ImportResourceStateResponse, error)
}

type ValidateResourceTypeConfigRequest struct {
	TypeName string
	Config   *CtyBlock
}

type ValidateResourceTypeConfigResponse struct {
	Diagnostics []*Diagnostic
}

type UpgradeResourceStateRequest struct {
	TypeName string
	Version  int64
	RawState *RawState
}

type UpgradeResourceStateResponse struct {
	UpgradedState *CtyBlock // TODO: figure out how to represent state
	Diagnostics   []*Diagnostic
}

type ReadResourceRequest struct {
	TypeName     string
	CurrentState *CtyBlock // TODO: figure out how to represent state
	Private      []byte    // TODO: should we handle this ourselves and not surface it?
	ProviderMeta *CtyBlock
}

type ReadResourceResponse struct {
	NewState    *CtyBlock // TODO: figure out how to represent state
	Diagnostics []*Diagnostic
	Private     []byte // TODO: should we handle this ourselves and not surface it?
}

type PlanResourceChangeRequest struct {
	TypeName         string
	PriorState       *CtyBlock // TODO: figure out how to represent state
	ProposedNewState *CtyBlock // TODO: figure out how to represent state
	Config           *CtyBlock
	PriorPrivate     []byte // TODO: should we handle this ourselves and not surface it?
	ProviderMeta     *CtyBlock
}

type PlanResourceChangeResponse struct {
	PlannedState    *CtyBlock // TODO: figure out how to represent state
	RequiresReplace []*AttributePath
	PlannedPrivate  []byte // TODO: should we handle this ourselves and not surface it?
	Diagnostics     []*Diagnostic
}

type ApplyResourceChangeRequest struct {
	TypeName       string
	PriorState     *CtyBlock // TODO: figure out how to represent state
	PlannedState   *CtyBlock // TODO: figure out how to represent state
	Config         *CtyBlock
	PlannedPrivate []byte // TODO: should we handle this ourselves and not surface it?
	ProviderMeta   *CtyBlock
}

type ApplyResourceChangeResponse struct {
	NewState    *CtyBlock // TODO: figure out how to represent state
	Private     []byte    // TODO: should we handle this ourselves and not surface it?
	Diagnostics []*Diagnostic
}

type ImportResourceStateRequest struct {
	TypeName string
	ID       string
}

type ImportResourceStateResponse struct {
	ImportedResources []*ImportedResource
	Diagnostics       []*Diagnostic
}

type ImportedResource struct {
	TypeName string
	State    *CtyBlock // TODO: figure out how to represent state
	Private  []byte    // TODO: should we handle this ourselves and not surface it?
}
