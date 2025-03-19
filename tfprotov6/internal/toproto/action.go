// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ActionSchema(in *tfprotov6.ActionSchema) *tfplugin6.ActionSchema {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.ActionSchema{
		LinkedResources: LinkedResources(in.LinkedResources),
		Block:           Schema_Block(in.Block),
		Version:         in.Version,
	}

	return resp
}

func CancelAction_Response(in *tfprotov6.CancelActionResponse) *tfplugin6.CancelAction_Response {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.CancelAction_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}

	return resp
}

func PlanAction_Response(in *tfprotov6.PlanActionResponse) *tfplugin6.PlanAction_Response {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.PlanAction_Response{
		Diagnostics:            Diagnostics(in.Diagnostics),
		PlannedResourceChanges: ActionResourceChanges(in.PlannedResourceChanges),
	}

	return resp
}

func InvokeAction_Event_Started_(in *tfprotov6.StartedActionEvent) *tfplugin6.InvokeAction_Event_Started_ {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.InvokeAction_Event_Started_{
		Started: &tfplugin6.InvokeAction_Event_Started{
			CancelationToken: in.CancellationToken,
		},
	}

	return resp
}

func InvokeAction_Event_Progress_(in *tfprotov6.ProgressActionEvent) *tfplugin6.InvokeAction_Event_Progress_ {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.InvokeAction_Event_Progress_{
		Progress: &tfplugin6.InvokeAction_Event_Progress{
			Stdout: in.StdOut,
			Stderr: in.StdErr,
		},
	}

	return resp
}

func InvokeAction_Event_Finished_(in *tfprotov6.FinishedActionEvent) *tfplugin6.InvokeAction_Event_Finished_ {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.InvokeAction_Event_Finished_{
		Finished: &tfplugin6.InvokeAction_Event_Finished{
			Outputs:         ActionResourceChanges(in.Outputs),
			ResourceChanges: ActionResourceChanges(in.ResourceChanges),
		},
	}

	return resp
}

func InvokeAction_Event_Diagnostics_(in *tfprotov6.DiagnosticsActionEvent) *tfplugin6.InvokeAction_Event_Diagnostics_ {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.InvokeAction_Event_Diagnostics_{
		Diagnostics: &tfplugin6.InvokeAction_Event_Diagnostics{
			Diagnostics: Diagnostics(in.Diagnostics),
		},
	}

	return resp
}

func InvokeAction_Event_Cancelled_(in *tfprotov6.CancelledActionEvent) *tfplugin6.InvokeAction_Event_Cancelled_ {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.InvokeAction_Event_Cancelled_{
		Cancelled: &tfplugin6.InvokeAction_Event_Cancelled{},
	}

	return resp
}
