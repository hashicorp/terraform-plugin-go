package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func GetMetadata_ResourceMetadata(in *tfprotov6.ResourceMetadata) *tfplugin6.GetMetadata_ResourceMetadata {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.GetMetadata_ResourceMetadata{
		TypeName: in.TypeName,
	}

	return resp
}

func ValidateResourceConfig_Response(in *tfprotov6.ValidateResourceConfigResponse) (*tfplugin6.ValidateResourceConfig_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ValidateResourceConfig_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func UpgradeResourceState_Response(in *tfprotov6.UpgradeResourceStateResponse) (*tfplugin6.UpgradeResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.UpgradeResourceState_Response{
		Diagnostics:   diags,
		UpgradedState: DynamicValue(in.UpgradedState),
	}

	return resp, nil
}

func ReadResource_Response(in *tfprotov6.ReadResourceResponse) (*tfplugin6.ReadResource_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ReadResource_Response{
		Diagnostics: diags,
		NewState:    DynamicValue(in.NewState),
		Private:     in.Private,
	}

	return resp, nil
}

func PlanResourceChange_Response(in *tfprotov6.PlanResourceChangeResponse) (*tfplugin6.PlanResourceChange_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	requiresReplace, err := AttributePaths(in.RequiresReplace)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.PlanResourceChange_Response{
		Diagnostics:      diags,
		LegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem, //nolint:staticcheck
		PlannedPrivate:   in.PlannedPrivate,
		PlannedState:     DynamicValue(in.PlannedState),
		RequiresReplace:  requiresReplace,
	}

	return resp, nil
}

func ApplyResourceChange_Response(in *tfprotov6.ApplyResourceChangeResponse) (*tfplugin6.ApplyResourceChange_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ApplyResourceChange_Response{
		Diagnostics:      diags,
		LegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem, //nolint:staticcheck
		NewState:         DynamicValue(in.NewState),
		Private:          in.Private,
	}

	return resp, nil
}

func ImportResourceState_Response(in *tfprotov6.ImportResourceStateResponse) (*tfplugin6.ImportResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ImportResourceState_Response{
		Diagnostics:       diags,
		ImportedResources: ImportResourceState_ImportedResources(in.ImportedResources),
	}

	return resp, nil
}

func ImportResourceState_ImportedResource(in *tfprotov6.ImportedResource) *tfplugin6.ImportResourceState_ImportedResource {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.ImportResourceState_ImportedResource{
		Private:  in.Private,
		State:    DynamicValue(in.State),
		TypeName: in.TypeName,
	}

	return resp
}

func ImportResourceState_ImportedResources(in []*tfprotov6.ImportedResource) []*tfplugin6.ImportResourceState_ImportedResource {
	resp := make([]*tfplugin6.ImportResourceState_ImportedResource, 0, len(in))

	for _, i := range in {
		resp = append(resp, ImportResourceState_ImportedResource(i))
	}

	return resp
}

func MoveResourceState_Response(in *tfprotov6.MoveResourceStateResponse) (*tfplugin6.MoveResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.MoveResourceState_Response{
		Diagnostics: diags,
		TargetState: DynamicValue(in.TargetState),
	}

	return resp, nil
}
