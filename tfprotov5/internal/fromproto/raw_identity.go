// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto

import (
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func RawIdentity(in []byte) *tfprotov5.RawIdentity {
	if in == nil {
		return nil
	}

	resp := &tfprotov5.RawIdentity{
		JSON: in,
	}

	return resp
}
