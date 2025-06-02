// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ValidateListResourceConfigRequest(in *tfplugin6.ValidateListResourceConfig_Request) *tfprotov6.ValidateListResourceConfigRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.ValidateListResourceConfigRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}
}
