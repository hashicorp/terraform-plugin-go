// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

// ResourceIdentityData contains the raw undecoded identity data
// for a resource.
type ResourceIdentityData struct {
	IdentityData *DynamicValue
}
