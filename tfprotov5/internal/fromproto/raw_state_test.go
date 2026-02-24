// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func testTfplugin5RawState(t *testing.T, json []byte) *tfplugin5.RawState {
	t.Helper()

	return &tfplugin5.RawState{
		// Flatmap is intentionally not supported, nor necessary.
		Json: json,
	}
}

func testTfprotov5RawState(t *testing.T, json []byte) *tfprotov5.RawState {
	t.Helper()

	return &tfprotov5.RawState{
		// Flatmap is intentionally not supported, nor necessary.
		JSON: json,
	}
}
