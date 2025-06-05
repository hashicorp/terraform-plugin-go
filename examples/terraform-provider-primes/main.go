// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"
)

func main() {
	providerFn := func() tfprotov5.ProviderServer {
		return PrimeNumberProvider{}
	}

	serveOpts := []tf5server.ServeOpt{
		tf5server.WithManagedDebug(),
		tf5server.WithManagedDebugEnvFilePath(".env.reattach"),
	}
	tf5server.Serve("terraform.io/playground/primes", providerFn, serveOpts...) //nolint:errcheck
}
