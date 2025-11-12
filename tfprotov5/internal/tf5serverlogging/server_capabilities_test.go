// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5serverlogging_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tf5serverlogging"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"
)

func TestServerCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ServerCapabilities
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
					"tf_server_capability_generate_resource_config":     false,
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ServerCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 false,
					"tf_server_capability_generate_resource_config":     false,
				},
			},
		},
		"get_provider_schema_optional": {
			capabilities: &tfprotov5.ServerCapabilities{
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
					"tf_server_capability_generate_resource_config":     false,
				},
			},
		},
		"move_resource_state": {
			capabilities: &tfprotov5.ServerCapabilities{
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
					"tf_server_capability_generate_resource_config":     false,
				},
			},
		},
		"plan_destroy": {
			capabilities: &tfprotov5.ServerCapabilities{
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
					"tf_server_capability_generate_resource_config":     false,
				},
			},
		},
		"generate_resource_config": {
			capabilities: &tfprotov5.ServerCapabilities{
				GenerateResourceConfig: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced server capabilities",
					"@module":  "sdk.proto",
					"tf_server_capability_get_provider_schema_optional": false,
					"tf_server_capability_move_resource_state":          false,
					"tf_server_capability_plan_destroy":                 false,
					"tf_server_capability_generate_resource_config":     true,
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

			tf5serverlogging.ServerCapabilities(ctx, testCase.capabilities)

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
