// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package tf5serverlogging_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"

	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tf5serverlogging"
)

func TestValidateResourceTypeConfigClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ValidateResourceTypeConfigClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ValidateResourceTypeConfigClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced client capabilities",
					"@module":  "sdk.proto",
					"tf_client_capability_write_only_attributes_allowed": false,
				},
			},
		},
		"write_only_attributes_allowed": {
			capabilities: &tfprotov5.ValidateResourceTypeConfigClientCapabilities{
				WriteOnlyAttributesAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Announced client capabilities",
					"@module":  "sdk.proto",
					"tf_client_capability_write_only_attributes_allowed": true,
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

			tf5serverlogging.ValidateResourceTypeConfigClientCapabilities(ctx, testCase.capabilities)

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

func TestConfigureProviderClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ConfigureProviderClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ConfigureProviderClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.ConfigureProviderClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.ConfigureProviderClientCapabilities(ctx, testCase.capabilities)

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

func TestReadDataSourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ReadDataSourceClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ReadDataSourceClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.ReadDataSourceClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.ReadDataSourceClientCapabilities(ctx, testCase.capabilities)

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

func TestReadResourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ReadResourceClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ReadResourceClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.ReadResourceClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.ReadResourceClientCapabilities(ctx, testCase.capabilities)

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

func TestPlanResourceChangeClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.PlanResourceChangeClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.PlanResourceChangeClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.PlanResourceChangeClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.PlanResourceChangeClientCapabilities(ctx, testCase.capabilities)

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

func TestImportResourceStateClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.ImportResourceStateClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.ImportResourceStateClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.ImportResourceStateClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.ImportResourceStateClientCapabilities(ctx, testCase.capabilities)

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

func TestOpenEphemeralResourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.OpenEphemeralResourceClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.OpenEphemeralResourceClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.OpenEphemeralResourceClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.OpenEphemeralResourceClientCapabilities(ctx, testCase.capabilities)

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

func TestPlanActionClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		capabilities *tfprotov5.PlanActionClientCapabilities
		expected     []map[string]interface{}
	}{
		"nil": {
			capabilities: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "No announced client capabilities",
					"@module":  "sdk.proto",
				},
			},
		},
		"empty": {
			capabilities: &tfprotov5.PlanActionClientCapabilities{},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": false,
				},
			},
		},
		"deferral_allowed": {
			capabilities: &tfprotov5.PlanActionClientCapabilities{
				DeferralAllowed: true,
			},
			expected: []map[string]interface{}{
				{
					"@level":                                "trace",
					"@message":                              "Announced client capabilities",
					"@module":                               "sdk.proto",
					"tf_client_capability_deferral_allowed": true,
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

			tf5serverlogging.PlanActionClientCapabilities(ctx, testCase.capabilities)

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
