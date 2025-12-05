// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
)

func TestDeferred(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.Deferred
		expected *tfplugin6.Deferred
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.Deferred{},
			expected: &tfplugin6.Deferred{
				Reason: tfplugin6.Deferred_UNKNOWN,
			},
		},
		"Reason-ResourceConfigUnknown": {
			in: &tfprotov6.Deferred{
				Reason: tfprotov6.DeferredReasonResourceConfigUnknown,
			},

			expected: &tfplugin6.Deferred{
				Reason: tfplugin6.Deferred_RESOURCE_CONFIG_UNKNOWN,
			},
		},
		"Reason-ProviderConfigUnknown": {
			in: &tfprotov6.Deferred{
				Reason: tfprotov6.DeferredReasonProviderConfigUnknown,
			},

			expected: &tfplugin6.Deferred{
				Reason: tfplugin6.Deferred_PROVIDER_CONFIG_UNKNOWN,
			},
		},
		"Reason-AbsentPrereq": {
			in: &tfprotov6.Deferred{
				Reason: tfprotov6.DeferredReasonAbsentPrereq,
			},

			expected: &tfplugin6.Deferred{
				Reason: tfplugin6.Deferred_ABSENT_PREREQ,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Deferred(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Deferred{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
