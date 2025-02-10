// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func RawIdentity(in []byte) *tfprotov6.RawIdentity {
	if in == nil {
		return nil
	}

	resp := &tfprotov6.RawIdentity{
		JSON: in,
	}

	return resp
}
