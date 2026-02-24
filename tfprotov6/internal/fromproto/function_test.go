// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
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
					testTfplugin6DynamicValue(),
				},
			},
			expected: &tfprotov6.CallFunctionRequest{
				Arguments: []*tfprotov6.DynamicValue{
					testTfprotov6DynamicValue(),
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
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.GetFunctionsRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
