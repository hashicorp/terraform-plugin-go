// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func testTfplugin6DynamicValue() *tfplugin6.DynamicValue {
	return toproto.DynamicValue(testTfprotov6DynamicValue())
}

func testTfprotov6DynamicValue() *tfprotov6.DynamicValue {
	dynamicValue, err := tfprotov6.NewDynamicValue(
		tftypes.Object{},
		tftypes.NewValue(tftypes.Object{}, nil),
	)

	if err != nil {
		panic("unable to create DynamicValue: " + err.Error())
	}

	return &dynamicValue
}
