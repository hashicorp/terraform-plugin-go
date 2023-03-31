package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/pulumi/terraform/pkg/tfplugin6"
)

func RawState(in *tfplugin6.RawState) *tfprotov6.RawState {
	return &tfprotov6.RawState{
		JSON:    in.Json,
		Flatmap: in.Flatmap,
	}
}
