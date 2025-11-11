package tfprotov5

import (
	"context"
)

// GenerateResourceConfigRequest is the request Terraform sends when it wants to generate configuration
// from a resource's state value
type GenerateResourceConfigRequest struct {
	// TODO comment
	TypeName string

	// TODO comment
	State *DynamicValue

	// Mux fills this in
	ResourceSchema *Schema
}

// GenerateResourceConfigResponse TODO
type GenerateResourceConfigResponse struct {
	// TODO comment
	Config *DynamicValue

	// TODO comment
	Diagnostics []*Diagnostic
}

// TODO comment
type GenerateResourceConfigServer interface {
	GenerateResourceConfig(context.Context, *GenerateResourceConfigRequest) (*GenerateResourceConfigResponse, error)
}
