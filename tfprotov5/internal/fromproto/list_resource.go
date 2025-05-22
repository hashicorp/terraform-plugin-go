// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func ValidateListResourceConfigRequest(in *tfplugin5.ValidateListResourceConfig_Request) *tfprotov5.ValidateListResourceConfigRequest {
	if in == nil {
		return nil
	}

	return &tfprotov5.ValidateListResourceConfigRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}
}
