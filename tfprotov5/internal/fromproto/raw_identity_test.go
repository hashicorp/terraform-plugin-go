// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

func testTfprotov5RawIdentity(t *testing.T, json []byte) *tfprotov5.RawIdentity {
	t.Helper()

	return &tfprotov5.RawIdentity{
		JSON: json,
	}
}
