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

func TestApplyResourceChangeRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ApplyResourceChange_Request
		expected *tfprotov6.ApplyResourceChangeRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ApplyResourceChange_Request{},
			expected: &tfprotov6.ApplyResourceChangeRequest{},
		},
		"Config": {
			in: &tfplugin6.ApplyResourceChange_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"PlannedPrivate": {
			in: &tfplugin6.ApplyResourceChange_Request{
				PlannedPrivate: []byte("{}"),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				PlannedPrivate: []byte("{}"),
			},
		},
		"PlannedState": {
			in: &tfplugin6.ApplyResourceChange_Request{
				PlannedState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				PlannedState: testTfprotov6DynamicValue(),
			},
		},
		"PriorState": {
			in: &tfplugin6.ApplyResourceChange_Request{
				PriorState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				PriorState: testTfprotov6DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin6.ApplyResourceChange_Request{
				ProviderMeta: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				ProviderMeta: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ApplyResourceChange_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				TypeName: "test",
			},
		},
		"PlannedIdentity": {
			in: &tfplugin6.ApplyResourceChange_Request{
				PlannedIdentity: testTfplugin6ResourceIdentityData(),
			},
			expected: &tfprotov6.ApplyResourceChangeRequest{
				PlannedIdentity: testTfprotov6ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ApplyResourceChangeRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestImportResourceStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ImportResourceState_Request
		expected *tfprotov6.ImportResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ImportResourceState_Request{},
			expected: &tfprotov6.ImportResourceStateRequest{},
		},
		"Id": {
			in: &tfplugin6.ImportResourceState_Request{
				Id: "test",
			},
			expected: &tfprotov6.ImportResourceStateRequest{
				ID: "test",
			},
		},
		"TypeName": {
			in: &tfplugin6.ImportResourceState_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ImportResourceStateRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.ImportResourceState_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.ImportResourceStateRequest{
				ClientCapabilities: &tfprotov6.ImportResourceStateClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"Identity": {
			in: &tfplugin6.ImportResourceState_Request{
				Identity: testTfplugin6ResourceIdentityData(),
			},
			expected: &tfprotov6.ImportResourceStateRequest{
				Identity: testTfprotov6ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ImportResourceStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestMoveResourceStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.MoveResourceState_Request
		expected *tfprotov6.MoveResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.MoveResourceState_Request{},
			expected: &tfprotov6.MoveResourceStateRequest{},
		},
		"SourcePrivate": {
			in: &tfplugin6.MoveResourceState_Request{
				SourcePrivate: []byte(`{}`),
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourcePrivate: []byte(`{}`),
			},
		},
		"SourceProviderAddress": {
			in: &tfplugin6.MoveResourceState_Request{
				SourceProviderAddress: "test",
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourceProviderAddress: "test",
			},
		},
		"SourceSchemaVersion": {
			in: &tfplugin6.MoveResourceState_Request{
				SourceSchemaVersion: 123,
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourceSchemaVersion: 123,
			},
		},
		"SourceState": {
			in: &tfplugin6.MoveResourceState_Request{
				SourceState: testTfplugin6RawState(t, []byte("{}")),
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourceState: testTfprotov6RawState(t, []byte("{}")),
			},
		},
		"SourceTypeName": {
			in: &tfplugin6.MoveResourceState_Request{
				SourceTypeName: "test",
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourceTypeName: "test",
			},
		},
		"TargetTypeName": {
			in: &tfplugin6.MoveResourceState_Request{
				TargetTypeName: "test",
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				TargetTypeName: "test",
			},
		},
		"SourceIdentity": {
			in: &tfplugin6.MoveResourceState_Request{
				SourceIdentity: testTfplugin6ResourceIdentityData(),
			},
			expected: &tfprotov6.MoveResourceStateRequest{
				SourceIdentity: testTfprotov6ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.MoveResourceStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanResourceChangeRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.PlanResourceChange_Request
		expected *tfprotov6.PlanResourceChangeRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.PlanResourceChange_Request{},
			expected: &tfprotov6.PlanResourceChangeRequest{},
		},
		"Config": {
			in: &tfplugin6.PlanResourceChange_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"PriorPrivate": {
			in: &tfplugin6.PlanResourceChange_Request{
				PriorPrivate: []byte("{}"),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				PriorPrivate: []byte("{}"),
			},
		},
		"PriorState": {
			in: &tfplugin6.PlanResourceChange_Request{
				PriorState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				PriorState: testTfprotov6DynamicValue(),
			},
		},
		"ProposedNewState": {
			in: &tfplugin6.PlanResourceChange_Request{
				ProposedNewState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				ProposedNewState: testTfprotov6DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin6.PlanResourceChange_Request{
				ProviderMeta: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				ProviderMeta: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.PlanResourceChange_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.PlanResourceChange_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				ClientCapabilities: &tfprotov6.PlanResourceChangeClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"PriorIdentity": {
			in: &tfplugin6.PlanResourceChange_Request{
				PriorIdentity: testTfplugin6ResourceIdentityData(),
			},
			expected: &tfprotov6.PlanResourceChangeRequest{
				PriorIdentity: testTfprotov6ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.PlanResourceChangeRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadResourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ReadResource_Request
		expected *tfprotov6.ReadResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ReadResource_Request{},
			expected: &tfprotov6.ReadResourceRequest{},
		},
		"CurrentState": {
			in: &tfplugin6.ReadResource_Request{
				CurrentState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ReadResourceRequest{
				CurrentState: testTfprotov6DynamicValue(),
			},
		},
		"Private": {
			in: &tfplugin6.ReadResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov6.ReadResourceRequest{
				Private: []byte("{}"),
			},
		},
		"ProviderMeta": {
			in: &tfplugin6.ReadResource_Request{
				ProviderMeta: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ReadResourceRequest{
				ProviderMeta: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ReadResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ReadResourceRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin6.ReadResource_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov6.ReadResourceRequest{
				ClientCapabilities: &tfprotov6.ReadResourceClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"CurrentIdentity": {
			in: &tfplugin6.ReadResource_Request{
				CurrentIdentity: testTfplugin6ResourceIdentityData(),
			},
			expected: &tfprotov6.ReadResourceRequest{
				CurrentIdentity: testTfprotov6ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ReadResourceRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestUpgradeResourceStateRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.UpgradeResourceState_Request
		expected *tfprotov6.UpgradeResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.UpgradeResourceState_Request{},
			expected: &tfprotov6.UpgradeResourceStateRequest{},
		},
		"RawState": {
			in: &tfplugin6.UpgradeResourceState_Request{
				RawState: testTfplugin6RawState(t, []byte("{}")),
			},
			expected: &tfprotov6.UpgradeResourceStateRequest{
				RawState: testTfprotov6RawState(t, []byte("{}")),
			},
		},
		"TypeName": {
			in: &tfplugin6.UpgradeResourceState_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.UpgradeResourceStateRequest{
				TypeName: "test",
			},
		},
		"Version": {
			in: &tfplugin6.UpgradeResourceState_Request{
				Version: 123,
			},
			expected: &tfprotov6.UpgradeResourceStateRequest{
				Version: 123,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.UpgradeResourceStateRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestUpgradeResourceIdentityRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.UpgradeResourceIdentity_Request
		expected *tfprotov6.UpgradeResourceIdentityRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.UpgradeResourceIdentity_Request{},
			expected: &tfprotov6.UpgradeResourceIdentityRequest{},
		},
		"RawIdentity": {
			in: &tfplugin6.UpgradeResourceIdentity_Request{
				RawIdentity: []byte("{}"),
			},
			expected: &tfprotov6.UpgradeResourceIdentityRequest{
				RawIdentity: testTfprotov6RawIdentity(t, []byte("{}")),
			},
		},
		"TypeName": {
			in: &tfplugin6.UpgradeResourceIdentity_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.UpgradeResourceIdentityRequest{
				TypeName: "test",
			},
		},
		"Version": {
			in: &tfplugin6.UpgradeResourceIdentity_Request{
				Version: 123,
			},
			expected: &tfprotov6.UpgradeResourceIdentityRequest{
				Version: 123,
			},
		},
	}

	for name, testCase := range testCases {

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.UpgradeResourceIdentityRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateResourceConfig_Request
		expected *tfprotov6.ValidateResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateResourceConfig_Request{},
			expected: &tfprotov6.ValidateResourceConfigRequest{},
		},
		"ClientCapabilities": {
			in: &tfplugin6.ValidateResourceConfig_Request{
				ClientCapabilities: &tfplugin6.ClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
			},
			expected: &tfprotov6.ValidateResourceConfigRequest{
				ClientCapabilities: &tfprotov6.ValidateResourceConfigClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
			},
		},
		"Config": {
			in: &tfplugin6.ValidateResourceConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateResourceConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ValidateResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ValidateResourceConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateResourceConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
