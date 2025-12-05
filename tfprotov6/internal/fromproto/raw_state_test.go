// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func testTfplugin6RawState(t *testing.T, json []byte) *tfplugin6.RawState {
	t.Helper()

	return &tfplugin6.RawState{
		// Flatmap is intentionally not supported, nor necessary.
		Json: json,
	}
}

func testTfprotov6RawState(t *testing.T, json []byte) *tfprotov6.RawState {
	t.Helper()

	return &tfprotov6.RawState{
		// Flatmap is intentionally not supported, nor necessary.
		JSON: json,
	}
}
