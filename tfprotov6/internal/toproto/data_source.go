// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func GetMetadata_DataSourceMetadata(in *tfprotov6.DataSourceMetadata) *tfplugin6.GetMetadata_DataSourceMetadata {
	if in == nil {
		return nil
	}

	return &tfplugin6.GetMetadata_DataSourceMetadata{
		TypeName: in.TypeName,
	}
}

func ValidateDataResourceConfig_Response(in *tfprotov6.ValidateDataResourceConfigResponse) (*tfplugin6.ValidateDataResourceConfig_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ValidateDataResourceConfig_Response{
		Diagnostics: diags,
	}

	return resp, nil
}

func ReadDataSource_Response(in *tfprotov6.ReadDataSourceResponse) (*tfplugin6.ReadDataSource_Response, error) {
	if in == nil {
		return nil, nil
	}

	diags, err := Diagnostics(in.Diagnostics)

	if err != nil {
		return nil, err
	}

	resp := &tfplugin6.ReadDataSource_Response{
		Diagnostics: diags,
		State:       DynamicValue(in.State),
	}

	return resp, nil
}
