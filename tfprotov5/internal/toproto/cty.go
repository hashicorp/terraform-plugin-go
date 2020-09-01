package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func Cty(in tftypes.RawValue) (tfplugin5.DynamicValue, error) {
	// TODO: figure out how to marshal tftypes.RawValue into DynamicValue
	// always populate msgpack, as per https://github.com/hashicorp/terraform/blob/doc-provider-value-wire-protocol/docs/plugin-protocol/object-wire-format.md
	return tfplugin5.DynamicValue{}, nil
}

func CtyType(in tftypes.Type) []byte {
	// TODO: figure out how to marshal tftypes.Type into []byte
	return nil
}
