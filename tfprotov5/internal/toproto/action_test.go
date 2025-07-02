// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
)

func TestGetMetadata_ActionMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ActionMetadata
		expected *tfplugin5.GetMetadata_ActionMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ActionMetadata{},
			expected: &tfplugin5.GetMetadata_ActionMetadata{},
		},
		"TypeName": {
			in: &tfprotov5.ActionMetadata{
				TypeName: "test",
			},
			expected: &tfplugin5.GetMetadata_ActionMetadata{
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
				tfplugin5.GetMetadata_ActionMetadata{},
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
		in       *tfprotov5.PlanActionResponse
		expected *tfplugin5.PlanAction_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.PlanActionResponse{},
			expected: &tfplugin5.PlanAction_Response{
				LinkedResources: []*tfplugin5.PlanAction_Response_LinkedResource{},
				Diagnostics:     []*tfplugin5.Diagnostic{},
			},
		},
		"LinkedResources - PlannedState": {
			in: &tfprotov5.PlanActionResponse{
				LinkedResources: []*tfprotov5.PlannedLinkedResource{
					{
						PlannedState: testTfprotov5DynamicValue(),
					},
				},
			},
			expected: &tfplugin5.PlanAction_Response{
				LinkedResources: []*tfplugin5.PlanAction_Response_LinkedResource{
					{
						PlannedState: testTfplugin5DynamicValue(),
					},
				},
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"LinkedResources - PlannedIdentity": {
			in: &tfprotov5.PlanActionResponse{
				LinkedResources: []*tfprotov5.PlannedLinkedResource{
					{
						PlannedIdentity: testTfprotov5ResourceIdentityData(),
					},
				},
			},
			expected: &tfplugin5.PlanAction_Response{
				LinkedResources: []*tfplugin5.PlanAction_Response_LinkedResource{
					{
						PlannedIdentity: testTfplugin5ResourceIdentityData(),
					},
				},
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.PlanActionResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.PlanAction_Response{
				LinkedResources: []*tfplugin5.PlanAction_Response_LinkedResource{},
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"Deferred": {
			in: &tfprotov5.PlanActionResponse{
				Deferred: &tfprotov5.Deferred{
					Reason: tfprotov5.DeferredReasonProviderConfigUnknown,
				},
			},
			expected: &tfplugin5.PlanAction_Response{
				LinkedResources: []*tfplugin5.PlanAction_Response_LinkedResource{},
				Diagnostics:     []*tfplugin5.Diagnostic{},
				Deferred: &tfplugin5.Deferred{
					Reason: tfplugin5.Deferred_PROVIDER_CONFIG_UNKNOWN,
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
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.PlanAction_Response{},
				tfplugin5.PlanAction_Response_LinkedResource{},
				tfplugin5.Deferred{},
				tfplugin5.ResourceIdentityData{},
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
		in       *tfprotov5.InvokeActionEvent
		expected *tfplugin5.InvokeAction_Event
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"ProgressInvokeActionEventType - Message": {
			in: &tfprotov5.InvokeActionEvent{
				Type: tfprotov5.ProgressInvokeActionEventType{
					Message: "test message",
				},
			},
			expected: &tfplugin5.InvokeAction_Event{
				Type: &tfplugin5.InvokeAction_Event_Progress_{
					Progress: &tfplugin5.InvokeAction_Event_Progress{
						Message: "test message",
					},
				},
			},
		},
		"CompletedInvokeActionEventType - Diagnostics": {
			in: &tfprotov5.InvokeActionEvent{
				Type: tfprotov5.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov5.NewLinkedResource{},
					Diagnostics: []*tfprotov5.Diagnostic{
						testTfprotov5Diagnostic,
					},
				},
			},
			expected: &tfplugin5.InvokeAction_Event{
				Type: &tfplugin5.InvokeAction_Event_Completed_{
					Completed: &tfplugin5.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin5.InvokeAction_Event_Completed_LinkedResource{},
						Diagnostics: []*tfplugin5.Diagnostic{
							testTfplugin5Diagnostic,
						},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - NewState": {
			in: &tfprotov5.InvokeActionEvent{
				Type: tfprotov5.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov5.NewLinkedResource{
						{
							NewState: testTfprotov5DynamicValue(),
						},
					},
					Diagnostics: []*tfprotov5.Diagnostic{},
				},
			},
			expected: &tfplugin5.InvokeAction_Event{
				Type: &tfplugin5.InvokeAction_Event_Completed_{
					Completed: &tfplugin5.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin5.InvokeAction_Event_Completed_LinkedResource{
							{
								NewState: testTfplugin5DynamicValue(),
							},
						},
						Diagnostics: []*tfplugin5.Diagnostic{},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - NewIdentity": {
			in: &tfprotov5.InvokeActionEvent{
				Type: tfprotov5.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov5.NewLinkedResource{
						{
							NewIdentity: testTfprotov5ResourceIdentityData(),
						},
					},
					Diagnostics: []*tfprotov5.Diagnostic{},
				},
			},
			expected: &tfplugin5.InvokeAction_Event{
				Type: &tfplugin5.InvokeAction_Event_Completed_{
					Completed: &tfplugin5.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin5.InvokeAction_Event_Completed_LinkedResource{
							{
								NewIdentity: testTfplugin5ResourceIdentityData(),
							},
						},
						Diagnostics: []*tfplugin5.Diagnostic{},
					},
				},
			},
		},
		"CompletedInvokeActionEventType - LinkedResources - RequiresReplace": {
			in: &tfprotov5.InvokeActionEvent{
				Type: tfprotov5.CompletedInvokeActionEventType{
					LinkedResources: []*tfprotov5.NewLinkedResource{
						{
							RequiresReplace: true,
						},
					},
					Diagnostics: []*tfprotov5.Diagnostic{},
				},
			},
			expected: &tfplugin5.InvokeAction_Event{
				Type: &tfplugin5.InvokeAction_Event_Completed_{
					Completed: &tfplugin5.InvokeAction_Event_Completed{
						LinkedResources: []*tfplugin5.InvokeAction_Event_Completed_LinkedResource{
							{
								RequiresReplace: true,
							},
						},
						Diagnostics: []*tfplugin5.Diagnostic{},
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
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.InvokeAction_Event{},
				tfplugin5.InvokeAction_Event_Progress{},
				tfplugin5.InvokeAction_Event_Progress_{},
				tfplugin5.InvokeAction_Event_Completed{},
				tfplugin5.InvokeAction_Event_Completed_{},
				tfplugin5.InvokeAction_Event_Completed_LinkedResource{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
