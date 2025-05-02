package toproto

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

type InvokeActionCallBackServer struct {
	ProtoServer grpc.ServerStreamingServer[tfplugin6.InvokeAction_Event]
}

func (i InvokeActionCallBackServer) Send(ctx context.Context, event tfprotov6.InvokeActionEvent) error {
	switch actionEvent := event.(type) {
	case *tfprotov6.StartedActionEvent:
		logging.ProtocolTrace(ctx, "Sending StartedActionEvent")
		tfplugin6Event := &tfplugin6.InvokeAction_Event{
			Event: InvokeAction_Event_Started_(actionEvent),
		}

		return i.ProtoServer.Send(tfplugin6Event)
	case *tfprotov6.ProgressActionEvent:
		logging.ProtocolTrace(ctx, "Sending ProgressActionEvent")
		tfplugin6Event := &tfplugin6.InvokeAction_Event{
			Event: InvokeAction_Event_Progress_(actionEvent),
		}

		return i.ProtoServer.Send(tfplugin6Event)
	case *tfprotov6.FinishedActionEvent:
		logging.ProtocolTrace(ctx, "Sending FinishedActionEvent")
		tfplugin6Event := &tfplugin6.InvokeAction_Event{
			Event: InvokeAction_Event_Finished_(actionEvent),
		}

		return i.ProtoServer.Send(tfplugin6Event)
	case *tfprotov6.CancelledActionEvent:
		logging.ProtocolTrace(ctx, "Sending CancelledActionEvent")
		tfplugin6Event := &tfplugin6.InvokeAction_Event{
			Event: InvokeAction_Event_Cancelled_(actionEvent),
		}

		return i.ProtoServer.Send(tfplugin6Event)
	default:
		return fmt.Errorf("unknown InvokeActionEvent type: %T", actionEvent)
	}
}
