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
