// Copyright (c) HashiCorp, Inc.

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestCallFunctionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.CallFunction_Request
		expected *tfprotov5.CallFunctionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfplugin5.CallFunction_Request{},
			expected: &tfprotov5.CallFunctionRequest{
				Arguments: []*tfprotov5.DynamicValue{},
			},
		},
		"Arguments": {
			in: &tfplugin5.CallFunction_Request{
				Arguments: []*tfplugin5.DynamicValue{
					testTfplugin5DynamicValue(t,
						tftypes.Object{},
						tftypes.NewValue(tftypes.Object{}, nil),
					),
				},
			},
			expected: &tfprotov5.CallFunctionRequest{
				Arguments: []*tfprotov5.DynamicValue{
					testTfprotov5DynamicValue(t,
						tftypes.Object{},
						tftypes.NewValue(tftypes.Object{}, nil),
					),
				},
			},
		},
		"Name": {
			in: &tfplugin5.CallFunction_Request{
				Name: "test",
			},
			expected: &tfprotov5.CallFunctionRequest{
				Arguments: []*tfprotov5.DynamicValue{},
				Name:      "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.CallFunctionRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetFunctionsRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.GetFunctions_Request
		expected *tfprotov5.GetFunctionsRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.GetFunctions_Request{},
			expected: &tfprotov5.GetFunctionsRequest{},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.GetFunctionsRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
