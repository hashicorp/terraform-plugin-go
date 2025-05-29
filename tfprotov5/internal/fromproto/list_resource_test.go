// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestValidateListResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ValidateListResourceConfig_Request
		expected *tfprotov5.ValidateListResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ValidateListResourceConfig_Request{},
			expected: &tfprotov5.ValidateListResourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin5.ValidateListResourceConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateListResourceConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ValidateListResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ValidateListResourceConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateListResourceConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
