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
			in: &tfprotov5.CallFunctionResponse{},
			expected: &tfplugin5.CallFunction_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.CallFunctionResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.CallFunction_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"Result": {
			in: &tfprotov5.CallFunctionResponse{
				Result: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.CallFunction_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Result:      testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

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
			in:       &tfprotov5.Function{},
			expected: nil,
		},
		"Description": {
			in: &tfprotov5.Function{
				Description: "test",
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				Description: "test",
				Parameters:  []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.Function{
				DescriptionKind: tfprotov5.StringKindMarkdown,
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
				Parameters:      []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"DeprecationMessage": {
			in: &tfprotov5.Function{
				DeprecationMessage: "test",
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				DeprecationMessage: "test",
				Parameters:         []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"Parameters": {
			in: &tfprotov5.Function{
				Parameters: []*tfprotov5.FunctionParameter{
					{
						Type: tftypes.Bool,
					},
				},
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{
					{
						Type: []byte(`"bool"`),
					},
				},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
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
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
				Summary: "test",
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
				Summary: "test",
			},
		},
		"VariadicParameter": {
			in: &tfprotov5.Function{
				Return: &tfprotov5.FunctionReturn{
					Type: tftypes.Bool,
				},
				VariadicParameter: &tfprotov5.FunctionParameter{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin5.Function{
				Parameters: []*tfplugin5.Function_Parameter{},
				Return: &tfplugin5.Function_Return{
					Type: []byte(`"bool"`),
				},
				VariadicParameter: &tfplugin5.Function_Parameter{
					Type: []byte(`"bool"`),
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it will be removed
			// in a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.Function(testCase.in)

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
			expected: nil,
		},
		"AllowNullValue": {
			in: &tfprotov5.FunctionParameter{
				AllowNullValue: true,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				AllowNullValue: true,
				Type:           []byte(`"bool"`),
			},
		},
		"AllowUnknownValues": {
			in: &tfprotov5.FunctionParameter{
				AllowUnknownValues: true,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				AllowUnknownValues: true,
				Type:               []byte(`"bool"`),
			},
		},
		"Description": {
			in: &tfprotov5.FunctionParameter{
				Description: "test",
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				Description: "test",
				Type:        []byte(`"bool"`),
			},
		},
		"DescriptionKind": {
			in: &tfprotov5.FunctionParameter{
				DescriptionKind: tfprotov5.StringKindMarkdown,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				DescriptionKind: tfplugin5.StringKind_MARKDOWN,
				Type:            []byte(`"bool"`),
			},
		},
		"Name": {
			in: &tfprotov5.FunctionParameter{
				Name: "test",
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.Function_Parameter{
				Name: "test",
				Type: []byte(`"bool"`),
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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it will be removed
			// in a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.Function_Parameter(testCase.in)

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
			expected: nil,
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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it will be removed
			// in a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.Function_Return(testCase.in)

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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.GetFunctions_Response(testCase.in)

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
		name, testCase := name, testCase

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
