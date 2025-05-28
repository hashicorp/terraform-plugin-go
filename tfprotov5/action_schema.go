// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import "github.com/hashicorp/terraform-plugin-go/tftypes"

type ActionSchema struct {
	// Version indicates which version of the action schema this is. Versions
	// should be monotonically incrementing numbers. When Terraform
	// encounters an action ... //TODO: what then?
	Version int64

	// Block is the root level of the action schema, the collection of attributes
	// and blocks that make up the actions config
	Block *SchemaBlock

	LinkedResources []*ActionSchemaLinkedResource
}

type ActionSchemaLinkedResource struct {
	Attribute *tftypes.AttributePath

	// Type of the linked resource
	Type string

	LinkedAttributes []*ActionSchemaLinkedResourceLinkedAttribute
}

type ActionSchemaLinkedResourceLinkedAttribute struct {
	Attribute *tftypes.AttributePath
}
