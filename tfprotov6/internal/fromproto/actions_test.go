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
			in: &tfplugin6.ValidateActionConfig_Request{},
			expected: &tfprotov6.ValidateActionConfigRequest{
				LinkedResources: []*tfprotov6.LinkedResourceConfig{},
			},
		},
		"Config": {
			in: &tfplugin6.ValidateActionConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateActionConfigRequest{
				Config:          testTfprotov6DynamicValue(),
				LinkedResources: []*tfprotov6.LinkedResourceConfig{},
			},
		},
		"ActionType": {
			in: &tfplugin6.ValidateActionConfig_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.ValidateActionConfigRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov6.LinkedResourceConfig{},
			},
		},
		"LinkedResources": {
			in: &tfplugin6.ValidateActionConfig_Request{
				LinkedResources: []*tfplugin6.LinkedResourceConfig{
					{
						TypeName: "test_linked",
						Config:   testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.ValidateActionConfigRequest{
				LinkedResources: []*tfprotov6.LinkedResourceConfig{
					{
						TypeName: "test_linked",
						Config:   testTfprotov6DynamicValue(),
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
		in       *tfplugin6.PlanAction_Request
		expected *tfprotov6.PlanActionRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfplugin6.PlanAction_Request{},
			expected: &tfprotov6.PlanActionRequest{
				LinkedResources: []*tfprotov6.ProposedLinkedResource{},
			},
		},
		"ActionType": {
			in: &tfplugin6.PlanAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.PlanActionRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov6.ProposedLinkedResource{},
			},
		},
		"Config": {
			in: &tfplugin6.PlanAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanActionRequest{
				Config:          testTfprotov6DynamicValue(),
				LinkedResources: []*tfprotov6.ProposedLinkedResource{},
			},
		},
		"LinkedResources - PriorState": {
			in: &tfplugin6.PlanAction_Request{
				LinkedResources: []*tfplugin6.PlanAction_Request_LinkedResource{
					{
						PriorState: testTfplugin6DynamicValue(),
					},
					{
						PriorState: testTfplugin6DynamicValue(),
					},
					{
						PriorState: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.PlanActionRequest{
				LinkedResources: []*tfprotov6.ProposedLinkedResource{
					{
						PriorState: testTfprotov6DynamicValue(),
					},
					{
						PriorState: testTfprotov6DynamicValue(),
					},
					{
						PriorState: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfplugin6.PlanAction_Request{
				LinkedResources: []*tfplugin6.PlanAction_Request_LinkedResource{
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.PlanActionRequest{
				LinkedResources: []*tfprotov6.ProposedLinkedResource{
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - Config": {
			in: &tfplugin6.PlanAction_Request{
				LinkedResources: []*tfplugin6.PlanAction_Request_LinkedResource{
					{
						Config: testTfplugin6DynamicValue(),
					},
					{
						Config: testTfplugin6DynamicValue(),
					},
					{
						Config: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.PlanActionRequest{
				LinkedResources: []*tfprotov6.ProposedLinkedResource{
					{
						Config: testTfprotov6DynamicValue(),
					},
					{
						Config: testTfprotov6DynamicValue(),
					},
					{
						Config: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PriorIdentity": {
			in: &tfplugin6.PlanAction_Request{
				LinkedResources: []*tfplugin6.PlanAction_Request_LinkedResource{
					{
						PriorIdentity: testTfplugin6ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfplugin6ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfplugin6ResourceIdentityData(),
					},
				},
			},
			expected: &tfprotov6.PlanActionRequest{
				LinkedResources: []*tfprotov6.ProposedLinkedResource{
					{
						PriorIdentity: testTfprotov6ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfprotov6ResourceIdentityData(),
					},
					{
						PriorIdentity: testTfprotov6ResourceIdentityData(),
					},
				},
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
				LinkedResources: []*tfprotov6.ProposedLinkedResource{},
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
			in: &tfplugin6.InvokeAction_Request{},
			expected: &tfprotov6.InvokeActionRequest{
				LinkedResources: []*tfprotov6.InvokeLinkedResource{},
			},
		},
		"ActionType": {
			in: &tfplugin6.InvokeAction_Request{
				ActionType: "test",
			},
			expected: &tfprotov6.InvokeActionRequest{
				ActionType:      "test",
				LinkedResources: []*tfprotov6.InvokeLinkedResource{},
			},
		},
		"Config": {
			in: &tfplugin6.InvokeAction_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.InvokeActionRequest{
				Config:          testTfprotov6DynamicValue(),
				LinkedResources: []*tfprotov6.InvokeLinkedResource{},
			},
		},
		"LinkedResources - PriorState": {
			in: &tfplugin6.InvokeAction_Request{
				LinkedResources: []*tfplugin6.InvokeAction_Request_LinkedResource{
					{
						PriorState: testTfplugin6DynamicValue(),
					},
					{
						PriorState: testTfplugin6DynamicValue(),
					},
					{
						PriorState: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.InvokeActionRequest{
				LinkedResources: []*tfprotov6.InvokeLinkedResource{
					{
						PriorState: testTfprotov6DynamicValue(),
					},
					{
						PriorState: testTfprotov6DynamicValue(),
					},
					{
						PriorState: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfplugin6.InvokeAction_Request{
				LinkedResources: []*tfplugin6.InvokeAction_Request_LinkedResource{
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.InvokeActionRequest{
				LinkedResources: []*tfprotov6.InvokeLinkedResource{
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - Config": {
			in: &tfplugin6.InvokeAction_Request{
				LinkedResources: []*tfplugin6.InvokeAction_Request_LinkedResource{
					{
						Config: testTfplugin6DynamicValue(),
					},
					{
						Config: testTfplugin6DynamicValue(),
					},
					{
						Config: testTfplugin6DynamicValue(),
					},
				},
			},
			expected: &tfprotov6.InvokeActionRequest{
				LinkedResources: []*tfprotov6.InvokeLinkedResource{
					{
						Config: testTfprotov6DynamicValue(),
					},
					{
						Config: testTfprotov6DynamicValue(),
					},
					{
						Config: testTfprotov6DynamicValue(),
					},
				},
			},
		},
		"LinkedResources - PlannedIdentity": {
			in: &tfplugin6.InvokeAction_Request{
				LinkedResources: []*tfplugin6.InvokeAction_Request_LinkedResource{
					{
						PlannedIdentity: testTfplugin6ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfplugin6ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfplugin6ResourceIdentityData(),
					},
				},
			},
			expected: &tfprotov6.InvokeActionRequest{
				LinkedResources: []*tfprotov6.InvokeLinkedResource{
					{
						PlannedIdentity: testTfprotov6ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfprotov6ResourceIdentityData(),
					},
					{
						PlannedIdentity: testTfprotov6ResourceIdentityData(),
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
