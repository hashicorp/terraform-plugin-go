// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

import (
	"context"
	"iter"
)

// StateStoreServer is an interface containing the methods an list resource
// implementation needs to fill.
type StateStoreServer interface {
	// ValidateStateStoreConfig performs configuration validation
	ValidateStateStoreConfig(context.Context, *ValidateStateStoreRequest) (*ValidateStateStoreResponse, error)

	// ConfigureStateStore configures the state store, such as S3 connection in the context of already configured provider
	ConfigureStateStore(context.Context, *ConfigureStateStoreRequest) (*ConfigureStateStoreResponse, error)

	// ReadStateBytes streams byte chunks of a given state file from a state store
	ReadStateBytes(context.Context, *ReadStateBytesRequest) (*ReadStateBytesStream, error)

	WriteStateBytes(context.Context, *WriteStateBytesStream) (*WriteStateBytesResponse, error)

	// GetStates returns a list of all states (i.e. CE workspaces) managed by a given state store
	GetStates(context.Context, *GetStatesRequest) (*GetStatesResponse, error)

	// DeleteState instructs a given state store to delete a specific state (i.e. a CE workspace)
	DeleteState(context.Context, *DeleteStateRequest) (*DeleteStateResponse, error)

	// LockState instructs a given state store to lock a specific state (i.e. a CE workspace)
	LockState(context.Context, *LockStateRequest) (*LockStateResponse, error)

	// UnlockState instructs a given state store to unlock a specific state (i.e. a CE workspace)
	UnlockState(context.Context, *UnlockStateRequest) (*UnlockStateResponse, error)
}

type ValidateStateStoreRequest struct {
	TypeName string
	Config   *DynamicValue
}

type ValidateStateStoreResponse struct {
	Diagnostics []*Diagnostic
}

type ConfigureStateStoreRequest struct {
	TypeName     string
	Config       *DynamicValue
	Capabilities StateStoreClientCapabilities
}

type ConfigureStateStoreResponse struct {
	Diagnostics  []*Diagnostic
	Capabilities StateStoreServerCapabilities
}

type StateStoreClientCapabilities struct {
	ChunkSize int64 // suggested chunk size by Core
}

type StateStoreServerCapabilities struct {
	ChunkSize int64 // chosen chunk size by plugin
}

type ReadStateBytesRequest struct {
	TypeName string
	StateId  string
}

type ReadStateBytesStream struct {
	Chunks iter.Seq[ReadStateByteChunk]
}

type WriteStateBytesStream struct {
	Chunks iter.Seq[WriteStateBytesChunk]
}

// WriteStateBytesChunk contains:
//  1. A chunk of state data, received from Terraform core to be persisted.
//  2. Any gRPC-related errors the provider server encountered when
//     receiving data from Terraform core.
//
// If a gRPC error is set, then the chunk should be empty.
type WriteStateBytesChunk struct {
	Meta *WriteStateChunkMeta
	StateByteChunk
	Err error
}

type WriteStateChunkMeta struct {
	TypeName string
	StateId  string
}

type WriteStateBytesResponse struct {
	Diagnostics []*Diagnostic
}

type ReadStateByteChunk struct {
	StateByteChunk
	Diagnostics []*Diagnostic
}

type StateByteChunk struct {
	Bytes       []byte
	TotalLength int64
	Range       StateByteRange
}

type StateByteRange struct {
	Start, End int64
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

type LockStateRequest struct {
	TypeName  string
	StateId   string
	Operation string
}

type LockStateResponse struct {
	LockId      string
	Diagnostics []*Diagnostic
}

type UnlockStateRequest struct {
	TypeName string
	StateId  string
	LockId   string
}

type UnlockStateResponse struct {
	Diagnostics []*Diagnostic
}
