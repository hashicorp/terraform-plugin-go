// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// ErrUnknownRawIdentityType is returned when a RawIdentity has no Flatmap or JSON
// bytes set. This should never be returned during the normal operation of a
// provider, and indicates one of the following:
//
// 1. terraform-plugin-go is out of sync with the protocol and should be
// updated.
//
// 2. terrafrom-plugin-go has a bug.
//
// 3. The `RawIdentity` was generated or modified by something other than
// terraform-plugin-go and is no longer a valid value.
var ErrUnknownRawIdentityType = errors.New("RawIdentity had no JSON or flatmap data set")

// RawIdentity is the raw, undecoded state for providers to upgrade. It is
// undecoded as Terraform, for whatever reason, doesn't have the previous
// schema available to it, and so cannot decode the state itself and pushes
// that responsibility off onto providers.
type RawIdentity struct {
	JSON []byte
}

// Unmarshal returns a `tftypes.Value` that represents the information
// contained in the RawIdentity in an easy-to-interact-with way. It is the
// main purpose of the RawIdentity type, and is how provider developers should
// obtain state values from the UpgradeResourceIdentity RPC call.
//
// Pass in the type you want the `Value` to be interpreted as. Terraform's type
// system encodes in a lossy manner, meaning the type information is not
// preserved losslessly when going over the wire. Sets, lists, and tuples all
// look the same. Objects and maps all look the same, as well, as do
// user-specified values when DynamicPseudoType is used in the schema.
// Fortunately, the provider should already know the type; it should be the
// type of the schema, or DynamicPseudoType if that's what's in the schema.
// `Unmarshal` will then parse the value as though it belongs to that type, if
// possible, and return a `tftypes.Value` with the appropriate information. If
// the data can't be interpreted as that type, an error will be returned saying
// so. In these cases, double check to make sure the schema is declaring the
// same type being passed into `Unmarshal`.
//
// In the event an ErrUnknownRawIdentityType is returned, one of three things
// has happened:
//
// 1. terraform-plugin-go is out of date and out of sync with the protocol, and
// an issue should be opened on its repo to get it updated.
//
// 2. terraform-plugin-go has a bug somewhere, and an issue should be opened on
// its repo to get it fixed.
//
// 3. The provider or a dependency has modified the `RawIdentity` in an
// unsupported way, or has created one from scratch, and should treat it as
// opaque and not modify it, only calling `Unmarshal` on `RawIdentity`s received
// from RPC requests.
func (s RawIdentity) Unmarshal(typ tftypes.Type) (tftypes.Value, error) {
	if s.JSON != nil {
		return tftypes.ValueFromJSON(s.JSON, typ) //nolint:staticcheck
	}
	return tftypes.Value{}, ErrUnknownRawIdentityType
}

// UnmarshalWithOpts is identical to Unmarshal but also accepts a tftypes.UnmarshalOpts which contains
// options that can be used to modify the behaviour when unmarshalling JSON.
func (s RawIdentity) UnmarshalWithOpts(typ tftypes.Type, opts UnmarshalOpts) (tftypes.Value, error) {
	if s.JSON != nil {
		return tftypes.ValueFromJSONWithOpts(s.JSON, typ, opts.ValueFromJSONOpts) //nolint:staticcheck
	}
	return tftypes.Value{}, ErrUnknownRawIdentityType
}
