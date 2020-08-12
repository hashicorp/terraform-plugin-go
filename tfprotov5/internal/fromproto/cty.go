package fromproto

import (
	tfproto "github.com/hashicorp/terraform-plugin-go"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func CtyBlock(in tfplugin5.DynamicValue) (tfprotov5.CtyBlock, error) {
	// TODO: figure out how to do this
	return tfprotov5.CtyBlock{}, tfproto.ErrUnimplemented
}
