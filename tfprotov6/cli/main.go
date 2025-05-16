// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <provider-binary>")
		os.Exit(1)
	}
	providerBinary := os.Args[1]
	cmd := exec.Command(providerBinary)

	pluginMap := map[int]plugin.PluginSet{
		6: {
			"provider": &GRPCPluginClient{
				ClientFunc: func(c *grpc.ClientConn) any {
					return &ProviderClient{
						client: tfplugin6.NewProviderClient(c),
					}
				},
			},
		},
	}

	magicCookie, ok := os.LookupEnv("TF_PLUGIN_MAGIC_COOKIE")
	if !ok {
		fmt.Println("TF_PLUGIN_MAGIC_COOKIE environment variable not set")
		os.Exit(1)
	}

	config := &plugin.ClientConfig{
		HandshakeConfig: plugin.HandshakeConfig{
			ProtocolVersion:  6,
			MagicCookieKey:   "TF_PLUGIN_MAGIC_COOKIE",
			MagicCookieValue: magicCookie,
		},
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolGRPC},
		VersionedPlugins: pluginMap,
		Cmd:              cmd,
		Managed:          true,
	}

	fail := func(msg string, args ...interface{}) {
		fmt.Printf(msg+"\n", args...)
		os.Exit(1)
	}

	pluginClient := plugin.NewClient(config)
	_, err := pluginClient.Start()
	if err != nil {
		fail("failed to start plugin client: %v", err)
	}

	defer pluginClient.Kill()

	protocolClient, err := pluginClient.Client()
	if err != nil {
		fail("failed to get protocol client: %v", err)
	}

	rawProviderClient, err := protocolClient.Dispense("provider")
	if err != nil {
		fail("failed to dispense provider: %v", err)
	}

	providerClient, ok := rawProviderClient.(*ProviderClient)
	if !ok {
		fail("failed to cast to provider client")
	}

	ctx := context.Background()
	response, err := providerClient.GetProviderSchema(ctx)
	if err != nil {
		fmt.Printf("failed to get provider schema: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("=== provider schema ===")
	fmt.Printf("%v\n", response.Provider)
}
