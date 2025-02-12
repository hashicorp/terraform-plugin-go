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
	testTfplugin6Error = &tfplugin6.FunctionError{
		Text: "test function error",
	}
	testTfprotov6Error = &tfprotov6.FunctionError{
		Text: "test function error",
	}
)

func TestCallFunction_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.CallFunctionResponse
		expected *tfplugin6.CallFunction_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.CallFunctionResponse{},
			expected: &tfplugin6.CallFunction_Response{},
		},
		"Error": {
			in: &tfprotov6.CallFunctionResponse{
				Error: testTfprotov6Error,
			},
			expected: &tfplugin6.CallFunction_Response{
				Error: testTfplugin6Error,
			},
		},
		"Result": {
			in: &tfprotov6.CallFunctionResponse{
				Result: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.CallFunction_Response{
				Result: testTfplugin6DynamicValue(),
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
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.CallFunction_Response{},
				tfplugin6.FunctionError{},
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
		in       *tfprotov6.Function
		expected *tfplugin6.Function
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.Function{},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
			},
		},
		"Description": {
			in: &tfprotov6.Function{
				Description: "test",
			},
			expected: &tfplugin6.Function{
				Description: "test",
				Parameters:  []*tfplugin6.Function_Parameter{},
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.Function{
				DescriptionKind: tfprotov6.StringKindMarkdown,
			},
			expected: &tfplugin6.Function{
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
				Parameters:      []*tfplugin6.Function_Parameter{},
			},
		},
		"DeprecationMessage": {
			in: &tfprotov6.Function{
				DeprecationMessage: "test",
			},
			expected: &tfplugin6.Function{
				DeprecationMessage: "test",
				Parameters:         []*tfplugin6.Function_Parameter{},
			},
		},
		"Parameters": {
			in: &tfprotov6.Function{
				Parameters: []*tfprotov6.FunctionParameter{
					{
						Type: tftypes.Bool,
					},
				},
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{
					{
						Type: []byte(`"bool"`),
					},
				},
			},
		},
		"Return": {
			in: &tfprotov6.Function{
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"Summary": {
			in: &tfprotov6.Function{
				Summary: "test",
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
				Summary:    "test",
			},
		},
		"VariadicParameter": {
			in: &tfprotov6.Function{
				VariadicParameter: &tfprotov6.FunctionParameter{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
				VariadicParameter: &tfplugin6.Function_Parameter{
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
				tfplugin6.Function{},
				tfplugin6.Function_Parameter{},
				tfplugin6.Function_Return{},
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
		in       *tfprotov6.FunctionParameter
		expected *tfplugin6.Function_Parameter
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.FunctionParameter{},
			expected: &tfplugin6.Function_Parameter{},
		},
		"AllowNullValue": {
			in: &tfprotov6.FunctionParameter{
				AllowNullValue: true,
			},
			expected: &tfplugin6.Function_Parameter{
				AllowNullValue: true,
			},
		},
		"AllowUnknownValues": {
			in: &tfprotov6.FunctionParameter{
				AllowUnknownValues: true,
			},
			expected: &tfplugin6.Function_Parameter{
				AllowUnknownValues: true,
			},
		},
		"Description": {
			in: &tfprotov6.FunctionParameter{
				Description: "test",
			},
			expected: &tfplugin6.Function_Parameter{
				Description: "test",
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.FunctionParameter{
				DescriptionKind: tfprotov6.StringKindMarkdown,
			},
			expected: &tfplugin6.Function_Parameter{
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
			},
		},
		"Name": {
			in: &tfprotov6.FunctionParameter{
				Name: "test",
			},
			expected: &tfplugin6.Function_Parameter{
				Name: "test",
			},
		},
		"Type": {
			in: &tfprotov6.FunctionParameter{
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
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
				tfplugin6.Function_Parameter{},
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
		in       *tfprotov6.FunctionReturn
		expected *tfplugin6.Function_Return
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.FunctionReturn{},
			expected: &tfplugin6.Function_Return{},
		},
		"Type": {
			in: &tfprotov6.FunctionReturn{
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Return{
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
				tfplugin6.Function_Return{},
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
		in       *tfprotov6.GetFunctionsResponse
		expected *tfplugin6.GetFunctions_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.GetFunctionsResponse{},
			expected: &tfplugin6.GetFunctions_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Functions:   map[string]*tfplugin6.Function{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.GetFunctionsResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.GetFunctions_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				Functions: map[string]*tfplugin6.Function{},
			},
		},
		"Functions": {
			in: &tfprotov6.GetFunctionsResponse{
				Functions: map[string]*tfprotov6.Function{
					"test": {
						Return: &tfprotov6.FunctionReturn{
							Type: tftypes.Bool,
						},
					},
				},
			},
			expected: &tfplugin6.GetFunctions_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Functions: map[string]*tfplugin6.Function{
					"test": {
						Parameters: []*tfplugin6.Function_Parameter{},
						Return: &tfplugin6.Function_Return{
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
				tfplugin6.Diagnostic{},
				tfplugin6.Function{},
				tfplugin6.Function_Parameter{},
				tfplugin6.Function_Return{},
				tfplugin6.GetFunctions_Response{},
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
		in       *tfprotov6.FunctionMetadata
		expected *tfplugin6.GetMetadata_FunctionMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.FunctionMetadata{},
			expected: &tfplugin6.GetMetadata_FunctionMetadata{},
		},
		"Name": {
			in: &tfprotov6.FunctionMetadata{
				Name: "test",
			},
			expected: &tfplugin6.GetMetadata_FunctionMetadata{
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
				tfplugin6.GetMetadata_FunctionMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
