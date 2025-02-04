// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import "github.com/hashicorp/terraform-plugin-go/tftypes"

type ResourceIdentitySchema struct {
	// Version indicates which version of the schema this is. Versions
	// should be monotonically incrementing numbers. When Terraform
	// encounters a resource identity stored in state with a schema version
	// lower that the identity schema version the provider advertises for
	// that resource, Terraform requests the provider upgrade the resource's
	// identity state.
	Version int64

	// IdentityAttributes is a list of attributes that uniquely identify a
	// resource. These attributes are used to identify a resource in the
	// state and to import existing resources into the state.
	IdentityAttributes []*ResourceIdentitySchemaAttribute
}

type ResourceIdentitySchemaAttribute struct {
	// Name is the name of the attribute. This is what the user will put
	Name string

	// Type indicates the type of data the attribute expects. See the
	// documentation for the tftypes package for information on what types
	// are supported and their behaviors.
	Type tftypes.Type

	RequiredForImport bool
}
