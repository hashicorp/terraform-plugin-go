// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tf6serverlogging_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tf6serverlogging"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"
)

func TestDeferred(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		deferred *tfprotov6.Deferred
		expected []map[string]interface{}
	}{
		"nil": {
			deferred: nil,
			expected: nil,
		},
		"empty": {
			deferred: &tfprotov6.Deferred{},
			expected: []map[string]interface{}{
				{
					"@level":             "trace",
					"@message":           "Received downstream deferred response",
					"@module":            "sdk.proto",
					"tf_deferred_reason": "UNKNOWN",
				},
			},
		},
		"deferred": {
			deferred: &tfprotov6.Deferred{
				Reason: tfprotov6.DeferredReasonProviderConfigUnknown,
			},
			expected: []map[string]interface{}{
				{
					"@level":             "trace",
					"@message":           "Received downstream deferred response",
					"@module":            "sdk.proto",
					"tf_deferred_reason": "PROVIDER_CONFIG_UNKNOWN",
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer

			ctx := tfsdklogtest.RootLogger(context.Background(), &output)
			ctx = logging.ProtoSubsystemContext(ctx, tfsdklog.Options{})

			tf6serverlogging.Deferred(ctx, testCase.deferred)

			entries, err := tfsdklogtest.MultilineJSONDecode(&output)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(entries, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
