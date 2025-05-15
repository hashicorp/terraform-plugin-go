// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
)

var _ plugin.GRPCPlugin = (*GRPCPluginClient)(nil)

type GRPCPluginClient struct {
	plugin.Plugin
	ClientFunc func(*grpc.ClientConn) any
}

func (p *GRPCPluginClient) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return p.ClientFunc(c), nil
}

func (p *GRPCPluginClient) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	return nil
}
