package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

func DynamicValue(in *tfprotov5.DynamicValue) *tfplugin5.DynamicValue {
	return &tfplugin5.DynamicValue{
		Msgpack: in.MsgPack,
		Json:    in.JSON,
	}
}

func CtyType(in tftypes.Type) ([]byte, error) {
	switch {
	case in.Is(tftypes.String), in.Is(tftypes.Bool), in.Is(tftypes.Number),
		in.Is(tftypes.List{}), in.Is(tftypes.Map{}),
		in.Is(tftypes.Set{}), in.Is(tftypes.Object{}),
		in.Is(tftypes.Tuple{}), in.Is(tftypes.DynamicPseudoType):
		return in.MarshalJSON()
	}
	return nil, fmt.Errorf("unknown type %s", in)
}
