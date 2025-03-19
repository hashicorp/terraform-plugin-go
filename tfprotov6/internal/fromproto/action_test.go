// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestCancelActionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.CancelAction_Request
		expected *tfprotov6.CancelActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.CancelAction_Request{},
			expected: &tfprotov6.CancelActionRequest{},
		},
		"Cancellation Token": {
			in: &tfplugin6.CancelAction_Request{
				CancelationToken: "test token",
			},
			expected: &tfprotov6.CancelActionRequest{
				CancellationToken: "test token",
				CancellationType:  tfprotov6.ActionCancelTypeSoft,
			},
		},
		"Cancellation Type": {
			in: &tfplugin6.CancelAction_Request{
				CancelationToken: "test token",
				Type:             tfplugin6.CancelAction_HARD,
			},
			expected: &tfprotov6.CancelActionRequest{
				CancellationToken: "test token",
				CancellationType:  tfprotov6.ActionCancelTypeHard,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.CancelActionRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInvokeActionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.InvokeAction_Request
		expected *tfprotov6.InvokeActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.InvokeAction_Request{},
			expected: &tfprotov6.InvokeActionRequest{},
		},
		"Config": {
			in: &tfplugin6.InvokeAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.InvokeActionRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.InvokeAction_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.InvokeActionRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.InvokeActionRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanActionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.PlanAction_Request
		expected *tfprotov6.PlanActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.PlanAction_Request{},
			expected: &tfprotov6.PlanActionRequest{},
		},
		"Config": {
			in: &tfplugin6.PlanAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanActionRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.PlanAction_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.PlanActionRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.PlanActionRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
