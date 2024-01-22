// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	testTfplugin6Diagnostic = &tfplugin6.Diagnostic{
		Detail:   "test detail",
		Severity: tfplugin6.Diagnostic_ERROR,
		Summary:  "test summary",
	}
	testTfprotov6Diagnostic = &tfprotov6.Diagnostic{
		Detail:   "test detail",
		Severity: tfprotov6.DiagnosticSeverityError,
		Summary:  "test summary",
	}
)

func TestDiagnostic(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.Diagnostic
		expected *tfplugin6.Diagnostic
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.Diagnostic{},
			expected: &tfplugin6.Diagnostic{
				Severity: tfplugin6.Diagnostic_INVALID,
			},
		},
		"Attribute": {
			in: &tfprotov6.Diagnostic{
				Attribute: tftypes.NewAttributePath().WithAttributeName("test"),
			},

			expected: &tfplugin6.Diagnostic{
				Attribute: &tfplugin6.AttributePath{
					Steps: []*tfplugin6.AttributePath_Step{
						{
							Selector: &tfplugin6.AttributePath_Step_AttributeName{
								AttributeName: "test",
							},
						},
					},
				},
				Severity: tfplugin6.Diagnostic_INVALID,
			},
		},
		"Detail": {
			in: &tfprotov6.Diagnostic{
				Detail: "test",
			},
			expected: &tfplugin6.Diagnostic{
				Detail:   "test",
				Severity: tfplugin6.Diagnostic_INVALID,
			},
		},
		"FunctionArgument": {
			in: &tfprotov6.Diagnostic{
				FunctionArgument: pointer(int64(123)),
			},
			expected: &tfplugin6.Diagnostic{
				FunctionArgument: pointer(int64(123)),
				Severity:         tfplugin6.Diagnostic_INVALID,
			},
		},
		"Severity": {
			in: &tfprotov6.Diagnostic{
				Severity: tfprotov6.DiagnosticSeverityError,
			},
			expected: &tfplugin6.Diagnostic{
				Severity: tfplugin6.Diagnostic_ERROR,
			},
		},
		"Summary": {
			in: &tfprotov6.Diagnostic{
				Summary: "test",
			},
			expected: &tfplugin6.Diagnostic{
				Severity: tfplugin6.Diagnostic_INVALID,
				Summary:  "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.Diagnostic(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
				tfplugin6.Diagnostic{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestDiagnostics(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tfprotov6.Diagnostic
		expected []*tfplugin6.Diagnostic
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.Diagnostic{},
		},
		"zero": {
			in:       []*tfprotov6.Diagnostic{},
			expected: []*tfplugin6.Diagnostic{},
		},
		"one": {
			in: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test",
				},
			},
			expected: []*tfplugin6.Diagnostic{
				{
					Severity: tfplugin6.Diagnostic_ERROR,
					Summary:  "test",
				},
			},
		},
		"two": {
			in: []*tfprotov6.Diagnostic{
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test1",
				},
				{
					Severity: tfprotov6.DiagnosticSeverityError,
					Summary:  "test2",
				},
			},
			expected: []*tfplugin6.Diagnostic{
				{
					Severity: tfplugin6.Diagnostic_ERROR,
					Summary:  "test1",
				},
				{
					Severity: tfplugin6.Diagnostic_ERROR,
					Summary:  "test2",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.Diagnostics(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
				tfplugin6.Diagnostic{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestForceValidUTF8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Input string
		Want  string
	}{
		{
			"hello",
			"hello",
		},
		{
			"こんにちは",
			"こんにちは",
		},
		{
			"baﬄe", // NOTE: "ﬄ" is a single-character ligature
			"baﬄe", // ligature is preserved exactly
		},
		{
			"wé́́é́́é́́!", // NOTE: These "e" have multiple combining diacritics
			"wé́́é́́é́́!", // diacritics are preserved exactly
		},
		{
			"😸😾", // Astral-plane characters
			"😸😾", // preserved exactly
		},
		{
			"\xff\xff",     // neither byte is valid UTF-8
			"\ufffd\ufffd", // both are replaced by replacement character
		},
		{
			"\xff\xff\xff\xff\xff",           // more than three invalid bytes
			"\ufffd\ufffd\ufffd\ufffd\ufffd", // still expanded even though it exceeds our initial slice capacity in the implementation
		},
		{
			"t\xffe\xffst",     // invalid bytes interleaved with other content
			"t\ufffde\ufffdst", // the valid content is preserved
		},
		{
			"\xffこんにちは\xffこんにちは",     // invalid bytes interacting with multibyte sequences
			"\ufffdこんにちは\ufffdこんにちは", // the valid content is preserved
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Input, func(t *testing.T) {
			t.Parallel()

			got := toproto.ForceValidUTF8(test.Input)
			if got != test.Want {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, test.Want)
			}
		})
	}
}
