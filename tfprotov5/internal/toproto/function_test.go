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
	testTfplugin5Error = &tfplugin5.FunctionError{
		Text: "test function error",
	}
	testTfprotov5Error = &tfprotov5.FunctionError{
		Text: "test function error",
	}
)

func TestCallFunction_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.CallFunctionResponse
		expected *tfplugin5.CallFunction_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.CallFunctionResponse{},
			expected: &tfplugin5.CallFunction_Response{},
		},
		"Error": {
			in: &tfprotov5.CallFunctionResponse{
				Error: testTfprotov5Error,
			},
			expected: &tfplugin5.CallFunction_Response{
				Error: testTfplugin5Error,
			},
		},
		"Result": {
			in: &tfprotov5.CallFunctionResponse{
				Result: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.CallFunction_Response{
				Result: testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.CallFunction_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.CallFunction_Response{},
				tfplugin5.FunctionError{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFunction(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.Function
		expected *tfplugin5.Function
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.Function{},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
			},
		},
		"Description": {
			in: &tfprotov5.Function{
				Description: "test",
			},
			expected: &tfplugin5.Function{
				Description: "test",
				Parameters:  []*tfplugin5.Function_Parameter{},
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.Function{
				DescriptionKind: tfprotov5.StringKindMarkdown,
			},
			expected: &tfplugin5.Function{
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
				Parameters:      []*tfplugin5.Function_Parameter{},
			},
		},
		"DeprecationMessage": {
			in: &tfprotov5.Function{
				DeprecationMessage: "test",
			},
			expected: &tfplugin5.Function{
				DeprecationMessage: "test",
				Parameters:         []*tfplugin5.Function_Parameter{},
			},
		},
		"Parameters": {
			in: &tfprotov5.Function{
				Parameters: []*tfprotov5.FunctionParameter{
					{
						Type: tftypes.Bool,
					},
				},
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{
					{
						Type: []byte(`"bool"`),
					},
				},
			},
		},
		"Return": {
			in: &tfprotov5.Function{
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"Summary": {
			in: &tfprotov5.Function{
				Summary: "test",
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
				Summary:    "test",
			},
		},
		"VariadicParameter": {
			in: &tfprotov5.Function{
				VariadicParameter: &tfprotov5.FunctionParameter{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
				VariadicParameter: &tfplugin5.Function_Parameter{
					Type: []byte(`"bool"`),
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Function(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Function{},
				tfplugin5.Function_Parameter{},
				tfplugin5.Function_Return{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFunction_Parameter(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.FunctionParameter
		expected *tfplugin5.Function_Parameter
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.FunctionParameter{},
			expected: &tfplugin5.Function_Parameter{},
		},
		"AllowNullValue": {
			in: &tfprotov5.FunctionParameter{
				AllowNullValue: true,
			},
			expected: &tfplugin5.Function_Parameter{
				AllowNullValue: true,
			},
		},
		"AllowUnknownValues": {
			in: &tfprotov5.FunctionParameter{
				AllowUnknownValues: true,
			},
			expected: &tfplugin5.Function_Parameter{
				AllowUnknownValues: true,
			},
		},
		"Description": {
			in: &tfprotov5.FunctionParameter{
				Description: "test",
			},
			expected: &tfplugin5.Function_Parameter{
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.FunctionParameter{
				DescriptionKind: tfprotov5.StringKindMarkdown,
			},
			expected: &tfplugin5.Function_Parameter{
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
			},
		},
		"Name": {
			in: &tfprotov5.FunctionParameter{
				Name: "test",
			},
			expected: &tfplugin5.Function_Parameter{
				Name: "test",
			},
		},
		"Type": {
			in: &tfprotov5.FunctionParameter{
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				Type: []byte(`"bool"`),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Function_Parameter(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Function_Parameter{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFunction_Return(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.FunctionReturn
		expected *tfplugin5.Function_Return
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.FunctionReturn{},
			expected: &tfplugin5.Function_Return{},
		},
		"Type": {
			in: &tfprotov5.FunctionReturn{
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Return{
				Type: []byte(`"bool"`),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.Function_Return(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Function_Return{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetFunctions_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.GetFunctionsResponse
		expected *tfplugin5.GetFunctions_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.GetFunctionsResponse{},
			expected: &tfplugin5.GetFunctions_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Functions:   map[string]*tfplugin5.Function{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.GetFunctionsResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.GetFunctions_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				Functions: map[string]*tfplugin5.Function{},
			},
		},
		"Functions": {
			in: &tfprotov5.GetFunctionsResponse{
				Functions: map[string]*tfprotov5.Function{
					"test": {
						Return: &tfprotov5.FunctionReturn{
							Type: tftypes.Bool,
						},
					},
				},
			},
			expected: &tfplugin5.GetFunctions_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Functions: map[string]*tfplugin5.Function{
					"test": {
						Parameters: []*tfplugin5.Function_Parameter{},
						Return: &tfplugin5.Function_Return{
							Type: []byte(`"bool"`),
						},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetFunctions_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.Function{},
				tfplugin5.Function_Parameter{},
				tfplugin5.Function_Return{},
				tfplugin5.GetFunctions_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetMetadata_FunctionMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.FunctionMetadata
		expected *tfplugin5.GetMetadata_FunctionMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.FunctionMetadata{},
			expected: &tfplugin5.GetMetadata_FunctionMetadata{},
		},
		"Name": {
			in: &tfprotov5.FunctionMetadata{
				Name: "test",
			},
			expected: &tfplugin5.GetMetadata_FunctionMetadata{
				Name: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_FunctionMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.GetMetadata_FunctionMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
