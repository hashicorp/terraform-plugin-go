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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var (
	testTfplugin5Diagnostic = &tfplugin5.Diagnostic{
		Detail:   "test detail",
		Severity: tfplugin5.Diagnostic_ERROR,
		Summary:  "test summary",
	}
	testTfprotov5Diagnostic = &tfprotov5.Diagnostic{
		Detail:   "test detail",
		Severity: tfprotov5.DiagnosticSeverityError,
		Summary:  "test summary",
	}
)

func TestDiagnostic(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.Diagnostic
		expected *tfplugin5.Diagnostic
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.Diagnostic{},
			expected: &tfplugin5.Diagnostic{
				Severity: tfplugin5.Diagnostic_INVALID,
			},
		},
		"Attribute": {
			in: &tfprotov5.Diagnostic{
				Attribute: tftypes.NewAttributePath().WithAttributeName("test"),
			},

			expected: &tfplugin5.Diagnostic{
				Attribute: &tfplugin5.AttributePath{
					Steps: []*tfplugin5.AttributePath_Step{
						{
							Selector: &tfplugin5.AttributePath_Step_AttributeName{
								AttributeName: "test",
							},
						},
					},
				},
				Severity: tfplugin5.Diagnostic_INVALID,
			},
		},
		"Detail": {
			in: &tfprotov5.Diagnostic{
				Detail: "test",
			},
			expected: &tfplugin5.Diagnostic{
				Detail:   "test",
				Severity: tfplugin5.Diagnostic_INVALID,
			},
		},
		"Severity": {
			in: &tfprotov5.Diagnostic{
				Severity: tfprotov5.DiagnosticSeverityError,
			},
			expected: &tfplugin5.Diagnostic{
				Severity: tfplugin5.Diagnostic_ERROR,
			},
		},
		"Summary": {
			in: &tfprotov5.Diagnostic{
				Summary: "test",
			},
			expected: &tfplugin5.Diagnostic{
				Severity: tfplugin5.Diagnostic_INVALID,
				Summary:  "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Diagnostic(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.AttributePath{},
				tfplugin5.AttributePath_Step{},
				tfplugin5.Diagnostic{},
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
		in       []*tfprotov5.Diagnostic
		expected []*tfplugin5.Diagnostic
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin5.Diagnostic{},
		},
		"zero": {
			in:       []*tfprotov5.Diagnostic{},
			expected: []*tfplugin5.Diagnostic{},
		},
		"one": {
			in: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "test",
				},
			},
			expected: []*tfplugin5.Diagnostic{
				{
					Severity: tfplugin5.Diagnostic_ERROR,
					Summary:  "test",
				},
			},
		},
		"two": {
			in: []*tfprotov5.Diagnostic{
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "test1",
				},
				{
					Severity: tfprotov5.DiagnosticSeverityError,
					Summary:  "test2",
				},
			},
			expected: []*tfplugin5.Diagnostic{
				{
					Severity: tfplugin5.Diagnostic_ERROR,
					Summary:  "test1",
				},
				{
					Severity: tfplugin5.Diagnostic_ERROR,
					Summary:  "test2",
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Diagnostics(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.AttributePath{},
				tfplugin5.AttributePath_Step{},
				tfplugin5.Diagnostic{},
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
