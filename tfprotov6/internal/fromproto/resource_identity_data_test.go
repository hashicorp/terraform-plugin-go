// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
)

func testTfplugin6ResourceIdentityData() *tfplugin6.ResourceIdentityData {
	return toproto.ResourceIdentityData(testTfprotov6ResourceIdentityData())
}

func testTfprotov6ResourceIdentityData() *tfprotov6.ResourceIdentityData {
	return &tfprotov6.ResourceIdentityData{
		IdentityData: testTfprotov6DynamicValue(),
	}
}
