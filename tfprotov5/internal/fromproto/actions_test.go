// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestValidateActionConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ValidateActionConfig_Request
		expected *tfprotov5.ValidateActionConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ValidateActionConfig_Request{},
			expected: &tfprotov5.ValidateActionConfigRequest{},
		},
		"Config": {
			in: &tfplugin5.ValidateActionConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateActionConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"ActionType": {
			in: &tfplugin5.ValidateActionConfig_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.ValidateActionConfigRequest{
				ActionType: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateActionConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanActionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.PlanAction_Request
		expected *tfprotov5.PlanActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.PlanAction_Request{},
			expected: &tfprotov5.PlanActionRequest{},
		},
		"ActionType": {
			in: &tfplugin5.PlanAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.PlanActionRequest{
				ActionType: "test",
			},
		},
		"Config": {
			in: &tfplugin5.PlanAction_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanActionRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.PlanAction_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov5.PlanActionRequest{
				ClientCapabilities: &tfprotov5.PlanActionClientCapabilities{
					DeferralAllowed: true,
				},
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

func TestInvokeActionRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.InvokeAction_Request
		expected *tfprotov5.InvokeActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.InvokeAction_Request{},
			expected: &tfprotov5.InvokeActionRequest{},
		},
		"ActionType": {
			in: &tfplugin5.InvokeAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.InvokeActionRequest{
				ActionType: "test",
			},
		},
		"Config": {
			in: &tfplugin5.InvokeAction_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.InvokeActionRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.InvokeAction_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{},
			},
			expected: &tfprotov5.InvokeActionRequest{
				ClientCapabilities: &tfprotov5.InvokeActionClientCapabilities{},
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
