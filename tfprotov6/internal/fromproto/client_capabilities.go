package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ClientCapabilities(in *tfplugin6.ClientCapabilities) *tfprotov6.ClientCapabilities {
	if in == nil {
		return nil
	}

	resp := &tfprotov6.ClientCapabilities{
		DeferralAllowed: in.DeferralAllowed,
	}

	return resp
}
