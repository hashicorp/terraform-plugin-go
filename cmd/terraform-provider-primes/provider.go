package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"time"
)

type PrimeNumberProvider struct {
}

func (p PrimeNumberProvider) GetMetadata(ctx context.Context, request *tfprotov5.GetMetadataRequest) (*tfprotov5.GetMetadataResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) GetProviderSchema(ctx context.Context, request *tfprotov5.GetProviderSchemaRequest) (*tfprotov5.GetProviderSchemaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) PrepareProviderConfig(ctx context.Context, request *tfprotov5.PrepareProviderConfigRequest) (*tfprotov5.PrepareProviderConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ConfigureProvider(ctx context.Context, request *tfprotov5.ConfigureProviderRequest) (*tfprotov5.ConfigureProviderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) StopProvider(ctx context.Context, request *tfprotov5.StopProviderRequest) (*tfprotov5.StopProviderResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ValidateResourceTypeConfig(ctx context.Context, request *tfprotov5.ValidateResourceTypeConfigRequest) (*tfprotov5.ValidateResourceTypeConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) UpgradeResourceState(ctx context.Context, request *tfprotov5.UpgradeResourceStateRequest) (*tfprotov5.UpgradeResourceStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ReadResource(ctx context.Context, request *tfprotov5.ReadResourceRequest) (*tfprotov5.ReadResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) PlanResourceChange(ctx context.Context, request *tfprotov5.PlanResourceChangeRequest) (*tfprotov5.PlanResourceChangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ApplyResourceChange(ctx context.Context, request *tfprotov5.ApplyResourceChangeRequest) (*tfprotov5.ApplyResourceChangeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ImportResourceState(ctx context.Context, request *tfprotov5.ImportResourceStateRequest) (*tfprotov5.ImportResourceStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) MoveResourceState(ctx context.Context, request *tfprotov5.MoveResourceStateRequest) (*tfprotov5.MoveResourceStateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ValidateDataSourceConfig(ctx context.Context, request *tfprotov5.ValidateDataSourceConfigRequest) (*tfprotov5.ValidateDataSourceConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ReadDataSource(ctx context.Context, request *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) CallFunction(ctx context.Context, request *tfprotov5.CallFunctionRequest) (*tfprotov5.CallFunctionResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) GetFunctions(ctx context.Context, request *tfprotov5.GetFunctionsRequest) (*tfprotov5.GetFunctionsResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ValidateEphemeralResourceConfig(ctx context.Context, request *tfprotov5.ValidateEphemeralResourceConfigRequest) (*tfprotov5.ValidateEphemeralResourceConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) OpenEphemeralResource(ctx context.Context, request *tfprotov5.OpenEphemeralResourceRequest) (*tfprotov5.OpenEphemeralResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) RenewEphemeralResource(ctx context.Context, request *tfprotov5.RenewEphemeralResourceRequest) (*tfprotov5.RenewEphemeralResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) CloseEphemeralResource(ctx context.Context, request *tfprotov5.CloseEphemeralResourceRequest) (*tfprotov5.CloseEphemeralResourceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ValidateListResourceConfig(ctx context.Context, request *tfprotov5.ValidateListResourceConfigRequest) (*tfprotov5.ValidateListResourceConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p PrimeNumberProvider) ListResource(ctx context.Context, request *tfprotov5.ListResourceRequest) (*tfprotov5.ListResourceResponse, error) {
	events := func(yield func(tfprotov5.ListResourceEvent) bool) {
		primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
		for _, prime := range primes {
			displayName := fmt.Sprintf("%d is a prime number", prime)

			resourceObject, err := tfprotov5.NewDynamicValue(tftypes.Number, tftypes.NewValue(tftypes.Number, prime))
			if err != nil {
				panic(err)
			}

			yield(tfprotov5.ListResourceEvent{
				DisplayName:    displayName,
				ResourceObject: &resourceObject,
			})

			fmt.Println("1 second nap")
			time.Sleep(1 * time.Second)
		}
	}

	return &tfprotov5.ListResourceResponse{
		ListResourceEvents: events,
	}, nil
}
