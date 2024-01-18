// Copyright (c) HashiCorp, Inc.

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
			in: &tfprotov6.CallFunctionResponse{},
			expected: &tfplugin6.CallFunction_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.CallFunctionResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.CallFunction_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Result": {
			in: &tfprotov6.CallFunctionResponse{
				Result: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.CallFunction_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Result:      testTfplugin6DynamicValue(),
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
			got, _ := toproto.CallFunction_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.CallFunction_Response{},
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
			in:       &tfprotov6.Function{},
			expected: nil,
		},
		"Description": {
			in: &tfprotov6.Function{
				Description: "test",
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				Description: "test",
				Parameters:  []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.Function{
				DescriptionKind: tfprotov6.StringKindMarkdown,
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
				Parameters:      []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"DeprecationMessage": {
			in: &tfprotov6.Function{
				DeprecationMessage: "test",
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				DeprecationMessage: "test",
				Parameters:         []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
			},
		},
		"Parameters": {
			in: &tfprotov6.Function{
				Parameters: []*tfprotov6.FunctionParameter{
					{
						Type: tftypes.Bool,
					},
				},
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{
					{
						Type: []byte(`"bool"`),
					},
				},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
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
				// Return will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
				Summary: "test",
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
				Summary: "test",
			},
		},
		"VariadicParameter": {
			in: &tfprotov6.Function{
				Return: &tfprotov6.FunctionReturn{
					Type: tftypes.Bool,
				},
				VariadicParameter: &tfprotov6.FunctionParameter{
					Type: tftypes.Bool,
				},
			},
			expected: &tfplugin6.Function{
				Parameters: []*tfplugin6.Function_Parameter{},
				Return: &tfplugin6.Function_Return{
					Type: []byte(`"bool"`),
				},
				VariadicParameter: &tfplugin6.Function_Parameter{
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
			expected: nil,
		},
		"AllowNullValue": {
			in: &tfprotov6.FunctionParameter{
				AllowNullValue: true,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
				AllowNullValue: true,
				Type:           []byte(`"bool"`),
			},
		},
		"AllowUnknownValues": {
			in: &tfprotov6.FunctionParameter{
				AllowUnknownValues: true,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
				AllowUnknownValues: true,
				Type:               []byte(`"bool"`),
			},
		},
		"Description": {
			in: &tfprotov6.FunctionParameter{
				Description: "test",
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
				Description: "test",
				Type:        []byte(`"bool"`),
			},
		},
		"DescriptionKind": {
			in: &tfprotov6.FunctionParameter{
				DescriptionKind: tfprotov6.StringKindMarkdown,
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
				DescriptionKind: tfplugin6.StringKind_MARKDOWN,
				Type:            []byte(`"bool"`),
			},
		},
		"Name": {
			in: &tfprotov6.FunctionParameter{
				Name: "test",
				// Type will no longer be required in a future change.
				// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.Function_Parameter{
				Name: "test",
				Type: []byte(`"bool"`),
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
			expected: nil,
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
		name, testCase := name, testCase

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
