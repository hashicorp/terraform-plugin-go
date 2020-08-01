package fromproto

import (
	"github.com/hashicorp/terraform-protocol-go/internal/tfplugin5"
	"github.com/hashicorp/terraform-protocol-go/tfprotov5"
)

func StringKind(in tfplugin5.StringKind) tfprotov5.StringKind {
	return tfprotov5.StringKind(in)
}
