// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func PlanAction_Response(in *tfprotov5.PlanActionResponse) *tfplugin5.PlanAction_Response {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.PlanAction_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
		NewConfig:   DynamicValue(in.NewConfig),
	}

	return resp
}

func ActionSchema(in *tfprotov5.ActionSchema) *tfplugin5.ActionSchema {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.ActionSchema{
		Version:         in.Version,
		Block:           Schema_Block(in.Block),
		LinkedResources: make([]*tfplugin5.ActionSchema_LinkedResource, 0, len(in.LinkedResources)),
	}

	for _, linkedResource := range in.LinkedResources {
		resp.LinkedResources = append(resp.LinkedResources, ActionSchema_LinkedResource(linkedResource))
	}

	return resp
}

func ActionSchema_LinkedResource(in *tfprotov5.ActionSchemaLinkedResource) *tfplugin5.ActionSchema_LinkedResource {
	if in == nil {
		return nil
	}

	resp := &tfplugin5.ActionSchema_LinkedResource{
		Attribute:        AttributePath(in.Attribute),
		Type:             in.Type,
		LinkedAttributes: make([]*tfplugin5.ActionSchema_LinkedResource_LinkedAttribute, 0, len(in.LinkedAttributes)),
	}

	for _, linkedAttribute := range in.LinkedAttributes {
		resp.LinkedAttributes = append(resp.LinkedAttributes, ActionSchema_LinkedResource_LinkedAttribute(linkedAttribute))
	}

	return resp
}

func ActionSchema_LinkedResource_LinkedAttribute(in *tfprotov5.ActionSchemaLinkedResourceLinkedAttribute) *tfplugin5.ActionSchema_LinkedResource_LinkedAttribute {
	if in == nil {
		return nil
	}

	return &tfplugin5.ActionSchema_LinkedResource_LinkedAttribute{
		Attribute: AttributePath(in.Attribute),
	}
}

func InvokeActionEvent(in *tfprotov5.InvokeActionEvent) *tfplugin5.InvokeAction_Event {
	if in == nil {
		return nil
	}
	switch e := (*in).(type) {
	case *tfprotov5.InvokeActionEventStarted:
		return &tfplugin5.InvokeAction_Event{
			Event: &tfplugin5.InvokeAction_Event_Started_{
				Started: &tfplugin5.InvokeAction_Event_Started{
					CancellationToken: e.CancellationToken,
					Diagnostics:       Diagnostics(e.Diagnostics),
				},
			},
		}
	case *tfprotov5.InvokeActionEventProgress:
		return &tfplugin5.InvokeAction_Event{
			Event: &tfplugin5.InvokeAction_Event_Progress_{
				Progress: &tfplugin5.InvokeAction_Event_Progress{
					Stdout:      e.Stdout,
					Stderr:      e.Stderr,
					Diagnostics: Diagnostics(e.Diagnostics),
				},
			},
		}
	case *tfprotov5.InvokeActionEventFinished:
		return &tfplugin5.InvokeAction_Event{
			Event: &tfplugin5.InvokeAction_Event_Finished_{
				Finished: &tfplugin5.InvokeAction_Event_Finished{
					NewConfig:   DynamicValue(e.NewConfig),
					Cancelled:   e.Cancelled,
					Diagnostics: Diagnostics(e.Diagnostics),
				},
			},
		}
	}
	return nil // TODO: error instead?
}
