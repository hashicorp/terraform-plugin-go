// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func testTfprotov6ResourceIdentityData() *tfprotov6.ResourceIdentityData {
	return fromproto.ResourceIdentityData(testTfplugin6ResourceIdentityData())
}

func testTfplugin6ResourceIdentityData() *tfplugin6.ResourceIdentityData {
	return &tfplugin6.ResourceIdentityData{
		IdentityData: testTfplugin6DynamicValue(),
	}
}
