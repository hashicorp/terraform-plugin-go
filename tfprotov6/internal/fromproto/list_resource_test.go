// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestValidateListResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateListResourceConfig_Request
		expected *tfprotov6.ValidateListResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateListResourceConfig_Request{},
			expected: &tfprotov6.ValidateListResourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateListResourceConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateListResourceConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ValidateListResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ValidateListResourceConfigRequest{
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
