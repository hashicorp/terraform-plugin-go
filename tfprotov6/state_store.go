// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

import (
	"context"
	"iter"
)

// StateStoreMetadata describes metadata for a state store in the GetMetadata RPC.
type StateStoreMetadata struct {
	// TypeName is the name of the state store.
	TypeName string
}

// StateStoreServer is an interface containing the methods an list resource
// implementation needs to fill.
type StateStoreServer interface {
	// ValidateStateStoreConfig performs configuration validation
	ValidateStateStoreConfig(context.Context, *ValidateStateStoreConfigRequest) (*ValidateStateStoreConfigResponse, error)

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

type ValidateStateStoreConfigRequest struct {
	TypeName string
	Config   *DynamicValue
}

type ValidateStateStoreConfigResponse struct {
	Diagnostics []*Diagnostic
}

type ConfigureStateStoreRequest struct {
	TypeName     string
	Config       *DynamicValue
	Capabilities *ConfigureStateStoreClientCapabilities
}

type ConfigureStateStoreResponse struct {
	Diagnostics  []*Diagnostic
	Capabilities *StateStoreServerCapabilities
}

type StateStoreServerCapabilities struct {
	ChunkSize int64 // chosen chunk size by plugin
}

type ReadStateBytesRequest struct {
	TypeName string
	StateID  string
}

type ReadStateBytesStream struct {
	Chunks iter.Seq[ReadStateByteChunk]
}

type WriteStateBytesStream struct {
	Chunks iter.Seq2[*WriteStateBytesChunk, []*Diagnostic]
}

// WriteStateBytesChunk contains a chunk of state data, received from Terraform core to be persisted.
type WriteStateBytesChunk struct {
	Meta *WriteStateChunkMeta
	StateByteChunk
}

type WriteStateChunkMeta struct {
	TypeName string
	StateID  string
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
	StateIDs    []string
	Diagnostics []*Diagnostic
}

type DeleteStateRequest struct {
	TypeName string
	StateID  string
}

type DeleteStateResponse struct {
	Diagnostics []*Diagnostic
}

type LockStateRequest struct {
	TypeName  string
	StateID   string
	Operation string
}

type LockStateResponse struct {
	LockID      string
	Diagnostics []*Diagnostic
}

type UnlockStateRequest struct {
	TypeName string
	StateID  string
	LockID   string
}

type UnlockStateResponse struct {
	Diagnostics []*Diagnostic
}
