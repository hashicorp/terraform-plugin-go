package fromproto

import (
	"github.com/hashicorp/terraform-protocol-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-protocol-go/tfprotov5"
)

func RawState(in tfplugin5.RawState) tfprotov5.RawState {
	return tfprotov5.RawState{
		JSON:    in.Json,
		Flatmap: in.Flatmap,
	}
}
