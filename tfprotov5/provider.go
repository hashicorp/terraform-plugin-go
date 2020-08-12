package tfprotov5

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tftypes"
)

type ProviderServer interface {
	GetProviderSchema(context.Context, *GetProviderSchemaRequest) (*GetProviderSchemaResponse, error)
	PrepareProviderConfig(context.Context, *PrepareProviderConfigRequest) (*PrepareProviderConfigResponse, error)
	ConfigureProvider(context.Context, *ConfigureProviderRequest) (*ConfigureProviderResponse, error)
	StopProvider(context.Context, *StopProviderRequest) (*StopProviderResponse, error)

	ResourceServer
	DataSourceServer
}

type GetProviderSchemaRequest struct{}

type GetProviderSchemaResponse struct {
	Provider          *Schema
	ProviderMeta      *Schema
	ResourceSchemas   map[string]*Schema
	DataSourceSchemas map[string]*Schema
	Diagnostics       []*Diagnostic
}

type PrepareProviderConfigRequest struct {
	Config *tftypes.RawValue
}

type PrepareProviderConfigResponse struct {
	PreparedConfig *tftypes.RawValue
	Diagnostics    []*Diagnostic
}

type ConfigureProviderRequest struct {
	TerraformVersion string
	Config           *tftypes.RawValue
}

type ConfigureProviderResponse struct {
	Diagnostics []*Diagnostic
}

type StopProviderRequest struct{}

type StopProviderResponse struct {
	Error string
}
