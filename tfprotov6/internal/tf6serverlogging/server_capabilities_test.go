// Copyright (c) HashiCorp, Inc.
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

func TestServerCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov6.ServerCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 false,
				},
			},
		},
		"empty": {
			capabilities: &tfprotov6.ServerCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 false,
				},
			},
		},
		"get_provider_schema_optional": {
			capabilities: &tfprotov6.ServerCapabilities{
				GetProviderSchemaOptional: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": true,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 false,
				},
			},
		},
		"move_resource_state": {
			capabilities: &tfprotov6.ServerCapabilities{
				MoveResourceState: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          true,
					"tf_server_capability_plan_destroy":                 false,
				},
			},
		},
		"plan_destroy": {
			capabilities: &tfprotov6.ServerCapabilities{
				PlanDestroy: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 true,
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

			tf6serverlogging.ServerCapabilities(ctx, testCase.capabilities)

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

func TestStateStoreServerCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov6.StateStoreServerCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced server capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov6.StateStoreServerCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                          "trace",
					"@message":                        "Announced server capabilities",
					"@module":                         "sdk.proto",
					"tf_server_capability_chunk_size": "0MB", // This is invalid provider behavior, so this log should never occur in practice.
				},
			},
		},
		"chunk_size": {
			capabilities: &tfprotov6.StateStoreServerCapabilities{
				ChunkSize: 8 << 20,
			},
			expected: []map[string]interface{}{
				{
					"@level":                          "trace",
					"@message":                        "Announced server capabilities",
					"@module":                         "sdk.proto",
					"tf_server_capability_chunk_size": "8MB",
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

			tf6serverlogging.StateStoreServerCapabilities(ctx, testCase.capabilities)

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
