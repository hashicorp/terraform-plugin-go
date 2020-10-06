package tfprotov5

import (
	"context"
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
	Config *DynamicValue
}

type PrepareProviderConfigResponse struct {
	PreparedConfig *DynamicValue
	Diagnostics    []*Diagnostic
}

type ConfigureProviderRequest struct {
	TerraformVersion string
	Config           *DynamicValue
}

type ConfigureProviderResponse struct {
	Diagnostics []*Diagnostic
}

type StopProviderRequest struct{}

type StopProviderResponse struct {
	Error string
}
