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
	RequiresReplace []*tftypes.AttributePath
	PlannedPrivate  []byte
	Diagnostics     []*Diagnostic

	// This field should only be set by hashicorp/terraform-plugin-sdk.
	// It modifies Terraform's behavior to work with the legacy
	// expectations of that SDK.
	//
	// Nobody else should use this. Ever. For any reason. Just don't do it.
	//
	// We have to expose it here for terraform-plugin-sdk to be muxable, or
	// we wouldn't even be including it in this type. Don't use it. It may
	// go away or change behavior on you with no warning. It is
	// explicitly unsupported and not part of our SemVer guarantees.
	//
	// Deprecated: Really, just don't use this, you don't need it.
	UnsafeToUseLegacyTypeSystem bool
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

	// This field should only be set by hashicorp/terraform-plugin-sdk.
	// It modifies Terraform's behavior to work with the legacy
	// expectations of that SDK.
	//
	// Nobody else should use this. Ever. For any reason. Just don't do it.
	//
	// We have to expose it here for terraform-plugin-sdk to be muxable, or
	// we wouldn't even be including it in this type. Don't use it. It may
	// go away or change behavior on you with no warning. It is
	// explicitly unsupported and not part of our SemVer guarantees.
	//
	// Deprecated: Really, just don't use this, you don't need it.
	UnsafeToUseLegacyTypeSystem bool
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
