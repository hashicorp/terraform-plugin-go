// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

// ClientCapabilities allows Terraform to publish information regarding optionally supported
// protocol features, such as forward-compatible Terraform behavior changes.
type ClientCapabilities struct {
	// DeferralAllowed signals that the Terraform client is able to
	// handle deferred responses from the provider.
	DeferralAllowed bool
}
