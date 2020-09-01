package fromproto

import (
	"errors"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

var ErrUnknownDynamicValueType = errors.New("DynamicValue had no JSON or msgpack data set")

func TerraformTypesRawValue(in tfplugin5.DynamicValue) (tftypes.RawValue, error) {
	if len(in.Msgpack) > 0 {
		return msgpackToRawValue(in.Msgpack)
	}
	if len(in.Json) > 0 {
		return jsonToRawValue(in.Json)
	}
	return tftypes.RawValue{}, ErrUnknownDynamicValueType
}

func msgpackToRawValue(in []byte) (tftypes.RawValue, error) {
	// TODO: parse msgpack
	return tftypes.RawValue{}, nil
}

func jsonToRawValue(in []byte) (tftypes.RawValue, error) {
	// TODO: parse json
	return tftypes.RawValue{}, nil
}

func TerraformTypesType(in []byte) tftypes.Type {
	// TODO: figure out how to unmarshal a cty []byte to tftypes.Type
	var resp tftypes.Type
	return resp
}
