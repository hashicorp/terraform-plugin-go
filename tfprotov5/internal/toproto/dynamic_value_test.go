// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func testTfplugin5DynamicValue() *tfplugin5.DynamicValue {
	return toproto.DynamicValue(testTfprotov5DynamicValue())
}

func testTfprotov5DynamicValue() *tfprotov5.DynamicValue {
	dynamicValue, err := tfprotov5.NewDynamicValue(
		tftypes.Object{},
		tftypes.NewValue(tftypes.Object{}, nil),
	)

	if err != nil {
		panic("unable to create DynamicValue: " + err.Error())
	}

	return &dynamicValue
}
