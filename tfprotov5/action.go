// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import (
	"context"
	"iter"
)

// ActionMetadata describes metadata for an action in the GetMetadata
// RPC.
type ActionMetadata struct {
	// TypeName is the name of the action.
	TypeName string
}

// ActionServer is an interface containing the methods an action
// implementation needs to fill.
type ActionServer interface {
	// TODO: pkg docs
	PlanAction(context.Context, *PlanActionRequest) (*PlanActionResponse, error)
	// TODO: pkg docs
	InvokeAction(context.Context, *InvokeActionRequest) (*InvokeActionServerStream, error)
}

// TODO: pkg docs
type PlanActionRequest struct {
	// ActionType is the name of the action being called.
	ActionType string

	LinkedResources []*ProposedLinkedResource

	Config *DynamicValue

	// ClientCapabilities defines optionally supported protocol features for the
	// PlanAction RPC, such as forward-compatible Terraform behavior changes.
	ClientCapabilities *PlanActionClientCapabilities
}

// TODO: pkg docs
type ProposedLinkedResource struct {
	PriorState    *DynamicValue
	PlannedState  *DynamicValue
	Config        *DynamicValue
	PriorIdentity *ResourceIdentityData
}

// TODO: pkg docs
type PlanActionResponse struct {
	LinkedResources []*PlannedLinkedResource

	Diagnostics []*Diagnostic

	// Deferred is used to indicate to Terraform that the PlanAction operation
	// needs to be deferred for a reason.
	Deferred *Deferred
}

// TODO: pkg docs
type PlannedLinkedResource struct {
	PlannedState    *DynamicValue
	PlannedIdentity *ResourceIdentityData
}

// TODO: pkg docs
type InvokeActionRequest struct {
	// ActionType is the name of the action being called.
	ActionType string

	LinkedResources []*InvokeLinkedResource

	Config *DynamicValue
}

// TODO: pkg docs
type InvokeLinkedResource struct {
	PriorState      *DynamicValue
	PlannedState    *DynamicValue
	Config          *DynamicValue
	PlannedIdentity *ResourceIdentityData
}

// TODO: pkg docs
type InvokeActionServerStream struct {
	Events iter.Seq[InvokeActionEvent]
}

// TODO: pkg docs
type InvokeActionEvent struct {
	Type InvokeActionEventType
}

// TODO: pkg docs
type InvokeActionEventType interface {
	isInvokeActionEventType() // this interface is only implementable in this package
}

var (
	_ InvokeActionEventType = ProgressInvokeActionEventType{}
	_ InvokeActionEventType = CompletedInvokeActionEventType{}
)

// TODO: pkg docs
type ProgressInvokeActionEventType struct {
	Message string
}

func (a ProgressInvokeActionEventType) isInvokeActionEventType() {}

// TODO: pkg docs
type CompletedInvokeActionEventType struct {
	LinkedResources []*NewLinkedResource
	Diagnostics     []*Diagnostic
}

func (a CompletedInvokeActionEventType) isInvokeActionEventType() {}

// TODO: pkg docs
type NewLinkedResource struct {
	NewState    *DynamicValue
	NewIdentity *ResourceIdentityData

	RequiresReplace bool
}
