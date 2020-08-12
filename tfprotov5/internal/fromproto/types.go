package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func TerraformTypesRawValue(in tfplugin5.DynamicValue) tftypes.RawValue {
	// TODO: figure out how to unmarshal DynamicValues into tftypes.RawValue
	return tftypes.RawValue{}
}

func TerraformTypesType(in []byte) tftypes.Type {
	// TODO: figure out how to unmarshal a cty []byte to tftypes.Type
	var resp tftypes.Type
	return resp
}
