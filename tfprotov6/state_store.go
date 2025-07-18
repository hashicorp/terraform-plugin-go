// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

import "context"

// StateStoreServer is an interface containing the methods an list resource
// implementation needs to fill.
type StateStoreServer interface {
	// ValidateStateStoreConfig performs configuration validation
	ValidateStateStoreConfig(context.Context, *ValidateStateStoreRequest) (*ValidateStateStoreResponse, error)

	// ConfigureStateStore configures the state store, such as S3 connection in the context of already configured provider
	ConfigureStateStore(context.Context, *ConfigureStateStoreRequest) (*ConfigureStateStoreResponse, error)

	// GetStates returns a list of all states (i.e. CE workspaces) managed by a given state store
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)

	// DeleteState instructs a given state store to delete a specific state (i.e. a CE workspace)
	DeleteState(context.Context, *DeleteStateRequest) (*DeleteStateResponse, error)
}

type ValidateStateStoreRequest struct {
	TypeName string
	Config   *DynamicValue
}

type ValidateStateStoreResponse struct {
	Diagnostics []*Diagnostic
}

type ConfigureStateStoreRequest struct {
	TypeName string
	Config   *DynamicValue
}

type ConfigureStateStoreResponse struct {
	Diagnostics []*Diagnostic
}

type GetStatesRequest struct {
	TypeName string
}

type GetStatesResponse struct {
	StateId     []string
	Diagnostics []*Diagnostic
}

type DeleteStateRequest struct {
	TypeName string
	StateId  string
}

type DeleteStateResponse struct {
	Diagnostics []*Diagnostic
}
