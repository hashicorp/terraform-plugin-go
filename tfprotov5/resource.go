package tfprotov5

import (
	"context"
)

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
	Config   *DynamicValue
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
	UpgradedState *RawState
	Diagnostics   []*Diagnostic
}

type ReadResourceRequest struct {
	TypeName     string
	CurrentState *DynamicValue
	Private      []byte
	ProviderMeta *DynamicValue
}

type ReadResourceResponse struct {
	NewState    *DynamicValue
	Diagnostics []*Diagnostic
	Private     []byte
}

type PlanResourceChangeRequest struct {
	TypeName         string
	PriorState       *DynamicValue
	ProposedNewState *DynamicValue
	Config           *DynamicValue
	PriorPrivate     []byte
	ProviderMeta     *DynamicValue
}

type PlanResourceChangeResponse struct {
	PlannedState    *DynamicValue
	RequiresReplace []*AttributePath
	PlannedPrivate  []byte
	Diagnostics     []*Diagnostic
}

type ApplyResourceChangeRequest struct {
	TypeName       string
	PriorState     *DynamicValue
	PlannedState   *DynamicValue
	Config         *DynamicValue
	PlannedPrivate []byte
	ProviderMeta   *DynamicValue
}

type ApplyResourceChangeResponse struct {
	NewState    *DynamicValue
	Private     []byte
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
	State    *DynamicValue
	Private  []byte
}
