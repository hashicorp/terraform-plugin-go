// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov6

const (
	// TODO: doc
	DeferredReasonUnknown DeferredReason = 0

	// TODO: doc
	DeferredReasonResourceConfigUnknown DeferredReason = 1

	// TODO: doc
	DeferredReasonProviderConfigUnknown DeferredReason = 2

	// TODO: doc
	DeferredReasonAbsentPrereq DeferredReason = 3
)

// TODO: doc
type Deferred struct {
	// TODO: doc
	Reason DeferredReason
}

// TODO: doc
type DeferredReason int32

func (d DeferredReason) String() string {
	switch d {
	case 0:
		return "UNKNOWN"
	case 1:
		return "RESOURCE_CONFIG_UNKNOWN"
	case 2:
		return "PROVIDER_CONFIG_UNKNOWN"
	case 3:
		return "ABSENT_PREREQ"
	}
	return "UNKNOWN"
}
