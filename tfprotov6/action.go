// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

// Function describes the definition of a function. Result must be defined.
type Action struct {
	// Parameters is the ordered list of positional function parameters.
	Parameters []*FunctionParameter

	// VariadicParameter is an optional final parameter which accepts zero or
	// more argument values, in which Terraform will send an ordered list of the
	// parameter type.
	VariadicParameter *FunctionParameter

	// Return is the function result.
	Return *FunctionReturn

	// Summary is the shortened human-readable documentation for the function.
	Summary string

	// Description is the longer human-readable documentation for the function.
	Description string

	// DescriptionKind indicates the formatting and encoding that the
	// Description field is using.
	DescriptionKind StringKind

	// DeprecationMessage is the human-readable documentation if the function
	// is deprecated. This message should be practitioner oriented to explain
	// how their configuration should be updated.
	DeprecationMessage string
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
	Events            <-chan InvokeActionEvent
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
