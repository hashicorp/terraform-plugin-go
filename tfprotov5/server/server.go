package tfprotov5server

import (
	"context"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

type ServeOpt interface {
	ApplyServeOpt(*ServeConfig) error
}

type ServeConfig struct {
	logger       hclog.Logger
	debugCtx     context.Context
	debugCh      chan *plugin.ReattachConfig
	debugCloseCh chan struct{}
}

type serveConfigFunc func(*ServeConfig) error

func (s serveConfigFunc) ApplyServeOpt(in *ServeConfig) error {
	return s(in)
}

func WithDebug(ctx context.Context, config chan *plugin.ReattachConfig, closeCh chan struct{}) ServeOpt {
	return serveConfigFunc(func(in *ServeConfig) error {
		in.debugCtx = ctx
		in.debugCh = config
		in.debugCloseCh = closeCh
		return nil
	})
}

func WithGoPluginLogger(logger hclog.Logger) ServeOpt {
	return serveConfigFunc(func(in *ServeConfig) error {
		in.logger = logger
		return nil
	})
}

func Serve(name string, serverFactory func() tfprotov5.ProviderServer, opts ...ServeOpt) error {
	var conf ServeConfig
	for _, opt := range opts {
		err := opt.ApplyServeOpt(&conf)
		if err != nil {
			return err
		}
	}
	serveConfig := &plugin.ServeConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  5,
			MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: "d602bf8f470bc67ca7faa0386276bbdd4330efaf76d1a219cb4d6991ca9872b2",
		},
		Plugins: plugin.PluginSet{
			"provider": &GRPCProviderPlugin{
				GRPCProvider: serverFactory,
			},
		},
		GRPCServer: plugin.DefaultGRPCServer,
	}
	if conf.logger != nil {
		serveConfig.Logger = conf.logger
	}
	if conf.debugCh != nil {
		serveConfig.Test = &plugin.ServeTestConfig{
			Context:          conf.debugCtx,
			ReattachConfigCh: conf.debugCh,
			CloseCh:          conf.debugCloseCh,
		}
	}
	plugin.Serve(serveConfig)
	return nil
}

type server struct {
	downstream tfprotov5.ProviderServer
	tfplugin5.UnimplementedProviderServer
}

func New(serve tfprotov5.ProviderServer) tfplugin5.ProviderServer {
	return server{
		downstream: serve,
	}
}

func (s server) GetSchema(ctx context.Context, req *tfplugin5.GetProviderSchema_Request) (*tfplugin5.GetProviderSchema_Response, error) {
	r, err := fromproto.GetProviderSchemaRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.GetProviderSchema(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.GetProviderSchema_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) PrepareProviderConfig(ctx context.Context, req *tfplugin5.PrepareProviderConfig_Request) (*tfplugin5.PrepareProviderConfig_Response, error) {
	r, err := fromproto.PrepareProviderConfigRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.PrepareProviderConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.PrepareProviderConfig_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) Configure(ctx context.Context, req *tfplugin5.Configure_Request) (*tfplugin5.Configure_Response, error) {
	r, err := fromproto.ConfigureProviderRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ConfigureProvider(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.Configure_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) Stop(ctx context.Context, req *tfplugin5.Stop_Request) (*tfplugin5.Stop_Response, error) {
	r, err := fromproto.StopProviderRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.StopProvider(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.Stop_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ValidateDataSourceConfig(ctx context.Context, req *tfplugin5.ValidateDataSourceConfig_Request) (*tfplugin5.ValidateDataSourceConfig_Response, error) {
	r, err := fromproto.ValidateDataSourceConfigRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ValidateDataSourceConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ValidateDataSourceConfig_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ReadDataSource(ctx context.Context, req *tfplugin5.ReadDataSource_Request) (*tfplugin5.ReadDataSource_Response, error) {
	r, err := fromproto.ReadDataSourceRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ReadDataSource(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ReadDataSource_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ValidateResourceTypeConfig(ctx context.Context, req *tfplugin5.ValidateResourceTypeConfig_Request) (*tfplugin5.ValidateResourceTypeConfig_Response, error) {
	r, err := fromproto.ValidateResourceTypeConfigRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ValidateResourceTypeConfig(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ValidateResourceTypeConfig_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) UpgradeResourceState(ctx context.Context, req *tfplugin5.UpgradeResourceState_Request) (*tfplugin5.UpgradeResourceState_Response, error) {
	r, err := fromproto.UpgradeResourceStateRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.UpgradeResourceState(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.UpgradeResourceState_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ReadResource(ctx context.Context, req *tfplugin5.ReadResource_Request) (*tfplugin5.ReadResource_Response, error) {
	r, err := fromproto.ReadResourceRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ReadResource(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ReadResource_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) PlanResourceChange(ctx context.Context, req *tfplugin5.PlanResourceChange_Request) (*tfplugin5.PlanResourceChange_Response, error) {
	r, err := fromproto.PlanResourceChangeRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.PlanResourceChange(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.PlanResourceChange_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ApplyResourceChange(ctx context.Context, req *tfplugin5.ApplyResourceChange_Request) (*tfplugin5.ApplyResourceChange_Response, error) {
	r, err := fromproto.ApplyResourceChangeRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ApplyResourceChange(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ApplyResourceChange_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (s server) ImportResourceState(ctx context.Context, req *tfplugin5.ImportResourceState_Request) (*tfplugin5.ImportResourceState_Response, error) {
	r, err := fromproto.ImportResourceStateRequest(req)
	if err != nil {
		return nil, err
	}
	resp, err := s.downstream.ImportResourceState(ctx, r)
	if err != nil {
		return nil, err
	}
	ret, err := toproto.ImportResourceState_Response(resp)
	if err != nil {
		return nil, err
	}
	return ret, nil
}