// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

// ActionSchema is how Terraform defines the shape of action data and
// how the practitioner can interact with the action.
type ActionSchema struct {
	// Schema is the definition for the action data itself, which will be specified in an action block in the user's configuration.
	Schema *Schema

	// Type defines how a practitioner can trigger an action, as well as what effect the action can have on the state
	// of the linked managed resources. There is currently only one type of action supported:
	//   - Unlinked actions are actions that cannot cause changes to resource states.
	Type ActionSchemaType
}

// ActionSchemaType is an intentionally unimplementable interface that
// functions as an enum, allowing us to use different strongly-typed action schema types
// that contain additional, but different data, as a generic "action" type.
//
// An action can currently only be one type (Unlinked), which is statically defined in the protocol. Future action types
// will be added to the protocol first, then implemented in this package.
type ActionSchemaType interface {
	isActionSchemaType()
}

var (
	_ ActionSchemaType = UnlinkedActionSchemaType{}
)

// UnlinkedActionSchemaType represents an unlinked action, which cannot cause changes to resource states.
type UnlinkedActionSchemaType struct{}

func (a UnlinkedActionSchemaType) isActionSchemaType() {}
