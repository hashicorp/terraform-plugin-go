package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateResourceTypeConfig_Request(in tfprotov5.ValidateResourceTypeConfigRequest) tfplugin5.ValidateResourceTypeConfig_Request {
	resp := tfplugin5.ValidateResourceTypeConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	return resp
}

func ValidateResourceTypeConfig_Response(in tfprotov5.ValidateResourceTypeConfigResponse) tfplugin5.ValidateResourceTypeConfig_Response {
	return tfplugin5.ValidateResourceTypeConfig_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func UpgradeResourceState_Request(in tfprotov5.UpgradeResourceStateRequest) tfplugin5.UpgradeResourceState_Request {
	resp := tfplugin5.UpgradeResourceState_Request{
		TypeName: in.TypeName,
		Version:  in.Version,
	}
	if in.RawState != nil {
		state := RawState(*in.RawState)
		resp.RawState = &state
	}
	return resp
}

func UpgradeResourceState_Response(in tfprotov5.UpgradeResourceStateResponse) tfplugin5.UpgradeResourceState_Response {
	return tfplugin5.UpgradeResourceState_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ReadResource_Request(in tfprotov5.ReadResourceRequest) tfplugin5.ReadResource_Request {
	resp := tfplugin5.ReadResource_Request{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.CurrentState != nil {
		state := Cty(*in.CurrentState)
		resp.CurrentState = &state
	}
	if in.ProviderMeta != nil {
		meta := Cty(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ReadResource_Response(in tfprotov5.ReadResourceResponse) tfplugin5.ReadResource_Response {
	resp := tfplugin5.ReadResource_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
		Private:     in.Private,
	}
	if in.NewState != nil {
		state := Cty(*in.NewState)
		resp.NewState = &state
	}
	return resp
}

func PlanResourceChange_Request(in tfprotov5.PlanResourceChangeRequest) tfplugin5.PlanResourceChange_Request {
	resp := tfplugin5.PlanResourceChange_Request{
		TypeName:     in.TypeName,
		PriorPrivate: in.PriorPrivate,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	if in.PriorState != nil {
		state := Cty(*in.PriorState)
		resp.PriorState = &state
	}
	if in.ProposedNewState != nil {
		state := Cty(*in.ProposedNewState)
		resp.ProposedNewState = &state
	}
	if in.ProviderMeta != nil {
		meta := Cty(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func PlanResourceChange_Response(in tfprotov5.PlanResourceChangeResponse) tfplugin5.PlanResourceChange_Response {
	resp := tfplugin5.PlanResourceChange_Response{
		RequiresReplace: AttributePaths(in.RequiresReplace),
		PlannedPrivate:  in.PlannedPrivate,
		Diagnostics:     Diagnostics(in.Diagnostics),
	}
	if in.PlannedState != nil {
		state := Cty(*in.PlannedState)
		resp.PlannedState = &state
	}
	return resp
}

func ApplyResourceChange_Request(in tfprotov5.ApplyResourceChangeRequest) tfplugin5.ApplyResourceChange_Request {
	resp := tfplugin5.ApplyResourceChange_Request{
		TypeName:       in.TypeName,
		PlannedPrivate: in.PlannedPrivate,
	}
	if in.Config != nil {
		config := Cty(*in.Config)
		resp.Config = &config
	}
	if in.PriorState != nil {
		state := Cty(*in.PriorState)
		resp.PriorState = &state
	}
	if in.PlannedState != nil {
		state := Cty(*in.PlannedState)
		resp.PlannedState = &state
	}
	if in.ProviderMeta != nil {
		meta := Cty(*in.ProviderMeta)
		resp.ProviderMeta = &meta
	}
	return resp
}

func ApplyResourceChange_Response(in tfprotov5.ApplyResourceChangeResponse) tfplugin5.ApplyResourceChange_Response {
	resp := tfplugin5.ApplyResourceChange_Response{
		Private:     in.Private,
		Diagnostics: Diagnostics(in.Diagnostics),
	}
	if in.NewState != nil {
		state := Cty(*in.NewState)
		resp.NewState = &state
	}
	return resp
}

func ImportResourceState_Request(in tfprotov5.ImportResourceStateRequest) tfplugin5.ImportResourceState_Request {
	return tfplugin5.ImportResourceState_Request{
		TypeName: in.TypeName,
		Id:       in.ID,
	}
}

func ImportResourceState_Response(in tfprotov5.ImportResourceStateResponse) tfplugin5.ImportResourceState_Response {
	return tfplugin5.ImportResourceState_Response{
		ImportedResources: ImportResourceState_ImportedResources(in.ImportedResources),
		Diagnostics:       Diagnostics(in.Diagnostics),
	}
}

func ImportResourceState_ImportedResource(in tfprotov5.ImportedResource) tfplugin5.ImportResourceState_ImportedResource {
	resp := tfplugin5.ImportResourceState_ImportedResource{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.State != nil {
		state := Cty(*in.State)
		resp.State = &state
	}
	return resp
}

func ImportResourceState_ImportedResources(in []*tfprotov5.ImportedResource) []*tfplugin5.ImportResourceState_ImportedResource {
	resp := make([]*tfplugin5.ImportResourceState_ImportedResource, 0, len(in))
	for _, i := range in {
		if i == nil {
			resp = append(resp, nil)
			continue
		}
		r := ImportResourceState_ImportedResource(*i)
		resp = append(resp, &r)
	}
	return resp
}
