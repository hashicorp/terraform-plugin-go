// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func testTfprotov5ResourceIdentityData() *tfprotov5.ResourceIdentityData {
	return fromproto.ResourceIdentityData(testTfplugin5ResourceIdentityData())
}

func testTfplugin5ResourceIdentityData() *tfplugin5.ResourceIdentityData {
	return &tfplugin5.ResourceIdentityData{
		IdentityData: testTfplugin5DynamicValue(),
	}
}
