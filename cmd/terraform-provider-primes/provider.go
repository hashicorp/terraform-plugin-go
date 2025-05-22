package main

import (
	"context"
	"fmt"
	"iter"
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
			"prime": {
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

func (p PrimeNumberProvider) ordinalSuffixes() iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 1; ; i++ {
			switch {
			case i == 1, i >= 20 && i%10 == 1:
				yield("st")
			case i == 2, i >= 20 && i%10 == 2:
				yield("nd")
			case i == 3, i >= 20 && i%10 == 3:
				yield("rd")
			default:
				yield("th")
			}

		}
	}
}

func (p PrimeNumberProvider) convertToDynamicValue(number int, ordinal int) (tfprotov5.DynamicValue, error) {
	typ := p.primeSchema().ValueType()
	value := map[string]tftypes.Value{
		"number":  tftypes.NewValue(tftypes.Number, number),
		"ordinal": tftypes.NewValue(tftypes.Number, ordinal),
	}

	return tfprotov5.NewDynamicValue(typ, tftypes.NewValue(typ, value))
}

func (p PrimeNumberProvider) ListResource(ctx context.Context, request *tfprotov5.ListResourceRequest) (*tfprotov5.ListResourceResponse, error) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	nextSuffix, stop := iter.Pull(p.ordinalSuffixes())
	defer stop()

	events := func(yield func(tfprotov5.ListResourceEvent) bool) {
		for i, prime := range primes {
			ordinal := i + 1
			suffix, _ := nextSuffix()
			displayName := fmt.Sprintf("The %d%s prime is %d", ordinal, suffix, prime)

			resourceObject, err := p.convertToDynamicValue(prime, ordinal)
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

func (p PrimeNumberProvider) StopProvider(ctx context.Context, request *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	return &tfprotov5.StopProviderResponse{}, nil
}
