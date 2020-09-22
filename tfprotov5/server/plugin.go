package tfprotov5server

import (
	"context"
	"errors"
	"net/rpc"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"google.golang.org/grpc"
)

type GRPCProviderPlugin struct {
	GRPCProvider func() tfprotov5.ProviderServer
}

func (p *GRPCProviderPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers.")
}

func (p *GRPCProviderPlugin) Client(*plugin.MuxBroker, *rpc.Client) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers.")
}

func (p *GRPCProviderPlugin) GRPCClient(context.Context, *plugin.GRPCBroker, *grpc.ClientConn) (interface{}, error) {
	return nil, errors.New("terraform-plugin-go only implements gRPC servers.")
}

func (p *GRPCProviderPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	tfplugin5.RegisterProviderServer(s, New(p.GRPCProvider()))
	return nil
}
