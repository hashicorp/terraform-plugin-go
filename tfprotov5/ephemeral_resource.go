// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import (
	"context"
	"time"
)

// EphemeralResourceMetadata describes metadata for an ephemeral resource in the GetMetadata
// RPC.
type EphemeralResourceMetadata struct {
	// TypeName is the name of the ephemeral resource.
	TypeName string
}

// EphemeralResourceServer is an interface containing the methods an ephemeral resource
// implementation needs to fill.
type EphemeralResourceServer interface {
	// ValidateEphemeralResourceConfig is called when Terraform is checking that a
	// ephemeral resource's configuration is valid. It is guaranteed to have types
	// conforming to your schema, but it is not guaranteed that all values
	// will be known. This is your opportunity to do custom or advanced
	// validation prior to ephemeral resource creation.
	ValidateEphemeralResourceConfig(context.Context, *ValidateEphemeralResourceConfigRequest) (*ValidateEphemeralResourceConfigResponse, error)

	// OpenEphemeralResource is called when Terraform wants to open the ephemeral resource,
	// usually during planning. If the config for the ephemeral resource contains unknown
	// values, Terraform will defer the OpenEphemeralResource call until apply.
	OpenEphemeralResource(context.Context, *OpenEphemeralResourceRequest) (*OpenEphemeralResourceResponse, error)

	// RenewEphemeralResource is called when Terraform detects that the previously specified
	// RenewAt timestamp has passed. The RenewAt timestamp is supplied either from the
	// OpenEphemeralResource call or a previous RenewEphemeralResource call.
	RenewEphemeralResource(context.Context, *RenewEphemeralResourceRequest) (*RenewEphemeralResourceResponse, error)

	// CloseEphemeralResource is called when Terraform is closing the ephemeral resource at
	// the end of the Terraform run.
	CloseEphemeralResource(context.Context, *CloseEphemeralResourceRequest) (*CloseEphemeralResourceResponse, error)
}

// ValidateEphemeralResourceConfigRequest is the request Terraform sends when it
// wants to validate an ephemeral resource's configuration.
type ValidateEphemeralResourceConfigRequest struct {
	// TypeName is the type of resource Terraform is validating.
	TypeName string

	// Config is the configuration the user supplied for that ephemeral resource. See
	// the documentation on `DynamicValue` for more information about
	// safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	//
	// This configuration may contain unknown values if a user uses
	// interpolation or other functionality that would prevent Terraform
	// from knowing the value at request time. Any attributes not directly
	// set in the configuration will be null.
	Config *DynamicValue
}

// ValidateEphemeralResourceConfigResponse is the response from the provider about
// the validity of an ephemeral resource's configuration.
type ValidateEphemeralResourceConfigResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}

// OpenEphemeralResourceRequest is the request Terraform sends when it
// wants to open an ephemeral resource.
type OpenEphemeralResourceRequest struct {
	// TypeName is the type of resource Terraform is opening.
	TypeName string

	// Config is the configuration the user supplied for that ephemeral resource. See
	// the documentation on `DynamicValue` for more information about
	// safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	//
	// This configuration will always be fully known. If Config contains unknown values,
	// Terraform will defer the OpenEphemeralResource RPC until apply.
	Config *DynamicValue
}

// OpenEphemeralResourceResponse is the response from the provider about the current
// state of the opened ephemeral resource.
type OpenEphemeralResourceResponse struct {
	// State is the provider's understanding of what the ephemeral resource's
	// state is after it has been opened, represented as a `DynamicValue`.
	// See the documentation for `DynamicValue` for information about
	// safely creating the `DynamicValue`.
	//
	// Any attribute, whether computed or not, that has a known value in
	// the Config in the OpenEphemeralResourceRequest must be preserved
	// exactly as it was in State.
	//
	// Any attribute in the Config in the OpenEphemeralResourceRequest
	// that is unknown must take on a known value at this time. No unknown
	// values are allowed in the State.
	//
	// The state should be represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	State *DynamicValue

	// Diagnostics report errors or warnings related to opening the
	// requested ephemeral resource. Returning an empty slice
	// indicates a successful creation with no warnings or errors
	// generated.
	Diagnostics []*Diagnostic

	// Private should be set to any state that the provider would like sent
	// with requests for this ephemeral resource. This state will be associated with
	// the ephemeral resource, but will not be considered when calculating diffs.
	Private []byte

	// RenewAt indicates to Terraform that the ephemeral resource
	// needs to be renewed at the specified time. Terraform will
	// call the RenewEphemeralResource RPC when the specified time has passed.
	RenewAt time.Time

	// IsClosable indicates to Terraform whether the ephemeral resource
	// implements the CloseEphemeralResource RPC.
	IsClosable bool
}

// RenewEphemeralResourceRequest is the request Terraform sends when it
// wants to renew an ephemeral resource.
type RenewEphemeralResourceRequest struct {
	// TypeName is the type of resource Terraform is renewing.
	TypeName string

	// State is the state of the ephemeral resource from the OpenEphemeralResource
	// RPC call. See the documentation on `DynamicValue` for more information
	// about safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	//
	// This prior state will always be fully known.
	State *DynamicValue

	// Private is any provider-defined private state stored with the
	// ephemeral resource. It is used for keeping state with the resource that is not
	// meant to be included when calculating diffs.
	//
	// To ensure private state data is preserved, copy any necessary data to
	// the RenewEphemeralResourceResponse type Private field.
	Private []byte
}

// RenewEphemeralResourceResponse is the response from the provider about the current
// state of the renewed ephemeral resource.
type RenewEphemeralResourceResponse struct {
	// Diagnostics report errors or warnings related to renewing the
	// requested ephemeral resource. Returning an empty slice
	// indicates a successful creation with no warnings or errors
	// generated.
	Diagnostics []*Diagnostic

	// Private should be set to any state that the provider would like sent
	// with requests for this ephemeral resource. This state will be associated with
	// the ephemeral resource, but will not be considered when calculating diffs.
	Private []byte

	// RenewAt indicates to Terraform that the ephemeral resource
	// needs to be renewed at the specified time. Terraform will
	// call the RenewEphemeralResource RPC when the specified time has passed.
	RenewAt time.Time
}

// CloseEphemeralResourceRequest is the request Terraform sends when it
// wants to close an ephemeral resource.
type CloseEphemeralResourceRequest struct {
	// TypeName is the type of resource Terraform is closing.
	TypeName string

	// State is the state of the ephemeral resource from the OpenEphemeralResource
	// RPC call. See the documentation on `DynamicValue` for more information
	// about safely accessing the configuration.
	//
	// The configuration is represented as a tftypes.Object, with each
	// attribute and nested block getting its own key and value.
	//
	// This prior state will always be fully known.
	State *DynamicValue

	// Private is any provider-defined private state stored with the
	// ephemeral resource. It is used for keeping state with the resource that is not
	// meant to be included when calculating diffs.
	Private []byte
}

// CloseEphemeralResourceResponse is the response from the provider about
// the closed ephemeral resource.
type CloseEphemeralResourceResponse struct {
	// Diagnostics report errors or warnings related to closing the
	// requested ephemeral resource. Returning an empty slice
	// indicates a successful creation with no warnings or errors
	// generated.
	Diagnostics []*Diagnostic
}
