//go:build tools
// +build tools

package tools

import (
	// Protocol Buffers compiler plugin for Go gRPC. This tool is versioned
	// separately from google.golang.org/grpc.
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
)
