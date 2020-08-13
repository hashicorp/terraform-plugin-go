package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func Cty(in tftypes.RawValue) tfplugin5.DynamicValue {
	// TODO: figure out how to marshal tftypes.RawValue into DynamicValue
	return tfplugin5.DynamicValue{}
}

func CtyType(in tftypes.Type) []byte {
	// TODO: figure out how to marshal tftypes.Type into []byte
	return nil
}
