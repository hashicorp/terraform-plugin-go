// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

func testTfprotov6RawIdentity(t *testing.T, json []byte) *tfprotov6.RawIdentity {
	t.Helper()

	return &tfprotov6.RawIdentity{
		JSON: json,
	}
}
