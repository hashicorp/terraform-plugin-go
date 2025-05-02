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
	Diagnostics []*Diagnostic
	NewConfig   *DynamicValue
}

type InvokeActionRequest struct {
	TypeName string
	Config   *DynamicValue
}

type InvokeActionResponse struct {
	CancellationToken string
	CallbackServer    InvokeActionCallBackServer
	Diagnostics       []*Diagnostic
}

type InvokeActionCallBackServer interface {
	Send(ctx context.Context, event InvokeActionEvent) error
}

// Invoke Action Events
type InvokeActionEvent interface {
	// TODO: make this interface unfillable to restrict implementations
	isInvokeActionEvent()
}

var _ InvokeActionEvent = &StartedActionEvent{}

type StartedActionEvent struct {
	CancellationToken string
	Diagnostics       []*Diagnostic
}

func (s *StartedActionEvent) isInvokeActionEvent() {}

type FinishedActionEvent struct {
	Outputs     map[string]*DynamicValue
	NewConfig   *DynamicValue
	Diagnostics []*Diagnostic
}

func (f *FinishedActionEvent) isInvokeActionEvent() {}

type CancelledActionEvent struct {
	Diagnostics []*Diagnostic
}

func (c *CancelledActionEvent) isInvokeActionEvent() {}

type ProgressActionEvent struct {
	StdOut      []string
	StdErr      []string
	Diagnostics []*Diagnostic
}

var _ InvokeActionEvent = &ProgressActionEvent{}

func (p *ProgressActionEvent) isInvokeActionEvent() {}
