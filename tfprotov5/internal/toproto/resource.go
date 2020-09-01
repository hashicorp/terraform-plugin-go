package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateResourceTypeConfig_Request(in tfprotov5.ValidateResourceTypeConfigRequest) (tfplugin5.ValidateResourceTypeConfig_Request, error) {
	resp := tfplugin5.ValidateResourceTypeConfig_Request{
		TypeName: in.TypeName,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	return resp, nil
}

func ValidateResourceTypeConfig_Response(in tfprotov5.ValidateResourceTypeConfigResponse) (tfplugin5.ValidateResourceTypeConfig_Response, error) {
	var resp tfplugin5.ValidateResourceTypeConfig_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func UpgradeResourceState_Request(in tfprotov5.UpgradeResourceStateRequest) (tfplugin5.UpgradeResourceState_Request, error) {
	resp := tfplugin5.UpgradeResourceState_Request{
		TypeName: in.TypeName,
		Version:  in.Version,
	}
	if in.RawState != nil {
		state := RawState(*in.RawState)
		resp.RawState = &state
	}
	return resp, nil
}

func UpgradeResourceState_Response(in tfprotov5.UpgradeResourceStateResponse) (tfplugin5.UpgradeResourceState_Response, error) {
	var resp tfplugin5.UpgradeResourceState_Response
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ReadResource_Request(in tfprotov5.ReadResourceRequest) (tfplugin5.ReadResource_Request, error) {
	resp := tfplugin5.ReadResource_Request{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.CurrentState != nil {
		state, err := Cty(*in.CurrentState)
		if err != nil {
			return resp, fmt.Errorf("Error converting current state: %w", err)
		}
		resp.CurrentState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := Cty(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ReadResource_Response(in tfprotov5.ReadResourceResponse) (tfplugin5.ReadResource_Response, error) {
	resp := tfplugin5.ReadResource_Response{
		Private: in.Private,
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.NewState != nil {
		state, err := Cty(*in.NewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting new state: %w", err)
		}
		resp.NewState = &state
	}
	return resp, nil
}

func PlanResourceChange_Request(in tfprotov5.PlanResourceChangeRequest) (tfplugin5.PlanResourceChange_Request, error) {
	resp := tfplugin5.PlanResourceChange_Request{
		TypeName:     in.TypeName,
		PriorPrivate: in.PriorPrivate,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config :%w", err)
		}
		resp.Config = &config
	}
	if in.PriorState != nil {
		state, err := Cty(*in.PriorState)
		if err != nil {
			return resp, fmt.Errorf("Error converting prior state: %w", err)
		}
		resp.PriorState = &state
	}
	if in.ProposedNewState != nil {
		state, err := Cty(*in.ProposedNewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting proposed new state: %w", err)
		}
		resp.ProposedNewState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := Cty(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func PlanResourceChange_Response(in tfprotov5.PlanResourceChangeResponse) (tfplugin5.PlanResourceChange_Response, error) {
	resp := tfplugin5.PlanResourceChange_Response{
		PlannedPrivate: in.PlannedPrivate,
	}
	requiresReplace, err := AttributePaths(in.RequiresReplace)
	if err != nil {
		return resp, err
	}
	resp.RequiresReplace = requiresReplace
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.PlannedState != nil {
		state, err := Cty(*in.PlannedState)
		if err != nil {
			return resp, fmt.Errorf("Error converting planned state: %w", err)
		}
		resp.PlannedState = &state
	}
	return resp, nil
}

func ApplyResourceChange_Request(in tfprotov5.ApplyResourceChangeRequest) (tfplugin5.ApplyResourceChange_Request, error) {
	resp := tfplugin5.ApplyResourceChange_Request{
		TypeName:       in.TypeName,
		PlannedPrivate: in.PlannedPrivate,
	}
	if in.Config != nil {
		config, err := Cty(*in.Config)
		if err != nil {
			return resp, fmt.Errorf("Error converting config: %w", err)
		}
		resp.Config = &config
	}
	if in.PriorState != nil {
		state, err := Cty(*in.PriorState)
		if err != nil {
			return resp, fmt.Errorf("Error converting prior state: %w", err)
		}
		resp.PriorState = &state
	}
	if in.PlannedState != nil {
		state, err := Cty(*in.PlannedState)
		if err != nil {
			return resp, fmt.Errorf("Error converting planned state: %w", err)
		}
		resp.PlannedState = &state
	}
	if in.ProviderMeta != nil {
		meta, err := Cty(*in.ProviderMeta)
		if err != nil {
			return resp, fmt.Errorf("Error converting provider_meta: %w", err)
		}
		resp.ProviderMeta = &meta
	}
	return resp, nil
}

func ApplyResourceChange_Response(in tfprotov5.ApplyResourceChangeResponse) (tfplugin5.ApplyResourceChange_Response, error) {
	resp := tfplugin5.ApplyResourceChange_Response{
		Private: in.Private,
	}
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	if in.NewState != nil {
		state, err := Cty(*in.NewState)
		if err != nil {
			return resp, fmt.Errorf("Error converting new state: %w", err)
		}
		resp.NewState = &state
	}
	return resp, nil
}

func ImportResourceState_Request(in tfprotov5.ImportResourceStateRequest) (tfplugin5.ImportResourceState_Request, error) {
	return tfplugin5.ImportResourceState_Request{
		TypeName: in.TypeName,
		Id:       in.ID,
	}, nil
}

func ImportResourceState_Response(in tfprotov5.ImportResourceStateResponse) (tfplugin5.ImportResourceState_Response, error) {
	var resp tfplugin5.ImportResourceState_Response
	importedResources, err := ImportResourceState_ImportedResources(in.ImportedResources)
	if err != nil {
		return resp, err
	}
	resp.ImportedResources = importedResources
	diags, err := Diagnostics(in.Diagnostics)
	if err != nil {
		return resp, err
	}
	resp.Diagnostics = diags
	return resp, nil
}

func ImportResourceState_ImportedResource(in tfprotov5.ImportedResource) (tfplugin5.ImportResourceState_ImportedResource, error) {
	resp := tfplugin5.ImportResourceState_ImportedResource{
		TypeName: in.TypeName,
		Private:  in.Private,
	}
	if in.State != nil {
		state, err := Cty(*in.State)
		if err != nil {
			return resp, fmt.Errorf("Error converting state: %w", err)
		}
		resp.State = &state
	}
	return resp, nil
}

func ImportResourceState_ImportedResources(in []*tfprotov5.ImportedResource) ([]*tfplugin5.ImportResourceState_ImportedResource, error) {
	resp := make([]*tfplugin5.ImportResourceState_ImportedResource, 0, len(in))
	for _, i := range in {
		if i == nil {
			resp = append(resp, nil)
			continue
		}
		r, err := ImportResourceState_ImportedResource(*i)
		if err != nil {
			return resp, err
		}
		resp = append(resp, &r)
	}
	return resp, nil
}
