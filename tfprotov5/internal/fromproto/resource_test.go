// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestApplyResourceChangeRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ApplyResourceChange_Request
		expected *tfprotov5.ApplyResourceChangeRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ApplyResourceChange_Request{},
			expected: &tfprotov5.ApplyResourceChangeRequest{},
		},
		"Config": {
			in: &tfplugin5.ApplyResourceChange_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"PlannedPrivate": {
			in: &tfplugin5.ApplyResourceChange_Request{
				PlannedPrivate: []byte("{}"),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				PlannedPrivate: []byte("{}"),
			},
		},
		"PlannedState": {
			in: &tfplugin5.ApplyResourceChange_Request{
				PlannedState: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				PlannedState: testTfprotov5DynamicValue(),
			},
		},
		"PriorState": {
			in: &tfplugin5.ApplyResourceChange_Request{
				PriorState: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				PriorState: testTfprotov5DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin5.ApplyResourceChange_Request{
				ProviderMeta: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				ProviderMeta: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ApplyResourceChange_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				TypeName: "test",
			},
		},
		"PlannedIdentity": {
			in: &tfplugin5.ApplyResourceChange_Request{
				PlannedIdentity: testTfplugin5ResourceIdentityData(),
			},
			expected: &tfprotov5.ApplyResourceChangeRequest{
				PlannedIdentity: testTfprotov5ResourceIdentityData(),
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
		in       *tfplugin5.ImportResourceState_Request
		expected *tfprotov5.ImportResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ImportResourceState_Request{},
			expected: &tfprotov5.ImportResourceStateRequest{},
		},
		"Id": {
			in: &tfplugin5.ImportResourceState_Request{
				Id: "test",
			},
			expected: &tfprotov5.ImportResourceStateRequest{
				ID: "test",
			},
		},
		"TypeName": {
			in: &tfplugin5.ImportResourceState_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ImportResourceStateRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.ImportResourceState_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov5.ImportResourceStateRequest{
				ClientCapabilities: &tfprotov5.ImportResourceStateClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"Identity": {
			in: &tfplugin5.ImportResourceState_Request{
				Identity: testTfplugin5ResourceIdentityData(),
			},
			expected: &tfprotov5.ImportResourceStateRequest{
				Identity: testTfprotov5ResourceIdentityData(),
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
		in       *tfplugin5.MoveResourceState_Request
		expected *tfprotov5.MoveResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.MoveResourceState_Request{},
			expected: &tfprotov5.MoveResourceStateRequest{},
		},
		"SourcePrivate": {
			in: &tfplugin5.MoveResourceState_Request{
				SourcePrivate: []byte(`{}`),
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourcePrivate: []byte(`{}`),
			},
		},
		"SourceProviderAddress": {
			in: &tfplugin5.MoveResourceState_Request{
				SourceProviderAddress: "test",
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourceProviderAddress: "test",
			},
		},
		"SourceSchemaVersion": {
			in: &tfplugin5.MoveResourceState_Request{
				SourceSchemaVersion: 123,
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourceSchemaVersion: 123,
			},
		},
		"SourceState": {
			in: &tfplugin5.MoveResourceState_Request{
				SourceState: testTfplugin5RawState(t, []byte("{}")),
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourceState: testTfprotov5RawState(t, []byte("{}")),
			},
		},
		"SourceTypeName": {
			in: &tfplugin5.MoveResourceState_Request{
				SourceTypeName: "test",
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourceTypeName: "test",
			},
		},
		"TargetTypeName": {
			in: &tfplugin5.MoveResourceState_Request{
				TargetTypeName: "test",
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				TargetTypeName: "test",
			},
		},
		"SourceIdentity": {
			in: &tfplugin5.MoveResourceState_Request{
				SourceIdentity: testTfplugin5RawState(t, []byte("{}")),
			},
			expected: &tfprotov5.MoveResourceStateRequest{
				SourceIdentity: testTfprotov5RawState(t, []byte("{}")),
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
		in       *tfplugin5.PlanResourceChange_Request
		expected *tfprotov5.PlanResourceChangeRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.PlanResourceChange_Request{},
			expected: &tfprotov5.PlanResourceChangeRequest{},
		},
		"Config": {
			in: &tfplugin5.PlanResourceChange_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"PriorPrivate": {
			in: &tfplugin5.PlanResourceChange_Request{
				PriorPrivate: []byte("{}"),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				PriorPrivate: []byte("{}"),
			},
		},
		"PriorState": {
			in: &tfplugin5.PlanResourceChange_Request{
				PriorState: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				PriorState: testTfprotov5DynamicValue(),
			},
		},
		"ProposedNewState": {
			in: &tfplugin5.PlanResourceChange_Request{
				ProposedNewState: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				ProposedNewState: testTfprotov5DynamicValue(),
			},
		},
		"ProviderMeta": {
			in: &tfplugin5.PlanResourceChange_Request{
				ProviderMeta: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				ProviderMeta: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.PlanResourceChange_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.PlanResourceChange_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				ClientCapabilities: &tfprotov5.PlanResourceChangeClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"PriorIdentity": {
			in: &tfplugin5.PlanResourceChange_Request{
				PriorIdentity: testTfplugin5ResourceIdentityData(),
			},
			expected: &tfprotov5.PlanResourceChangeRequest{
				PriorIdentity: testTfprotov5ResourceIdentityData(),
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
		in       *tfplugin5.ReadResource_Request
		expected *tfprotov5.ReadResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ReadResource_Request{},
			expected: &tfprotov5.ReadResourceRequest{},
		},
		"CurrentState": {
			in: &tfplugin5.ReadResource_Request{
				CurrentState: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ReadResourceRequest{
				CurrentState: testTfprotov5DynamicValue(),
			},
		},
		"Private": {
			in: &tfplugin5.ReadResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov5.ReadResourceRequest{
				Private: []byte("{}"),
			},
		},
		"ProviderMeta": {
			in: &tfplugin5.ReadResource_Request{
				ProviderMeta: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ReadResourceRequest{
				ProviderMeta: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ReadResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ReadResourceRequest{
				TypeName: "test",
			},
		},
		"ClientCapabilities": {
			in: &tfplugin5.ReadResource_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					DeferralAllowed: true,
				},
			},
			expected: &tfprotov5.ReadResourceRequest{
				ClientCapabilities: &tfprotov5.ReadResourceClientCapabilities{
					DeferralAllowed: true,
				},
			},
		},
		"CurrentIdentity": {
			in: &tfplugin5.ReadResource_Request{
				CurrentIdentity: testTfplugin5ResourceIdentityData(),
			},
			expected: &tfprotov5.ReadResourceRequest{
				CurrentIdentity: testTfprotov5ResourceIdentityData(),
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
		in       *tfplugin5.UpgradeResourceState_Request
		expected *tfprotov5.UpgradeResourceStateRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.UpgradeResourceState_Request{},
			expected: &tfprotov5.UpgradeResourceStateRequest{},
		},
		"RawState": {
			in: &tfplugin5.UpgradeResourceState_Request{
				RawState: testTfplugin5RawState(t, []byte("{}")),
			},
			expected: &tfprotov5.UpgradeResourceStateRequest{
				RawState: testTfprotov5RawState(t, []byte("{}")),
			},
		},
		"TypeName": {
			in: &tfplugin5.UpgradeResourceState_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.UpgradeResourceStateRequest{
				TypeName: "test",
			},
		},
		"Version": {
			in: &tfplugin5.UpgradeResourceState_Request{
				Version: 123,
			},
			expected: &tfprotov5.UpgradeResourceStateRequest{
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
		in       *tfplugin5.UpgradeResourceIdentity_Request
		expected *tfprotov5.UpgradeResourceIdentityRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.UpgradeResourceIdentity_Request{},
			expected: &tfprotov5.UpgradeResourceIdentityRequest{},
		},
		"RawIdentity": {
			in: &tfplugin5.UpgradeResourceIdentity_Request{
				RawIdentity: testTfplugin5RawState(t, []byte("{}")),
			},
			expected: &tfprotov5.UpgradeResourceIdentityRequest{
				RawIdentity: testTfprotov5RawState(t, []byte("{}")),
			},
		},
		"TypeName": {
			in: &tfplugin5.UpgradeResourceIdentity_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.UpgradeResourceIdentityRequest{
				TypeName: "test",
			},
		},
		"Version": {
			in: &tfplugin5.UpgradeResourceIdentity_Request{
				Version: 123,
			},
			expected: &tfprotov5.UpgradeResourceIdentityRequest{
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

func TestValidateResourceTypeConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ValidateResourceTypeConfig_Request
		expected *tfprotov5.ValidateResourceTypeConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ValidateResourceTypeConfig_Request{},
			expected: &tfprotov5.ValidateResourceTypeConfigRequest{},
		},
		"ClientCapabilities": {
			in: &tfplugin5.ValidateResourceTypeConfig_Request{
				ClientCapabilities: &tfplugin5.ClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
			},
			expected: &tfprotov5.ValidateResourceTypeConfigRequest{
				ClientCapabilities: &tfprotov5.ValidateResourceTypeConfigClientCapabilities{
					WriteOnlyAttributesAllowed: true,
				},
			},
		},
		"Config": {
			in: &tfplugin5.ValidateResourceTypeConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateResourceTypeConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ValidateResourceTypeConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ValidateResourceTypeConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateResourceTypeConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
