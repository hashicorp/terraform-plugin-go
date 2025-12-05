// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func testTfplugin5ResourceIdentityData() *tfplugin5.ResourceIdentityData {
	return toproto.ResourceIdentityData(testTfprotov5ResourceIdentityData())
}

func testTfprotov5ResourceIdentityData() *tfprotov5.ResourceIdentityData {
	return &tfprotov5.ResourceIdentityData{
		IdentityData: testTfprotov5DynamicValue(),
	}
}
