package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

type PrimeNumberProvider struct {
}

func (p PrimeNumberProvider) GetMetadata(ctx context.Context, request *tfprotov5.GetMetadataRequest) (*tfprotov5.GetMetadataResponse, error) {
	return &tfprotov5.GetMetadataResponse{
		ListResources: []tfprotov5.ListResourceMetadata{
			{
				TypeName: "prime",
			},
		},
	}, nil
}

func (p PrimeNumberProvider) GetResourceIdentitySchemas(ctx context.Context, request *tfprotov5.GetResourceIdentitySchemasRequest) (*tfprotov5.GetResourceIdentitySchemasResponse, error) {
	return &tfprotov5.GetResourceIdentitySchemasResponse{
		IdentitySchemas: map[string]*tfprotov5.ResourceIdentitySchema{
			"prime": &tfprotov5.ResourceIdentitySchema{
				IdentityAttributes: []*tfprotov5.ResourceIdentitySchemaAttribute{
					{
						Name:              "number",
						Type:              tftypes.Number,
						RequiredForImport: true,
					},
				},
			},
		},
	}, nil
}

func (p PrimeNumberProvider) GetProviderSchema(ctx context.Context, request *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	return &tfprotov5.GetProviderSchemaResponse{
		Provider: &tfprotov5.Schema{
			Block: &tfprotov5.SchemaBlock{
				Description:     "Prime Provider",
				DescriptionKind: tfprotov5.StringKindPlain,
			},
		},
		ResourceSchemas: map[string]*tfprotov5.Schema{
			"prime": p.primeSchema(),
		},
		ListResourceSchemas: map[string]*tfprotov5.Schema{
			"prime": p.primeSchema(),
		},
	}, nil
}

func (p PrimeNumberProvider) primeSchema() *tfprotov5.Schema {
	return &tfprotov5.Schema{
		Block: &tfprotov5.SchemaBlock{
			Attributes: []*tfprotov5.SchemaAttribute{
				{
					Name:            "number",
					Type:            tftypes.Number,
					Description:     "The nth prime",
					DescriptionKind: tfprotov5.StringKindPlain,
					Optional:        true,
				},
				{
					Name:            "ordinal",
					Type:            tftypes.Number,
					Description:     "n",
					DescriptionKind: tfprotov5.StringKindPlain,
					Optional:        true,
				},
			},
		},
	}
}

func (p PrimeNumberProvider) PrepareProviderConfig(ctx context.Context, request *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	return &tfprotov5.PrepareProviderConfigResponse{
		PreparedConfig: request.Config,
	}, nil
}

func (p PrimeNumberProvider) ConfigureProvider(ctx context.Context, request *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	return &tfprotov5.ConfigureProviderResponse{}, nil
}
func (p PrimeNumberProvider) ListResource(ctx context.Context, request *tfprotov5.ListResourceRequest) (*tfprotov5.ListResourceResponse, error) {
	events := func(yield func(tfprotov5.ListResourceEvent) bool) {
		primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

		var suffix string
		for i, prime := range primes {
			ordinal := i + 1
			switch {
			case ordinal == 1, ordinal >= 20 && ordinal%10 == 1:
				suffix = "st"
			case ordinal == 2, ordinal >= 20 && ordinal%10 == 2:
				suffix = "nd"
			case ordinal == 3, ordinal >= 20 && ordinal%10 == 3:
				suffix = "rd"
			default:
				suffix = "th"
			}
			displayName := fmt.Sprintf("The %d%s prime is %d", ordinal, suffix, prime)

			typ := p.primeSchema().ValueType()
			value := map[string]tftypes.Value{
				"number":  tftypes.NewValue(tftypes.Number, prime),
				"ordinal": tftypes.NewValue(tftypes.Number, ordinal),
			}

			resourceObject, err := tfprotov5.NewDynamicValue(typ, tftypes.NewValue(typ, value))
			if err != nil {
				panic(err)
			}

			protoEv := tfprotov5.ListResourceEvent{
				DisplayName:    displayName,
				ResourceObject: &resourceObject,
			}
			if !yield(protoEv) {
				fmt.Println("let's stop here")
				return
			}

			fmt.Println("1 second nap")
			time.Sleep(1 * time.Second)
		}
	}

	return &tfprotov5.ListResourceResponse{
		ListResourceEvents: events,
	}, nil
}

func (p PrimeNumberProvider) ValidateListResourceConfig(ctx context.Context, request *tfprotov5.ValidateListResourceConfigRequest) (*tfprotov5.ValidateListResourceConfigResponse, error) {
	return &tfprotov5.ValidateListResourceConfigResponse{}, nil
}
