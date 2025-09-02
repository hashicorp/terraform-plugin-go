// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func ActionSchema(in *tfprotov6.ActionSchema) *tfplugin6.ActionSchema {
	if in == nil {
		return nil
	}

	resp := &tfplugin6.ActionSchema{
		Schema: Schema(in.Schema),
	}

	switch in.Type.(type) {
	case tfprotov6.UnlinkedActionSchemaType:
		resp.Type = &tfplugin6.ActionSchema_Unlinked_{
			Unlinked: &tfplugin6.ActionSchema_Unlinked{},
		}
	default:
		// It is not currently possible to create tfprotov6.ActionSchemaType
		// implementations outside the tfprotov6 package. If this panic was reached,
		// it implies that a new event type was introduced and needs to be implemented
		// as a new case above.
		panic(fmt.Sprintf("unimplemented tfprotov6.ActionSchemaType type: %T", in.Type))
	}

	return resp
}
