package tfprotov5

import (
	"bytes"
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
	"github.com/vmihailenco/msgpack"
)

var ErrUnknownDynamicValueType = errors.New("DynamicValue had no JSON or msgpack data set")

// DynamicValue represents a nested encoding value that came from the protocol.
// The only way providers should ever interact with it is by calling its
// `Unmarshal` method to retrive a `tftypes.Value`. Although the type system
// allows for other interactions, they are explicitly not supported, and will
// not be considered when evaluation for breaking changes. Treat this type as
// an opaque value, and *only* call its `Unmarshal` method.
type DynamicValue struct {
	MsgPack []byte
	JSON    []byte
}

// Unmarshal returns a tftypes.Value that represents the information contained
// in the DynamicValue in an easy-to-interact-with way. It is the main purpose
// of the DynamicValue type, and is how provider developers should obtain
// config, state, and other values from the protocol.
//
// Pass in the type you want the value to be interpreted as. Terraform's type
// system encodes in a lossy manner, meaning the type information is not
// preserved losslessly when going over the wire. Sets, lists, and tuples all
// look the same, as do user-specified values when the provider has a
// DynamicPseudoType in its schema. Objects and maps all look the same, as
// well, as do PseudoDynamicType values sometimes. Fortunately, the provider
// should already know the type; it should be the type of the schema, or
// PseudoDynamicType if that's what's in the schema. `Unmarshal` will then
// parse the value as though it belongs to that type, if possible, and return a
// tftypes.Value with the appropriate information. If the data can't be
// interpreted as that type, an error will be returned saying so. In these
// cases, double check to make sure the schema is declaring the same type being
// passed into Unmarshal.
//
// In the event an ErrUnknownDynamicValueType is returned, one of three things
// has happened:
//
// 1. terraform-plugin-go is out of date and out of sync with the protocol, and
// an issue should be opened on its repo to get it updated.
//
// 2. terraform-plugin-go has a bug somewhere, and an issue should be opened on
// its repo to get it fixed.
//
// 3. The provider or a dependency has modified the DynamicValue in an
// unsupported way, and should treat it as opaque and not modify it, only call
// Unmarshal.
func (d DynamicValue) Unmarshal(typ tftypes.Type) (tftypes.Value, error) {
	if d.JSON != nil {
		return jsonUnmarshal(d.JSON, typ, nil)
	}
	if d.MsgPack != nil {
		r := bytes.NewReader(d.MsgPack)
		dec := msgpack.NewDecoder(r)
		return msgpackUnmarshal(dec, typ, nil)
	}
	return tftypes.Value{}, ErrUnknownDynamicValueType
}
