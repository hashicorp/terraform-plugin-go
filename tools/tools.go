// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

//go:build generate

package tools

import (
	// Protocol Buffers compiler plugin for Go.
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"

	// Protocol Buffers compiler plugin for Go gRPC. This tool is versioned
	// separately from google.golang.org/grpc.
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"

	// copywrite header generation
	_ "github.com/hashicorp/copywrite"
)

//go:generate go run github.com/hashicorp/copywrite headers -d .. --config ../.copywrite.hcl
