package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func ValidateResourceTypeConfigRequest(in tfplugin5.ValidateResourceTypeConfig_Request) tfprotov5.ValidateResourceTypeConfigRequest {
	return tfprotov5.ValidateResourceTypeConfigRequest{
		TypeName: in.TypeName,
		Config:   nil, // TODO: figure out how to unmarshal config
	}
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
		UpgradedState: nil, // TODO: figure out how to unmarshal state
		Diagnostics:   Diagnostics(in.Diagnostics),
	}
}

func ReadResourceRequest(in tfplugin5.ReadResource_Request) tfprotov5.ReadResourceRequest {
	return tfprotov5.ReadResourceRequest{
		TypeName:     in.TypeName,
		CurrentState: nil, // TODO: figure out how to unmarshal state
		Private:      in.Private,
		ProviderMeta: nil, // TODO: figure out how to unmarshal cty
	}
}

func ReadResourceResponse(in tfplugin5.ReadResource_Response) tfprotov5.ReadResourceResponse {
	return tfprotov5.ReadResourceResponse{
		NewState:    nil, // TODO: figure out how to marshal state
		Diagnostics: Diagnostics(in.Diagnostics),
		Private:     in.Private,
	}
}

func PlanResourceChangeRequest(in tfplugin5.PlanResourceChange_Request) tfprotov5.PlanResourceChangeRequest {
	return tfprotov5.PlanResourceChangeRequest{
		TypeName:         in.TypeName,
		PriorState:       nil, // TODO: figure out how to unmarshal state
		ProposedNewState: nil, // TODO: figure out how to unmarshal state
		Config:           nil, // TODO: figure out how unmarshal config
		PriorPrivate:     in.PriorPrivate,
		ProviderMeta:     nil, // TODO: figure out how to unmarshal cty
	}
}

func PlanResourceChangeResponse(in tfplugin5.PlanResourceChange_Response) tfprotov5.PlanResourceChangeResponse {
	return tfprotov5.PlanResourceChangeResponse{
		PlannedState:    nil, // TODO: figure out how to unmarshal state
		RequiresReplace: AttributePaths(in.RequiresReplace),
		PlannedPrivate:  in.PlannedPrivate,
		Diagnostics:     Diagnostics(in.Diagnostics),
	}
}

func ApplyResourceChangeRequest(in tfplugin5.ApplyResourceChange_Request) tfprotov5.ApplyResourceChangeRequest {
	return tfprotov5.ApplyResourceChangeRequest{
		TypeName:       in.TypeName,
		PriorState:     nil, // TODO: figure out how to unmarshal state
		PlannedState:   nil, // TODO: figure out how to unmarshal state
		Config:         nil, // TODO: figure out how to unmarshal cty
		PlannedPrivate: in.PlannedPrivate,
		ProviderMeta:   nil, // TODO: figure out how to unmarshal cty
	}
}

func ApplyResourceChangeResponse(in tfplugin5.ApplyResourceChange_Response) tfprotov5.ApplyResourceChangeResponse {
	return tfprotov5.ApplyResourceChangeResponse{
		NewState:    nil, // TODO: figure out how to unmarshal state
		Private:     in.Private,
		Diagnostics: Diagnostics(in.Diagnostics),
	}
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
	return tfprotov5.ImportedResource{
		TypeName: in.TypeName,
		State:    nil, // TODO: figure out how to unmarshal state
		Private:  in.Private,
	}
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
