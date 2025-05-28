// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func PlanActionRequest(in *tfplugin5.PlanAction_Request) *tfprotov5.PlanActionRequest {
	if in == nil {
		return nil
	}

	resp := &tfprotov5.PlanActionRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}

	return resp
}

func InvokeActionRequest(in *tfplugin5.InvokeAction_Request) *tfprotov5.InvokeActionRequest {
	if in == nil {
		return nil
	}

	resp := &tfprotov5.InvokeActionRequest{
		TypeName:      in.TypeName,
		PlannedConfig: DynamicValue(in.PlannedConfig),
	}

	return resp
}
