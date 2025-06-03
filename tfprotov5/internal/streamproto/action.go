// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamproto

import (
	"context"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"google.golang.org/grpc"
)

func InvokeActionStreamingServer(in grpc.ServerStreamingServer[tfplugin5.InvokeAction_Event]) tfprotov5.InvokeActionStreamingServer {
	return InvokeActionStreamingServerImpl{in}
}

var _ tfprotov5.InvokeActionStreamingServer = InvokeActionStreamingServerImpl{}

type InvokeActionStreamingServerImpl struct {
	protoServer grpc.ServerStreamingServer[tfplugin5.InvokeAction_Event]
}

func (s InvokeActionStreamingServerImpl) Send(event *tfprotov5.InvokeActionEvent) error {
	protoEvent := toproto.InvokeActionEvent(event)
	return s.protoServer.Send(protoEvent)
}

func (s InvokeActionStreamingServerImpl) Context() context.Context {
	return s.protoServer.Context()
}
