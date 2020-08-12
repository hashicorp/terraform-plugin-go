package fromproto

import (
	tfproto "github.com/hashicorp/terraform-plugin-go"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func TerraformTypesRawValue(in tfplugin5.DynamicValue) (tftypes.RawValue, error) {
	// TODO: figure out how to unmarshal DynamicValues into tftypes.RawValue
	return tftypes.RawValue{}, tfproto.ErrUnimplemented
}
