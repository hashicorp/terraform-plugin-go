// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tf5serverlogging_test

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"

	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/diag"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tf5serverlogging"
)

func TestDownstreamRequest(t *testing.T) {
	t.Parallel()

	var output bytes.Buffer

	ctx := tfsdklogtest.RootLogger(context.Background(), &output)
	ctx = logging.ProtoSubsystemContext(ctx, tfsdklog.Options{})

	got := tf5serverlogging.DownstreamRequest(ctx)

	if _, ok := got.Value(tf5serverlogging.ContextKeyDownstreamRequestStartTime{}).(time.Time); !ok {
		t.Error("missing downstream request start time context key")
	}

	entries, err := tfsdklogtest.MultilineJSONDecode(&output)

	if err != nil {
		t.Fatalf("unable to read multiple line JSON: %s", err)
	}

	expectedEntries := []map[string]interface{}{
		{
			"@level":   "trace",
			"@message": "Sending request downstream",
			"@module":  "sdk.proto",
		},
	}

	if diff := cmp.Diff(entries, expectedEntries); diff != "" {
		t.Errorf("unexpected difference: %s", diff)
	}
}

func TestDownstreamResponse(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		diagnostics diag.Diagnostics
		expected    []map[string]interface{}
	}{
		"diagnostics-nil": {
			diagnostics: nil,
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Received downstream response",
					"@module":  "sdk.proto",
					// go-hclog treats int as float64
					"diagnostic_error_count":   float64(0),
					"diagnostic_warning_count": float64(0),
				},
			},
		},
		"diagnostics-empty": {
			diagnostics: diag.Diagnostics{},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Received downstream response",
					"@module":  "sdk.proto",
					// go-hclog treats int as float64
					"diagnostic_error_count":   float64(0),
					"diagnostic_warning_count": float64(0),
				},
			},
		},
		"diagnostics": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityWarning,
					Summary:  "test warning summary 2",
					Detail:   "test warning detail 2",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":   "trace",
					"@message": "Received downstream response",
					"@module":  "sdk.proto",
					// go-hclog treats int as float64
					"diagnostic_error_count":   float64(2),
					"diagnostic_warning_count": float64(2),
				},
				{
					"@level":              "error",
					"@message":            "Response contains error diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test error detail 1",
					"diagnostic_severity": "ERROR",
					"diagnostic_summary":  "test error summary 1",
				},
				{
					"@level":              "warn",
					"@message":            "Response contains unknown diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test invalid detail 1",
					"diagnostic_severity": "INVALID",
					"diagnostic_summary":  "test invalid summary 1",
				},
				{
					"@level":              "warn",
					"@message":            "Response contains warning diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test warning detail 1",
					"diagnostic_severity": "WARNING",
					"diagnostic_summary":  "test warning summary 1",
				},
				{
					"@level":              "error",
					"@message":            "Response contains error diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test error detail 2",
					"diagnostic_severity": "ERROR",
					"diagnostic_summary":  "test error summary 2",
				},
				{
					"@level":              "warn",
					"@message":            "Response contains unknown diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test invalid detail 2",
					"diagnostic_severity": "INVALID",
					"diagnostic_summary":  "test invalid summary 2",
				},
				{
					"@level":              "warn",
					"@message":            "Response contains warning diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test warning detail 2",
					"diagnostic_severity": "WARNING",
					"diagnostic_summary":  "test warning summary 2",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer

			ctx := tfsdklogtest.RootLogger(context.Background(), &output)
			ctx = logging.ProtoSubsystemContext(ctx, tfsdklog.Options{})

			tf5serverlogging.DownstreamResponse(ctx, testCase.diagnostics)

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

func TestDownstreamResponseWithError(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		functionError *tfprotov5.FunctionError
		expected      []map[string]interface{}
	}{
		"function-error-nil": {
			functionError: nil,
			expected: []map[string]interface{}{
				{
					"@level":                "trace",
					"@message":              "Received downstream response",
					"@module":               "sdk.proto",
					"function_error_exists": false,
				},
			},
		},
		"function-error-empty": {
			functionError: &tfprotov5.FunctionError{},
			expected: []map[string]interface{}{
				{
					"@level":                "trace",
					"@message":              "Received downstream response",
					"@module":               "sdk.proto",
					"function_error_exists": false,
				},
			},
		},
		"function-error": {
			functionError: &tfprotov5.FunctionError{
				Text: "test function error",
			},
			expected: []map[string]interface{}{
				{
					"@level":                "trace",
					"@message":              "Received downstream response",
					"@module":               "sdk.proto",
					"function_error_exists": true,
				},
				{
					"@level":             "error",
					"@message":           "Response contains function error",
					"@module":            "sdk.proto",
					"function_error_msg": "test function error",
				},
			},
		},
		"function-error-with-argument-only": {
			functionError: &tfprotov5.FunctionError{
				FunctionArgument: pointer(int64(0)),
			},
			expected: []map[string]interface{}{
				{
					"@level":                "trace",
					"@message":              "Received downstream response",
					"@module":               "sdk.proto",
					"function_error_exists": true,
				},
				{
					"@level":                  "error",
					"@message":                "Response contains function error",
					"@module":                 "sdk.proto",
					"function_error_argument": float64(0),
				},
			},
		},
		"function-error-with-argument-and-message": {
			functionError: &tfprotov5.FunctionError{
				Text:             "test function error",
				FunctionArgument: pointer(int64(0)),
			},
			expected: []map[string]interface{}{
				{
					"@level":                "trace",
					"@message":              "Received downstream response",
					"@module":               "sdk.proto",
					"function_error_exists": true,
				},
				{
					"@level":                  "error",
					"@message":                "Response contains function error",
					"@module":                 "sdk.proto",
					"function_error_msg":      "test function error",
					"function_error_argument": float64(0),
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer

			ctx := tfsdklogtest.RootLogger(context.Background(), &output)
			ctx = logging.ProtoSubsystemContext(ctx, tfsdklog.Options{})

			tf5serverlogging.DownstreamResponseWithError(ctx, testCase.functionError)

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
