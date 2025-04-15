package toproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func LinkedResources(in []*tfprotov6.LinkedResource) []*tfplugin6.ActionSchema_LinkedResource {
	resp := make([]*tfplugin6.ActionSchema_LinkedResource, len(in))

	for i, v := range in {
		resp[i] = &tfplugin6.ActionSchema_LinkedResource{
			Type:      v.TypeName,
			Attribute: AttributePath(v.Attribute),
		}
	}

	return resp
}
