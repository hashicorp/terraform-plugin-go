package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func LinkedResources(in map[string]*tfprotov6.LinkedResource) map[string]*tfplugin6.ActionSchema_LinkedResource {
	resp := make(map[string]*tfplugin6.ActionSchema_LinkedResource)

	for k, v := range in {
		resp[k] = &tfplugin6.ActionSchema_LinkedResource{
			Type: v.TypeName,
		}
	}

	return resp
}
