// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

import (
	"context"
)

// ListResourceMetadata describes metadata for an list resource in the GetMetadata
// RPC.
type ListResourceMetadata struct {
	// TypeName is the name of the list resource.
	TypeName string
}

// ListResourceServer is an interface containing the methods an list resource
// implementation needs to fill.
type ListResourceServer interface {
	// ValidateListResourceConfig is called when Terraform is checking that an
	// list resource configuration is valid. It is guaranteed to have types
	// conforming to your schema, but it is not guaranteed that all values
	// will be known. This is your opportunity to do custom or advanced
	// validation prior to an list resource being opened.
	ValidateListResourceConfig(context.Context, *ValidateListResourceConfigRequest) (*ValidateListResourceConfigResponse, error)
}

// ValidateListResourceConfigRequest is the request Terraform sends when it
// wants to validate an list resource's configuration.
type ValidateListResourceConfigRequest struct {
	// TypeName is the type of resource Terraform is validating.
	TypeName string

	// Config is the configuration the user supplied for that list resource. See
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

// ValidateListResourceConfigResponse is the response from the provider about
// the validity of an list resource's configuration.
type ValidateListResourceConfigResponse struct {
	// Diagnostics report errors or warnings related to the given
	// configuration. Returning an empty slice indicates a successful
	// validation with no warnings or errors generated.
	Diagnostics []*Diagnostic
}
