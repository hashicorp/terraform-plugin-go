// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func CancelActionRequest(in *tfplugin6.CancelAction_Request) *tfprotov6.CancelActionRequest {
	if in == nil {
		return nil
	}

	resp := &tfprotov6.CancelActionRequest{
		CancellationToken: in.CancelationToken,
		CancellationType:  tfprotov6.CancelType(in.Type.Number()),
	}

	return resp

}

func InvokeActionRequest(in *tfplugin6.InvokeAction_Request) *tfprotov6.InvokeActionRequest {
	if in == nil {
		return nil
	}

	resp := &tfprotov6.InvokeActionRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}

	return resp

}

func PlanActionRequest(in *tfplugin6.PlanAction_Request) *tfprotov6.PlanActionRequest {
	if in == nil {
		return nil
	}

	resp := &tfprotov6.PlanActionRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}

	return resp

}
