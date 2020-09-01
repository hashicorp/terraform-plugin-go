package fromproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateResourceTypeConfigRequest(in tfplugin5.ValidateResourceTypeConfig_Request) (tfprotov5.ValidateResourceTypeConfigRequest, error) {
	resp := tfprotov5.ValidateResourceTypeConfigRequest{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := TerraformTypesRawValue(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func ValidateResourceTypeConfigResponse(in tfplugin5.ValidateResourceTypeConfig_Response) (tfprotov5.ValidateResourceTypeConfigResponse, error) {
	var resp tfprotov5.ValidateResourceTypeConfigResponse
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func UpgradeResourceStateRequest(in tfplugin5.UpgradeResourceState_Request) (tfprotov5.UpgradeResourceStateRequest, error) {
	resp := tfprotov5.UpgradeResourceStateRequest{
		TypeName: in.TypeName,
		Version:  in.Version,
	}
	if in.RawState != nil {
		state := RawState(*in.RawState)
		resp.RawState = &state
	}
	return resp, nil
}

func UpgradeResourceStateResponse(in tfplugin5.UpgradeResourceState_Response) (tfprotov5.UpgradeResourceStateResponse, error) {
	var resp tfprotov5.UpgradeResourceStateResponse
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ReadResourceRequest(in tfplugin5.ReadResource_Request) (tfprotov5.ReadResourceRequest, error) {
	resp := tfprotov5.ReadResourceRequest{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.CurrentState != nil {
		state, err := TerraformTypesRawValue(*in.CurrentState)
		if err != nil {
			return resp, fmt.Errorf("Error converting current state: %w", err)
		}
		resp.CurrentState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := TerraformTypesRawValue(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ReadResourceResponse(in tfplugin5.ReadResource_Response) (tfprotov5.ReadResourceResponse, error) {
	resp := tfprotov5.ReadResourceResponse{
		Private: in.Private,
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.NewState != nil {
		state, err := TerraformTypesRawValue(*in.NewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting new state: %w", err)
		}
		resp.NewState = &state
	}
	return resp, nil
}

func PlanResourceChangeRequest(in tfplugin5.PlanResourceChange_Request) (tfprotov5.PlanResourceChangeRequest, error) {
	resp := tfprotov5.PlanResourceChangeRequest{
		TypeName:     in.TypeName,
		PriorPrivate: in.PriorPrivate,
	}
	if in.Config != nil {
		config, err := TerraformTypesRawValue(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	if in.PriorState != nil {
		state, err := TerraformTypesRawValue(*in.PriorState)
		if err != nil {
			return resp, fmt.Errorf("Error converting prior state: %w", err)
		}
		resp.PriorState = &state
	}
	if in.ProposedNewState != nil {
		state, err := TerraformTypesRawValue(*in.ProposedNewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting proposed new state: %w", err)
		}
		resp.ProposedNewState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := TerraformTypesRawValue(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func PlanResourceChangeResponse(in tfplugin5.PlanResourceChange_Response) (tfprotov5.PlanResourceChangeResponse, error) {
	resp := tfprotov5.PlanResourceChangeResponse{
		PlannedPrivate: in.PlannedPrivate,
	}
	attributePaths, err := AttributePaths(in.RequiresReplace)
	if err != nil {
		return resp, err
	}
	resp.RequiresReplace = attributePaths
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.PlannedState != nil {
		state, err := TerraformTypesRawValue(*in.PlannedState)
		if err != nil {
			return resp, fmt.Errorf("Error converting planned state: %w", err)
		}
		resp.PlannedState = &state
	}
	return resp, nil
}

func ApplyResourceChangeRequest(in tfplugin5.ApplyResourceChange_Request) (tfprotov5.ApplyResourceChangeRequest, error) {
	resp := tfprotov5.ApplyResourceChangeRequest{
		TypeName:       in.TypeName,
		PlannedPrivate: in.PlannedPrivate,
	}
	if in.Config != nil {
		config, err := TerraformTypesRawValue(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	if in.PriorState != nil {
		state, err := TerraformTypesRawValue(*in.PriorState)
		if err != nil {
			return resp, fmt.Errorf("Error converting prior state: %w", err)
		}
		resp.PriorState = &state
	}
	if in.PlannedState != nil {
		state, err := TerraformTypesRawValue(*in.PlannedState)
		if err != nil {
			return resp, fmt.Errorf("Error converting planned state: %w", err)
		}
		resp.PlannedState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := TerraformTypesRawValue(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ApplyResourceChangeResponse(in tfplugin5.ApplyResourceChange_Response) (tfprotov5.ApplyResourceChangeResponse, error) {
	resp := tfprotov5.ApplyResourceChangeResponse{
		Private: in.Private,
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.NewState != nil {
		state, err := TerraformTypesRawValue(*in.NewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting new state: %w", err)
		}
		resp.NewState = &state
	}
	return resp, nil
}

func ImportResourceStateRequest(in tfplugin5.ImportResourceState_Request) (tfprotov5.ImportResourceStateRequest, error) {
	return tfprotov5.ImportResourceStateRequest{
		TypeName: in.TypeName,
		ID:       in.Id,
	}, nil
}

func ImportResourceStateResponse(in tfplugin5.ImportResourceState_Response) (tfprotov5.ImportResourceStateResponse, error) {
	var resp tfprotov5.ImportResourceStateResponse
	imported, err := ImportedResources(in.ImportedResources)
	if err != nil {
		return resp, err
	}
	resp.ImportedResources = imported
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ImportedResource(in tfplugin5.ImportResourceState_ImportedResource) (tfprotov5.ImportedResource, error) {
	resp := tfprotov5.ImportedResource{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.State != nil {
		state, err := TerraformTypesRawValue(*in.State)
		if err != nil {
			return resp, fmt.Errorf("Error converting state: %w", err)
		}
		resp.State = &state
	}
	return resp, nil
}

func ImportedResources(in []*tfplugin5.ImportResourceState_ImportedResource) ([]*tfprotov5.ImportedResource, error) {
	resp := make([]*tfprotov5.ImportedResource, 0, len(in))
	for pos, i := range in {
		if i == nil {
			resp = append(resp, nil)
			continue
		}
		r, err := ImportedResource(*i)
		if err != nil {
			return resp, fmt.Errorf("Error converting imported resource %d/%d: %w", pos+1, len(in), err)
		}
		resp = append(resp, &r)
	}
	return resp, nil
}
