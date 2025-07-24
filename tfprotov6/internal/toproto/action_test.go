// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
)

func TestGetMetadata_ActionMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ActionMetadata
		expected *tfplugin6.GetMetadata_ActionMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ActionMetadata{},
			expected: &tfplugin6.GetMetadata_ActionMetadata{},
		},
		"TypeName": {
			in: &tfprotov6.ActionMetadata{
				TypeName: "test",
			},
			expected: &tfplugin6.GetMetadata_ActionMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_ActionMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.GetMetadata_ActionMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateActionConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ValidateActionConfigResponse
		expected *tfplugin6.ValidateActionConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ValidateActionConfigResponse{},
			expected: &tfplugin6.ValidateActionConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ValidateActionConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ValidateActionConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ValidateActionConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ValidateActionConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanAction_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.PlanActionResponse
		expected *tfplugin6.PlanAction_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.PlanActionResponse{},
			expected: &tfplugin6.PlanAction_Response{
				LinkedResources: []*tfplugin6.PlanAction_Response_LinkedResource{},
				Diagnostics:     []*tfplugin6.Diagnostic{},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfprotov6.PlanActionResponse{
				LinkedResources: []*tfprotov6.PlannedLinkedResource{
					{
						PlannedState: testTfprotov6DynamicValue(),
					},
				},
			},
			expected: &tfplugin6.PlanAction_Response{
				LinkedResources: []*tfplugin6.PlanAction_Response_LinkedResource{
					{
						PlannedState: testTfplugin6DynamicValue(),
					},
				},
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"LinkedResources - PlannedIdentity": {
			in: &tfprotov6.PlanActionResponse{
				LinkedResources: []*tfprotov6.PlannedLinkedResource{
					{
						PlannedIdentity: testTfprotov6ResourceIdentityData(),
					},
				},
			},
			expected: &tfplugin6.PlanAction_Response{
				LinkedResources: []*tfplugin6.PlanAction_Response_LinkedResource{
					{
						PlannedIdentity: testTfplugin6ResourceIdentityData(),
					},
				},
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.PlanActionResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.PlanAction_Response{
				LinkedResources: []*tfplugin6.PlanAction_Response_LinkedResource{},
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Deferred": {
			in: &tfprotov6.PlanActionResponse{
				Deferred: &tfprotov6.Deferred{
					Reason: tfprotov6.DeferredReasonProviderConfigUnknown,
				},
			},
			expected: &tfplugin6.PlanAction_Response{
				LinkedResources: []*tfplugin6.PlanAction_Response_LinkedResource{},
				Diagnostics:     []*tfplugin6.Diagnostic{},
				Deferred: &tfplugin6.Deferred{
					Reason: tfplugin6.Deferred_PROVIDER_CONFIG_UNKNOWN,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.PlanAction_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.PlanAction_Response{},
				tfplugin6.PlanAction_Response_LinkedResource{},
				tfplugin6.Deferred{},
				tfplugin6.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestInvokeAction_InvokeActionEvent(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.InvokeActionEvent
		expected *tfplugin6.InvokeAction_Event
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"ProgressInvokeActionEventType - Message": {
			in: &tfprotov6.InvokeActionEvent{
				Type: tfprotov6.ProgressInvokeActionEventType{
					Message: "test message",
				},
			},
			expected: &tfplugin6.InvokeAction_Event{
				Type: &tfplugin6.InvokeAction_Event_Progress_{
					Progress: &tfplugin6.InvokeAction_Event_Progress{
						Message: "test message",
					},
				},
			},
		},
		"CompletedInvokeActionEventType - Diagnostics": {
			in: &tfprotov6.InvokeActionEvent{
				Type: tfprotov6.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov6.NewLinkedResource{},
					Diagnostics: []*tfprotov6.Diagnostic{
						testTfprotov6Diagnostic,
					},
				},
			},
			expected: &tfplugin6.InvokeAction_Event{
				Type: &tfplugin6.InvokeAction_Event_Completed_{
					Completed: &tfplugin6.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin6.InvokeAction_Event_Completed_LinkedResource{},
						Diagnostics: []*tfplugin6.Diagnostic{
							testTfplugin6Diagnostic,
						},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - NewState": {
			in: &tfprotov6.InvokeActionEvent{
				Type: tfprotov6.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov6.NewLinkedResource{
						{
							NewState: testTfprotov6DynamicValue(),
						},
					},
					Diagnostics: []*tfprotov6.Diagnostic{},
				},
			},
			expected: &tfplugin6.InvokeAction_Event{
				Type: &tfplugin6.InvokeAction_Event_Completed_{
					Completed: &tfplugin6.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin6.InvokeAction_Event_Completed_LinkedResource{
							{
								NewState: testTfplugin6DynamicValue(),
							},
						},
						Diagnostics: []*tfplugin6.Diagnostic{},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - NewIdentity": {
			in: &tfprotov6.InvokeActionEvent{
				Type: tfprotov6.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov6.NewLinkedResource{
						{
							NewIdentity: testTfprotov6ResourceIdentityData(),
						},
					},
					Diagnostics: []*tfprotov6.Diagnostic{},
				},
			},
			expected: &tfplugin6.InvokeAction_Event{
				Type: &tfplugin6.InvokeAction_Event_Completed_{
					Completed: &tfplugin6.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin6.InvokeAction_Event_Completed_LinkedResource{
							{
								NewIdentity: testTfplugin6ResourceIdentityData(),
							},
						},
						Diagnostics: []*tfplugin6.Diagnostic{},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - RequiresReplace": {
			in: &tfprotov6.InvokeActionEvent{
				Type: tfprotov6.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov6.NewLinkedResource{
						{
							RequiresReplace: true,
						},
					},
					Diagnostics: []*tfprotov6.Diagnostic{},
				},
			},
			expected: &tfplugin6.InvokeAction_Event{
				Type: &tfplugin6.InvokeAction_Event_Completed_{
					Completed: &tfplugin6.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin6.InvokeAction_Event_Completed_LinkedResource{
							{
								RequiresReplace: true,
							},
						},
						Diagnostics: []*tfplugin6.Diagnostic{},
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.InvokeAction_InvokeActionEvent(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.InvokeAction_Event{},
				tfplugin6.InvokeAction_Event_Progress{},
				tfplugin6.InvokeAction_Event_Progress_{},
				tfplugin6.InvokeAction_Event_Completed{},
				tfplugin6.InvokeAction_Event_Completed_{},
				tfplugin6.InvokeAction_Event_Completed_LinkedResource{},
				tfplugin6.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
