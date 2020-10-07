package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func RawState(in *tfprotov5.RawState) *tfplugin5.RawState {
	return &tfplugin5.RawState{
		Json:    in.JSON,
		Flatmap: in.Flatmap,
	}
}
