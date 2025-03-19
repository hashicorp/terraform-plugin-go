// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ActionResourceChanges(in map[string]*tfprotov6.DynamicValue) map[string]*tfplugin6.DynamicValue {
	resp := make(map[string]*tfplugin6.DynamicValue)

	for k, v := range in {
		resp[k] = DynamicValue(v)
	}

	return resp
}
