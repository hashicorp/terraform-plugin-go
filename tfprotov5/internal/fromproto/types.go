package fromproto

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func DynamicValue(in *tfplugin5.DynamicValue) *tfprotov5.DynamicValue {
	return &tfprotov5.DynamicValue{
		MsgPack: in.Msgpack,
		JSON:    in.Json,
	}
}

func TerraformTypesType(in []byte) (tftypes.Type, error) {
	var raw interface{}
	err := json.Unmarshal(in, &raw)
	if err != nil {
		return nil, fmt.Errorf("error parsing type, not valid JSON: %w", err)
	}
	return tftypes.ParseType(raw)
}
