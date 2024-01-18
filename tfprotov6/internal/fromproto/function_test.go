// Copyright (c) HashiCorp, Inc.

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestCallFunctionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.CallFunction_Request
		expected *tfprotov6.CallFunctionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfplugin6.CallFunction_Request{},
			expected: &tfprotov6.CallFunctionRequest{
				Arguments: []*tfprotov6.DynamicValue{},
			},
		},
		"Arguments": {
			in: &tfplugin6.CallFunction_Request{
				Arguments: []*tfplugin6.DynamicValue{
					testTfplugin6DynamicValue(t,
						tftypes.Object{},
						tftypes.NewValue(tftypes.Object{}, nil),
					),
				},
			},
			expected: &tfprotov6.CallFunctionRequest{
				Arguments: []*tfprotov6.DynamicValue{
					testTfprotov6DynamicValue(t,
						tftypes.Object{},
						tftypes.NewValue(tftypes.Object{}, nil),
					),
				},
			},
		},
		"Name": {
			in: &tfplugin6.CallFunction_Request{
				Name: "test",
			},
			expected: &tfprotov6.CallFunctionRequest{
				Arguments: []*tfprotov6.DynamicValue{},
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
		in       *tfplugin6.GetFunctions_Request
		expected *tfprotov6.GetFunctionsRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.GetFunctions_Request{},
			expected: &tfprotov6.GetFunctionsRequest{},
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
