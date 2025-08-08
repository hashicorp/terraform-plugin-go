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

func GetStates_Response(in *tfprotov6.GetStatesResponse) *tfplugin6.GetStates_Response {
	if in == nil {
		return nil
	}

	return &tfplugin6.GetStates_Response{
		StateId:     in.StateId,
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}

func DeleteState_Response(in *tfprotov6.DeleteStateResponse) *tfplugin6.DeleteState_Response {
	if in == nil {
		return nil
	}

	return &tfplugin6.DeleteState_Response{
		Diagnostics: Diagnostics(in.Diagnostics),
	}
}
