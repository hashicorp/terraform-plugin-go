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
	CurrentState *tftypes.RawValue
	Private      []byte
	ProviderMeta *tftypes.RawValue
}

type ReadResourceResponse struct {
	NewState    *tftypes.RawValue
	Diagnostics []*Diagnostic
	Private     []byte
}

type PlanResourceChangeRequest struct {
	TypeName         string
	PriorState       *tftypes.RawValue
	ProposedNewState *tftypes.RawValue
	Config           *tftypes.RawValue
	PriorPrivate     []byte
	ProviderMeta     *tftypes.RawValue
}

type PlanResourceChangeResponse struct {
	PlannedState    *tftypes.RawValue
	RequiresReplace []*AttributePath
	PlannedPrivate  []byte
	Diagnostics     []*Diagnostic
}

type ApplyResourceChangeRequest struct {
	TypeName       string
	PriorState     *tftypes.RawValue
	PlannedState   *tftypes.RawValue
	Config         *tftypes.RawValue
	PlannedPrivate []byte
	ProviderMeta   *tftypes.RawValue
}

type ApplyResourceChangeResponse struct {
	NewState    *tftypes.RawValue
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
	State    *tftypes.RawValue
	Private  []byte
}
