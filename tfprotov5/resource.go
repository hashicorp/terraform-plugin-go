package tfprotov5

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
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
	Config   *tftypes.RawValue
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
	CurrentState *RawState
	Private      []byte // TODO: should we handle this ourselves and not surface it?
	ProviderMeta *tftypes.RawValue
}

type ReadResourceResponse struct {
	NewState    *RawState
	Diagnostics []*Diagnostic
	Private     []byte // TODO: should we handle this ourselves and not surface it?
}

type PlanResourceChangeRequest struct {
	TypeName         string
	PriorState       *RawState
	ProposedNewState *RawState
	Config           *tftypes.RawValue
	PriorPrivate     []byte // TODO: should we handle this ourselves and not surface it?
	ProviderMeta     *tftypes.RawValue
}

type PlanResourceChangeResponse struct {
	PlannedState    *RawState
	RequiresReplace []*AttributePath
	PlannedPrivate  []byte // TODO: should we handle this ourselves and not surface it?
	Diagnostics     []*Diagnostic
}

type ApplyResourceChangeRequest struct {
	TypeName       string
	PriorState     *RawState
	PlannedState   *RawState
	Config         *tftypes.RawValue
	PlannedPrivate []byte // TODO: should we handle this ourselves and not surface it?
	ProviderMeta   *tftypes.RawValue
}

type ApplyResourceChangeResponse struct {
	NewState    *RawState
	Private     []byte // TODO: should we handle this ourselves and not surface it?
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
	State    *RawState
	Private  []byte // TODO: should we handle this ourselves and not surface it?
}
