// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

type ProviderClient struct {
	client tfplugin6.ProviderClient
}

func (p *ProviderClient) GetProviderSchema(ctx context.Context) (*tfplugin6.GetProviderSchema_Response, error) {
	req := &tfplugin6.GetProviderSchema_Request{}
	return p.client.GetProviderSchema(ctx, req)
}
