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

func TestValidateActionConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateActionConfig_Request
		expected *tfprotov6.ValidateActionConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateActionConfig_Request{},
			expected: &tfprotov6.ValidateActionConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateActionConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateActionConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"ActionType": {
			in: &tfplugin6.ValidateActionConfig_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.ValidateActionConfigRequest{
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
		"ActionType": {
			in: &tfplugin6.PlanAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.PlanActionRequest{
				ActionType: "test",
			},
		},
		"Config": {
			in: &tfplugin6.PlanAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanActionRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.PlanAction_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.PlanActionRequest{
				ClientCapabilities: &tfprotov6.PlanActionClientCapabilities{
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
		"ActionType": {
			in: &tfplugin6.InvokeAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.InvokeActionRequest{
				ActionType: "test",
			},
		},
		"Config": {
			in: &tfplugin6.InvokeAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.InvokeActionRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.InvokeAction_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{},
			},
			expected: &tfprotov6.InvokeActionRequest{
				ClientCapabilities: &tfprotov6.InvokeActionClientCapabilities{},
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
