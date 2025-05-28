// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import (
	"context"
)

// ActionServer is an interface containing the methods an action
// implementation needs to fill.
type ActionServer interface {
	// PlanAction is called when Terraform is attempting to
	// calculate a plan for an action.
	PlanAction(context.Context, *PlanActionRequest) (*PlanActionResponse, error)

	// InvokeAction is called when Terraform wants to execute the logic of an action
	InvokeAction(*InvokeActionRequest, *InvokeActionStreamingServer) error
}

// PlanActionRequest is the request Terraform sends when it is attempting to
// calculate a plan for an action.
type PlanActionRequest struct {
	// TypeName is the name of the action being called.
	TypeName string

	// Config is the configuration value of the action being called.
	Config *DynamicValue
}

// PlanActionResponse is the response from the provider with the new configuration
// expected after the action is invoked
type PlanActionResponse struct {
	Diagnostics []*Diagnostic

	// NewConfig is the new configuration value after the action has been applied.
	NewConfig *DynamicValue
}

// InvokeActionRequest is the request Terraform sends when it wants to execute
// the logic of an action
type InvokeActionRequest struct {
	// TypeName is the type name of the action being called.
	TypeName string

	// PlannedConfig is the configuration value of the action being called.
	PlannedConfig *DynamicValue
}

type InvokeActionStreamingServer interface {
	// Send sends an event to the client.
	Send(event *InvokeActionEvent) error

	// Context returns the context for the streaming server.
	Context() context.Context
}

type InvokeActionEvent interface {
	// unexported marker method to only allow these events to be passed to
	// the streaming server at compile time
	isInvokeActionEvent()
}

func (*InvokeActionEventStarted) isInvokeActionEvent()  {}
func (*InvokeActionEventProgress) isInvokeActionEvent() {}
func (*InvokeActionEventFinished) isInvokeActionEvent() {}

type InvokeActionEventStarted struct {
	CancellationToken string
	Diagnostics       []*Diagnostic
}
type InvokeActionEventProgress struct {
	Stdout      []string
	Stderr      []string
	Diagnostics []*Diagnostic
}
type InvokeActionEventFinished struct {
	NewConfig   *DynamicValue
	Cancelled   bool
	Diagnostics []*Diagnostic
}
