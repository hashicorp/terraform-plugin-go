// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ValidateStateStoreRequest(in *tfplugin6.ValidateStateStore_Request) *tfprotov6.ValidateStateStoreRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.ValidateStateStoreRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
	}
}

func ConfigureStateStoreRequest(in *tfplugin6.ConfigureStateStore_Request) *tfprotov6.ConfigureStateStoreRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.ConfigureStateStoreRequest{
		TypeName: in.TypeName,
		Config:   DynamicValue(in.Config),
		Capabilities: tfprotov6.StateStoreClientCapabilities{
			ChunkSize: in.Capabilities.ChunkSize,
		},
	}
}

func ReadStateBytesRequest(in *tfplugin6.ReadStateBytes_Request) *tfprotov6.ReadStateBytesRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.ReadStateBytesRequest{
		TypeName: in.TypeName,
		StateId:  in.StateId,
	}
}

func GetStatesRequest(in *tfplugin6.GetStates_Request) *tfprotov6.GetStatesRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.GetStatesRequest{
		TypeName: in.TypeName,
	}
}

func DeleteStateRequest(in *tfplugin6.DeleteState_Request) *tfprotov6.DeleteStateRequest {
	if in == nil {
		return nil
	}

	return &tfprotov6.DeleteStateRequest{
		TypeName: in.TypeName,
		StateId:  in.StateId,
	}
}
