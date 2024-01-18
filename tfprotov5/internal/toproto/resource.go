// Copyright (c) HashiCorp, Inc.

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func GetMetadata_ResourceMetadata(in *tfprotov5.ResourceMetadata) *tfplugin5.GetMetadata_ResourceMetadata {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.GetMetadata_ResourceMetadata{
		TypeName: in.TypeName,
	}

	return resp
}

func ValidateResourceTypeConfig_Response(in *tfprotov5.ValidateResourceTypeConfigResponse) (*tfplugin5.ValidateResourceTypeConfig_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.ValidateResourceTypeConfig_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func UpgradeResourceState_Response(in *tfprotov5.UpgradeResourceStateResponse) (*tfplugin5.UpgradeResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.UpgradeResourceState_Response{
		Diagnostics:   diags,
		UpgradedState: DynamicValue(in.UpgradedState),
	}

	return resp, nil
}

func ReadResource_Response(in *tfprotov5.ReadResourceResponse) (*tfplugin5.ReadResource_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.ReadResource_Response{
		Diagnostics: diags,
		NewState:    DynamicValue(in.NewState),
		Private:     in.Private,
	}

	return resp, nil
}

func PlanResourceChange_Response(in *tfprotov5.PlanResourceChangeResponse) (*tfplugin5.PlanResourceChange_Response, error) {
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

	resp := &tfplugin5.PlanResourceChange_Response{
		Diagnostics:      diags,
		LegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem, //nolint:staticcheck
		PlannedPrivate:   in.PlannedPrivate,
		PlannedState:     DynamicValue(in.PlannedState),
		RequiresReplace:  requiresReplace,
	}

	return resp, nil
}

func ApplyResourceChange_Response(in *tfprotov5.ApplyResourceChangeResponse) (*tfplugin5.ApplyResourceChange_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.ApplyResourceChange_Response{
		Diagnostics:      diags,
		LegacyTypeSystem: in.UnsafeToUseLegacyTypeSystem, //nolint:staticcheck
		NewState:         DynamicValue(in.NewState),
		Private:          in.Private,
	}

	return resp, nil
}

func ImportResourceState_Response(in *tfprotov5.ImportResourceStateResponse) (*tfplugin5.ImportResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.ImportResourceState_Response{
		Diagnostics:       diags,
		ImportedResources: ImportResourceState_ImportedResources(in.ImportedResources),
	}

	return resp, nil
}

func ImportResourceState_ImportedResource(in *tfprotov5.ImportedResource) *tfplugin5.ImportResourceState_ImportedResource {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.ImportResourceState_ImportedResource{
		Private:  in.Private,
		State:    DynamicValue(in.State),
		TypeName: in.TypeName,
	}

	return resp
}

func ImportResourceState_ImportedResources(in []*tfprotov5.ImportedResource) []*tfplugin5.ImportResourceState_ImportedResource {
	resp := make([]*tfplugin5.ImportResourceState_ImportedResource, 0, len(in))

	for _, i := range in {
		resp = append(resp, ImportResourceState_ImportedResource(i))
	}

	return resp
}

func MoveResourceState_Response(in *tfprotov5.MoveResourceStateResponse) (*tfplugin5.MoveResourceState_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin5.MoveResourceState_Response{
		Diagnostics: diags,
		TargetState: DynamicValue(in.TargetState),
	}

	return resp, nil
}
