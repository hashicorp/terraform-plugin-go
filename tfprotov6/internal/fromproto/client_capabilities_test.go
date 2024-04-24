// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.ClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
