package diag_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/diag"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"
)

func TestDiagnosticsErrorCount(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		diagnostics diag.Diagnostics
		expected    int
	}{
		"nil": {
			diagnostics: nil,
			expected:    0,
		},
		"empty": {
			diagnostics: diag.Diagnostics{},
			expected:    0,
		},
		"error": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
			},
			expected: 1,
		},
		"errors": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
			},
			expected: 2,
		},
		"invalid": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
			},
			expected: 0,
		},
		"invalids": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
			},
			expected: 0,
		},
		"warning": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
			},
			expected: 0,
		},
		"warnings": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 2",
				},
			},
			expected: 0,
		},
		"mixed": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 2",
					Detail:   "test warning detail 2",
				},
			},
			expected: 2,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.diagnostics.ErrorCount()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestDiagnosticsLog(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		diagnostics diag.Diagnostics
		expected    []map[string]interface{}
	}{
		"nil": {
			diagnostics: nil,
			expected:    nil,
		},
		"empty": {
			diagnostics: diag.Diagnostics{},
			expected:    nil,
		},
		"attribute": {
			diagnostics: diag.Diagnostics{
				{
					Severity:  tfprotov6.DiagnosticSeverityError,
					Summary:   "test error summary 1",
					Detail:    "test error detail 1",
					Attribute: tftypes.NewAttributePath().WithAttributeName("test"),
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":               "error",
					"@message":             "Response contains error diagnostic",
					"@module":              "sdk.proto",
					"diagnostic_attribute": "AttributeName(\"test\")",
					"diagnostic_detail":    "test error detail 1",
					"diagnostic_severity":  "ERROR",
					"diagnostic_summary":   "test error summary 1",
				},
			},
		},
		"error": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":              "error",
					"@message":            "Response contains error diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test error detail 1",
					"diagnostic_severity": "ERROR",
					"diagnostic_summary":  "test error summary 1",
				},
			},
		},
		"errors": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":              "error",
					"@message":            "Response contains error diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test error detail 1",
					"diagnostic_severity": "ERROR",
					"diagnostic_summary":  "test error summary 1",
				},
				{
					"@level":              "error",
					"@message":            "Response contains error diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test error detail 2",
					"diagnostic_severity": "ERROR",
					"diagnostic_summary":  "test error summary 2",
				},
			},
		},
		"invalid": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":              "warn",
					"@message":            "Response contains unknown diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test invalid detail 1",
					"diagnostic_severity": "INVALID",
					"diagnostic_summary":  "test invalid summary 1",
				},
			},
		},
		"invalids": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
			},
			expected: []map[string]interface{}{
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
					"@message":            "Response contains unknown diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test invalid detail 2",
					"diagnostic_severity": "INVALID",
					"diagnostic_summary":  "test invalid summary 2",
				},
			},
		},
		"warning": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":              "warn",
					"@message":            "Response contains warning diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test warning detail 1",
					"diagnostic_severity": "WARNING",
					"diagnostic_summary":  "test warning summary 1",
				},
			},
		},
		"warnings": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 2",
					Detail:   "test warning detail 2",
				},
			},
			expected: []map[string]interface{}{
				{
					"@level":              "warn",
					"@message":            "Response contains warning diagnostic",
					"@module":             "sdk.proto",
					"diagnostic_detail":   "test warning detail 1",
					"diagnostic_severity": "WARNING",
					"diagnostic_summary":  "test warning summary 1",
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
		"mixed": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 2",
					Detail:   "test warning detail 2",
				},
			},
			expected: []map[string]interface{}{
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

			testCase.diagnostics.Log(ctx)

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

func TestDiagnosticsWarningCount(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		diagnostics diag.Diagnostics
		expected    int
	}{
		"nil": {
			diagnostics: nil,
			expected:    0,
		},
		"empty": {
			diagnostics: diag.Diagnostics{},
			expected:    0,
		},
		"error": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
			},
			expected: 0,
		},
		"errors": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
			},
			expected: 0,
		},
		"invalid": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
			},
			expected: 0,
		},
		"invalids": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
			},
			expected: 0,
		},
		"warning": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
			},
			expected: 1,
		},
		"warnings": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 2",
				},
			},
			expected: 2,
		},
		"mixed": {
			diagnostics: diag.Diagnostics{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 1",
					Detail:   "test error detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 1",
					Detail:   "test invalid detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 1",
					Detail:   "test warning detail 1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test error summary 2",
					Detail:   "test error detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityInvalid,
					Summary:  "test invalid summary 2",
					Detail:   "test invalid detail 2",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityWarning,
					Summary:  "test warning summary 2",
					Detail:   "test warning detail 2",
				},
			},
			expected: 2,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.diagnostics.WarningCount()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
