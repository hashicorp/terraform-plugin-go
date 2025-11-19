// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

// StateStoreSchema is how Terraform defines the shape of state store data.
type StateStoreSchema struct {
	// Schema is the definition for the state store data itself, which will be specified in an state store block in the user's configuration.
	Schema *Schema
}
