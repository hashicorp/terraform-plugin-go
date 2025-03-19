// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

import (
	"context"
)

type ActionServer interface {
	PlanAction(context.Context, *PlanActionRequest) (*PlanActionResponse, error)

	InvokeAction(context.Context, *InvokeActionRequest, *InvokeActionResponse) error //Todo: change response signature?}

	CancelAction(context.Context, *CancelActionRequest) (*CancelActionResponse, error)
}

type Action struct {
}

type CancelActionRequest struct {
	CancellationToken string
	CancellationType  CancelType
}

type CancelActionResponse struct {
	Diagnostics []*Diagnostic
}

type PlanActionRequest struct {
	TypeName string
	Config   *DynamicValue
}

type PlanActionResponse struct {
	Diagnostics            []*Diagnostic
	PlannedResourceChanges map[string]*DynamicValue
}

type InvokeActionRequest struct {
	TypeName string
	Config   *DynamicValue
}

type InvokeActionResponse struct {
	CancellationToken string
	Events            chan InvokeActionEvent
	Diagnostics       []*Diagnostic
}

// Invoke Action Events
type InvokeActionEvent interface {
	// TODO: make this interface unfillable to restrict implementations
	isInvokeActionEvent()
}

var _ InvokeActionEvent = &StartedActionEvent{}

type StartedActionEvent struct {
	CancellationToken string
}

func (s *StartedActionEvent) isInvokeActionEvent() {}

type FinishedActionEvent struct {
	Outputs         map[string]*DynamicValue
	ResourceChanges map[string]*DynamicValue
}

func (f *FinishedActionEvent) isInvokeActionEvent() {}

type DiagnosticsActionEvent struct {
	Diagnostics []*Diagnostic
}

func (d *DiagnosticsActionEvent) isInvokeActionEvent() {}

type CancelledActionEvent struct{}

func (c *CancelledActionEvent) isInvokeActionEvent() {}

type ProgressActionEvent struct {
	StdOut []string
	StdErr []string
}

var _ InvokeActionEvent = &ProgressActionEvent{}

func (p *ProgressActionEvent) isInvokeActionEvent() {}
