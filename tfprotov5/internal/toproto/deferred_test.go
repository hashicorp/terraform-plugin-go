// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func TestDeferred(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.Deferred
		expected *tfplugin5.Deferred
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.Deferred{},
			expected: &tfplugin5.Deferred{
				Reason: tfplugin5.Deferred_UNKNOWN,
			},
		},
		"Reason-ResourceConfigUnknown": {
			in: &tfprotov5.Deferred{
				Reason: tfprotov5.DeferredReasonResourceConfigUnknown,
			},

			expected: &tfplugin5.Deferred{
				Reason: tfplugin5.Deferred_RESOURCE_CONFIG_UNKNOWN,
			},
		},
		"Reason-ProviderConfigUnknown": {
			in: &tfprotov5.Deferred{
				Reason: tfprotov5.DeferredReasonProviderConfigUnknown,
			},

			expected: &tfplugin5.Deferred{
				Reason: tfplugin5.Deferred_PROVIDER_CONFIG_UNKNOWN,
			},
		},
		"Reason-AbsentPrereq": {
			in: &tfprotov5.Deferred{
				Reason: tfprotov5.DeferredReasonAbsentPrereq,
			},

			expected: &tfplugin5.Deferred{
				Reason: tfplugin5.Deferred_ABSENT_PREREQ,
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
				tfplugin5.Deferred{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
