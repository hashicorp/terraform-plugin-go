package tfprotov5

import (
	"context"
)

// ProviderServer is an interface that reflects that Terraform protocol.
// Providers must implement this interface.
type ProviderServer interface {
	// GetProviderSchema is called when Terraform needs to know what the
	// provider's schema is, along with the schemas of all its resources
	// and data sources.
	GetProviderSchema(context.Context, *GetProviderSchemaRequest) (*GetProviderSchemaResponse, error)

	// PrepareProviderConfig is called to give a provider a chance to
	// modify the configuration the user specified before validation.
	PrepareProviderConfig(context.Context, *PrepareProviderConfigRequest) (*PrepareProviderConfigResponse, error)

	// ConfigureProvider is called to pass the user-specified provider
	// configuration to the provider.
	ConfigureProvider(context.Context, *ConfigureProviderRequest) (*ConfigureProviderResponse, error)

	// StopProvider is called when Terraform would like providers to shut
	// down as quickly as possible, and usually represents an interrupt.
	StopProvider(context.Context, *StopProviderRequest) (*StopProviderResponse, error)

	// ResourceServer is an interface encapsulating all the
	// resource-related RPC requests. ProviderServer implementations must
	// implement them, but they are a handy interface for defining what a
	// resource is to terraform-plugin-go, so they're their own interface
	// that is composed into ProviderServer.
	//
	// It is common, but not required, to fill this part of the
	// ProviderServer interface by embeding a ResourceRouter into your
	// ProviderServer implementation and inheriting its implementation.
	ResourceServer

	// DataSourceServer is an interface encapsulating all the data
	// source-related RPC requests. ProviderServer implementations must
	// implement them, but they are a handy interface for defining what a
	// data source is to terraform-plugin-go, so they're their own
	// interface that is composed into ProviderServer.
	//
	// It is common, but not required, to fill this part of the
	// ProviderServer interface by embedding a DataSourceRouter into your
	// ProviderServer implementation and inheriting its implementation.
	DataSourceServer
}

// GetProviderSchemaRequest represents a Terraform RPC request for the
// provider's schemas.
type GetProviderSchemaRequest struct{}

// GetProviderSchemaResponse represents a Terraform RPC response containing the
// provider's schemas.
type GetProviderSchemaResponse struct {
	// Provider defines the schema for the provider configuration, which
	// will be specified in the provider block of the user's configuration.
	Provider *Schema

	// ProviderMeta defines the schema for the provider's metadta, which
	// will be specified in the provider_meta blocks of the terraform block
	// for a module. This is an advanced feature and its usage should be
	// coordinated with the Terraform Core team by opening an issue at
	// https://github.com/hashicorp/terraform/issues/new/choose.
	ProviderMeta *Schema

	// ResourceSchemas is a map of resource names to the schema for the
	// configuration specified in the resource. The name should be a
	// resource name, and should be prefixed with your provider's shortname
	// and an underscore. It should match the first label after `resource`
	// in a user's configuration.
	ResourceSchemas map[string]*Schema

	// DataSourceSchemas is a map of data source names to the schema for
	// the configuration specified in the data source. The name should be a
	// data source name, and should be prefixed with your provider's
	// shortname and an underscore. It should match the first label after
	// `data` in a user's configuration.
	DataSourceSchemas map[string]*Schema

	// Diagnostics report errors or warnings related to returning the
	// provider's schemas. Returning an empty slice indicates success, with
	// no errors or warnings generated.
	Diagnostics []*Diagnostic
}

// PrepareProviderConfigRequest represents a Terraform RPC request for the
// provider to modify the provider configuration in preparation for Terraform
// validating it.
type PrepareProviderConfigRequest struct {
	// Config is the configuration the user supplied for the provider. See
	// the documentation on `DynamicValue` for more information about
	// safely accessing the configuration.
	//
	// TODO: can this configuration contain unknown values?
	Config *DynamicValue
}

// PrepareProviderConfigResponse represents a Terraform RPC response containing
// a modified provider configuration that Terraform can now validate and use.
type PrepareProviderConfigResponse struct {
	// PreparedConfig should be set to the modified configuration. See the
	// documentation on `DynamicValue` for information about safely
	// creating the `DynamicValue`.
	//
	// This RPC call exists because early versions of the Terraform Plugin
	// SDK allowed providers to set defaults for provider configurations in
	// such a way that Terraform couldn't validate the provider config
	// without retrieving the default values first. As providers using
	// terraform-plugin-go directly and new frameworks built on top of it
	// have no such requirement, it is safe and recommended to simply set
	// PreparedConfig to the value of the PrepareProviderConfigRequest's
	// Config property, indicating that no changes are needed to the
	// configuration.
	//
	// TODO: should we provide an implementation that does that that
	// provider developers can just embed and not need to implement the
	// method themselves, then?
	PreparedConfig *DynamicValue

	// Diagnostics report errors or warnings related to preparing the
	// provider's configuration. Returning an empty slice indicates
	// success, with no errors or warnings generated.
	Diagnostics []*Diagnostic
}

// ConfigureProviderRequest represents a Terraform RPC request to supply the
// provider with information about what the user entered in the provider's
// configuration block.
type ConfigureProviderRequest struct {
	// TerraformVersion is the version of Terraform executing the request.
	// This is supplied for logging, analytics, and User-Agent purposes
	// *only*. Providers should not try to gate provider behavior on
	// Terraform versions. It will make you sad. We can't stop you from
	// doing it, but we really highly recommend you do not do it.
	TerraformVersion string

	// Config is the configuration the user supplied for the provider. This
	// information should usually be persisted to the underlying type
	// that's implementing the ProviderServer interface, for use in later
	// RPC requests. See the documentation on `DynamicValue` for more
	// information about safely accessing the configuration.
	//
	// TODO: can this configuration contain unknown values?
	Config *DynamicValue
}

// ConfigureProviderResponse represents a Terraform RPC response to the
// configuration block that Terraform supplied for the provider.
type ConfigureProviderResponse struct {
	// Diagnostics report errors or warnings related to the provider's
	// configuration. Returning an empty slice indicates success, with no
	// errors or warnings generated.
	Diagnostics []*Diagnostic
}

// StopProviderRequest represents a Terraform RPC request to interrupt a
// provider's work and terminate a provider's processes as soon as possible.
type StopProviderRequest struct{}

// StopProviderResponse represents a Terraform RPC response surfacing an issues
// the provider encountered in terminating.
type StopProviderResponse struct {
	// Error should be set to a string describing the error if the provider
	// cannot currently shut down for some reason. Because this always
	// represents a system error and not a user error, it is returned as a
	// string, not a Diagnostic.
	Error string
}
