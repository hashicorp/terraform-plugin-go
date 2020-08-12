package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateResourceTypeConfigRequest(in tfplugin5.ValidateResourceTypeConfig_Request) tfprotov5.ValidateResourceTypeConfigRequest {
	resp := tfprotov5.ValidateResourceTypeConfigRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	return resp
}

func ValidateResourceTypeConfigResponse(in tfplugin5.ValidateResourceTypeConfig_Response) tfprotov5.ValidateResourceTypeConfigResponse {
	return tfprotov5.ValidateResourceTypeConfigResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func UpgradeResourceStateRequest(in tfplugin5.UpgradeResourceState_Request) tfprotov5.UpgradeResourceStateRequest {
	resp := tfprotov5.UpgradeResourceStateRequest{
		TypeName: in.TypeName,
		Version:  in.Version,
	}
	if in.RawState != nil {
		state := RawState(*in.RawState)
		resp.RawState = &state
	}
	return resp
}

func UpgradeResourceStateResponse(in tfplugin5.UpgradeResourceState_Response) tfprotov5.UpgradeResourceStateResponse {
	return tfprotov5.UpgradeResourceStateResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ReadResourceRequest(in tfplugin5.ReadResource_Request) tfprotov5.ReadResourceRequest {
	resp := tfprotov5.ReadResourceRequest{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.CurrentState != nil {
		state := TerraformTypesRawValue(*in.CurrentState)
		resp.CurrentState = &state
	}
	if in.ProviderMeta != nil {
		meta := TerraformTypesRawValue(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ReadResourceResponse(in tfplugin5.ReadResource_Response) tfprotov5.ReadResourceResponse {
	resp := tfprotov5.ReadResourceResponse{
		Diagnostics: Diagnostics(in.Diagnostics),
		Private:     in.Private,
	}
	if in.NewState != nil {
		state := TerraformTypesRawValue(*in.NewState)
		resp.NewState = &state
	}
	return resp
}

func PlanResourceChangeRequest(in tfplugin5.PlanResourceChange_Request) tfprotov5.PlanResourceChangeRequest {
	resp := tfprotov5.PlanResourceChangeRequest{
		TypeName:     in.TypeName,
		PriorPrivate: in.PriorPrivate,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	if in.PriorState != nil {
		state := TerraformTypesRawValue(*in.PriorState)
		resp.PriorState = &state
	}
	if in.ProposedNewState != nil {
		state := TerraformTypesRawValue(*in.ProposedNewState)
		resp.ProposedNewState = &state
	}
	if in.ProviderMeta != nil {
		meta := TerraformTypesRawValue(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func PlanResourceChangeResponse(in tfplugin5.PlanResourceChange_Response) tfprotov5.PlanResourceChangeResponse {
	resp := tfprotov5.PlanResourceChangeResponse{
		RequiresReplace: AttributePaths(in.RequiresReplace),
		PlannedPrivate:  in.PlannedPrivate,
		Diagnostics:     Diagnostics(in.Diagnostics),
	}
	if in.PlannedState != nil {
		state := TerraformTypesRawValue(*in.PlannedState)
		resp.PlannedState = &state
	}
	return resp
}

func ApplyResourceChangeRequest(in tfplugin5.ApplyResourceChange_Request) tfprotov5.ApplyResourceChangeRequest {
	resp := tfprotov5.ApplyResourceChangeRequest{
		TypeName:       in.TypeName,
		PlannedPrivate: in.PlannedPrivate,
	}
	if in.Config != nil {
		config := TerraformTypesRawValue(*in.Config)
		resp.Config = &config
	}
	if in.PriorState != nil {
		state := TerraformTypesRawValue(*in.PriorState)
		resp.PriorState = &state
	}
	if in.PlannedState != nil {
		state := TerraformTypesRawValue(*in.PlannedState)
		resp.PlannedState = &state
	}
	if in.ProviderMeta != nil {
		meta := TerraformTypesRawValue(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ApplyResourceChangeResponse(in tfplugin5.ApplyResourceChange_Response) tfprotov5.ApplyResourceChangeResponse {
	resp := tfprotov5.ApplyResourceChangeResponse{
		Private:     in.Private,
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.NewState != nil {
		state := TerraformTypesRawValue(*in.NewState)
		resp.NewState = &state
	}
	return resp
}

func ImportResourceStateRequest(in tfplugin5.ImportResourceState_Request) tfprotov5.ImportResourceStateRequest {
	return tfprotov5.ImportResourceStateRequest{
		TypeName: in.TypeName,
		ID:       in.Id,
	}
}

func ImportResourceStateResponse(in tfplugin5.ImportResourceState_Response) tfprotov5.ImportResourceStateResponse {
	return tfprotov5.ImportResourceStateResponse{
		ImportedResources: ImportedResources(in.ImportedResources),
		Diagnostics:       Diagnostics(in.Diagnostics),
	}
}

func ImportedResource(in tfplugin5.ImportResourceState_ImportedResource) tfprotov5.ImportedResource {
	resp := tfprotov5.ImportedResource{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.State != nil {
		state := TerraformTypesRawValue(*in.State)
		resp.State = &state
	}
	return resp
}

func ImportedResources(in []*tfplugin5.ImportResourceState_ImportedResource) []*tfprotov5.ImportedResource {
	resp := make([]*tfprotov5.ImportedResource, 0, len(in))
	for _, i := range in {
		if i == nil {
			resp = append(resp, nil)
			continue
		}
		r := ImportedResource(*i)
		resp = append(resp, &r)
	}
	return resp
}
