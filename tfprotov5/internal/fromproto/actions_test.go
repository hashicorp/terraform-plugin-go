// Copyright (c) HashiCorp, Inc.
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
			in: &tfplugin5.ValidateActionConfig_Request{},
			expected: &tfprotov5.ValidateActionConfigRequest{
				LinkedResources: []*tfprotov5.LinkedResourceConfig{},
			},
		},
		"Config": {
			in: &tfplugin5.ValidateActionConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateActionConfigRequest{
				Config:          testTfprotov5DynamicValue(),
				LinkedResources: []*tfprotov5.LinkedResourceConfig{},
			},
		},
		"ActionType": {
			in: &tfplugin5.ValidateActionConfig_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.ValidateActionConfigRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov5.LinkedResourceConfig{},
			},
		},
		"LinkedResources": {
			in: &tfplugin5.ValidateActionConfig_Request{
				LinkedResources: []*tfplugin5.LinkedResourceConfig{
					{
						TypeName: "test_linked",
						Config:   testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.ValidateActionConfigRequest{
				LinkedResources: []*tfprotov5.LinkedResourceConfig{
					{
						TypeName: "test_linked",
						Config:   testTfprotov5DynamicValue(),
					},
				},
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
			in: &tfplugin5.PlanAction_Request{},
			expected: &tfprotov5.PlanActionRequest{
				LinkedResources: []*tfprotov5.ProposedLinkedResource{},
			},
		},
		"ActionType": {
			in: &tfplugin5.PlanAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.PlanActionRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov5.ProposedLinkedResource{},
			},
		},
		"Config": {
			in: &tfplugin5.PlanAction_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanActionRequest{
				Config:          testTfprotov5DynamicValue(),
				LinkedResources: []*tfprotov5.ProposedLinkedResource{},
			},
		},
		"LinkedResources - PriorState": {
			in: &tfplugin5.PlanAction_Request{
				LinkedResources: []*tfplugin5.PlanAction_Request_LinkedResource{
					{
						PriorState: testTfplugin5DynamicValue(),
					},
					{
						PriorState: testTfplugin5DynamicValue(),
					},
					{
						PriorState: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.PlanActionRequest{
				LinkedResources: []*tfprotov5.ProposedLinkedResource{
					{
						PriorState: testTfprotov5DynamicValue(),
					},
					{
						PriorState: testTfprotov5DynamicValue(),
					},
					{
						PriorState: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfplugin5.PlanAction_Request{
				LinkedResources: []*tfplugin5.PlanAction_Request_LinkedResource{
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.PlanActionRequest{
				LinkedResources: []*tfprotov5.ProposedLinkedResource{
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - Config": {
			in: &tfplugin5.PlanAction_Request{
				LinkedResources: []*tfplugin5.PlanAction_Request_LinkedResource{
					{
						Config: testTfplugin5DynamicValue(),
					},
					{
						Config: testTfplugin5DynamicValue(),
					},
					{
						Config: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.PlanActionRequest{
				LinkedResources: []*tfprotov5.ProposedLinkedResource{
					{
						Config: testTfprotov5DynamicValue(),
					},
					{
						Config: testTfprotov5DynamicValue(),
					},
					{
						Config: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PriorIdentity": {
			in: &tfplugin5.PlanAction_Request{
				LinkedResources: []*tfplugin5.PlanAction_Request_LinkedResource{
					{
						PriorIdentity: testTfplugin5ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfplugin5ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfplugin5ResourceIdentityData(),
					},
				},
			},
			expected: &tfprotov5.PlanActionRequest{
				LinkedResources: []*tfprotov5.ProposedLinkedResource{
					{
						PriorIdentity: testTfprotov5ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfprotov5ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfprotov5ResourceIdentityData(),
					},
				},
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
				LinkedResources: []*tfprotov5.ProposedLinkedResource{},
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
			in: &tfplugin5.InvokeAction_Request{},
			expected: &tfprotov5.InvokeActionRequest{
				LinkedResources: []*tfprotov5.InvokeLinkedResource{},
			},
		},
		"ActionType": {
			in: &tfplugin5.InvokeAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov5.InvokeActionRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov5.InvokeLinkedResource{},
			},
		},
		"Config": {
			in: &tfplugin5.InvokeAction_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.InvokeActionRequest{
				Config:          testTfprotov5DynamicValue(),
				LinkedResources: []*tfprotov5.InvokeLinkedResource{},
			},
		},
		"LinkedResources - PriorState": {
			in: &tfplugin5.InvokeAction_Request{
				LinkedResources: []*tfplugin5.InvokeAction_Request_LinkedResource{
					{
						PriorState: testTfplugin5DynamicValue(),
					},
					{
						PriorState: testTfplugin5DynamicValue(),
					},
					{
						PriorState: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.InvokeActionRequest{
				LinkedResources: []*tfprotov5.InvokeLinkedResource{
					{
						PriorState: testTfprotov5DynamicValue(),
					},
					{
						PriorState: testTfprotov5DynamicValue(),
					},
					{
						PriorState: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfplugin5.InvokeAction_Request{
				LinkedResources: []*tfplugin5.InvokeAction_Request_LinkedResource{
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.InvokeActionRequest{
				LinkedResources: []*tfprotov5.InvokeLinkedResource{
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - Config": {
			in: &tfplugin5.InvokeAction_Request{
				LinkedResources: []*tfplugin5.InvokeAction_Request_LinkedResource{
					{
						Config: testTfplugin5DynamicValue(),
					},
					{
						Config: testTfplugin5DynamicValue(),
					},
					{
						Config: testTfplugin5DynamicValue(),
					},
				},
			},
			expected: &tfprotov5.InvokeActionRequest{
				LinkedResources: []*tfprotov5.InvokeLinkedResource{
					{
						Config: testTfprotov5DynamicValue(),
					},
					{
						Config: testTfprotov5DynamicValue(),
					},
					{
						Config: testTfprotov5DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedIdentity": {
			in: &tfplugin5.InvokeAction_Request{
				LinkedResources: []*tfplugin5.InvokeAction_Request_LinkedResource{
					{
						PlannedIdentity: testTfplugin5ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfplugin5ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfplugin5ResourceIdentityData(),
					},
				},
			},
			expected: &tfprotov5.InvokeActionRequest{
				LinkedResources: []*tfprotov5.InvokeLinkedResource{
					{
						PlannedIdentity: testTfprotov5ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfprotov5ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfprotov5ResourceIdentityData(),
					},
				},
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
