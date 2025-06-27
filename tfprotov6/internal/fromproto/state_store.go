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
	}
}
