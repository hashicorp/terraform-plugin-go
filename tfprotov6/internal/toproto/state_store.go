// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ValidateStateStoreConfig_Response(in *tfprotov6.ValidateStateStoreResponse) *tfplugin6.ValidateStateStore_Response {
	if in == nil {
		return nil
	}

	return &tfplugin6.ValidateStateStore_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func ConfigureStateStore_Response(in *tfprotov6.ConfigureStateStoreResponse) *tfplugin6.ConfigureStateStore_Response {
	if in == nil {
		return nil
	}

	return &tfplugin6.ConfigureStateStore_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}
