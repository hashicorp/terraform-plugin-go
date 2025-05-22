// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func GetMetadata_ListResourceMetadata(in *tfprotov6.ListResourceMetadata) *tfplugin6.GetMetadata_ListResourceMetadata {
	if in == nil {
		return nil
	}

	return &tfplugin6.GetMetadata_ListResourceMetadata{
		TypeName: in.TypeName,
	}
}

func ValidateListResourceConfig_Response(in *tfprotov6.ValidateListResourceConfigResponse) *tfplugin6.ValidateListResourceConfig_Response {
	if in == nil {
		return nil
	}

	return &tfplugin6.ValidateListResourceConfig_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}
