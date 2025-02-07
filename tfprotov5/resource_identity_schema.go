// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import "github.com/hashicorp/terraform-plugin-go/tftypes"

// ResourceIdentitySchema is the identity schema for a Resource.
type ResourceIdentitySchema struct {
	// IdentityVersion indicates which version of the schema this is. Versions
	// should be monotonically incrementing numbers. When Terraform
	// encounters a resource identity stored in state with a schema version
	// lower that the identity schema version the provider advertises for
	// that resource, Terraform requests the provider upgrade the resource's
	// identity state.
	IdentityVersion int64

	// IdentityAttributes is a list of attributes that uniquely identify a
	// resource. These attributes are used to identify a resource in the
	// state and to import existing resources into the state.
	IdentityAttributes []*ResourceIdentitySchemaAttribute
}

// ResourceIdentitySchemaAttribute represents one value of data within
// resource identity.
// These are always used in resource identity comparisons.
type ResourceIdentitySchemaAttribute struct {
	// Name is the name of the attribute. This is what the user will put
	// before the equals sign to assign a value to this attribute.
	Name string

	// Type indicates the type of data the attribute expects. See the
	// documentation for the tftypes package for information on what types
	// are supported and their behaviors.
	Type tftypes.Type

	// RequiredForImport indicates whether this attribute is required to
	// import the resource. For example it might be false if the value
	// can be derived from provider configuration. Either this or OptionalForImport
	// needs to be true.
	RequiredForImport bool

	// OptionalForImport indicates whether this attribute is optional to
	// import the resource. For example it might be true if the value
	// can be derived from provider configuration. Either this or RequiredForImport
	// needs to be true.
	OptionalForImport bool

	// Description is a human-readable description of the attribute.
	Description string
}
