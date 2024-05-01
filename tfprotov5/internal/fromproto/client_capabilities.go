// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ClientCapabilities(in *tfplugin5.ClientCapabilities) *tfprotov5.ClientCapabilities {
	if in == nil {
		return nil
	}

	resp := &tfprotov5.ClientCapabilities{
		DeferralAllowed: in.DeferralAllowed,
	}

	return resp
}
