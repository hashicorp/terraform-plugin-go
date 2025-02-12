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
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestApplyResourceChange_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ApplyResourceChangeResponse
		expected *tfplugin5.ApplyResourceChange_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ApplyResourceChangeResponse{},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ApplyResourceChangeResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"Private": {
			in: &tfprotov5.ApplyResourceChangeResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"NewState": {
			in: &tfprotov5.ApplyResourceChangeResponse{
				NewState: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				NewState:    testTfplugin5DynamicValue(),
			},
		},
		"UnsafeToUseLegacyTypeSystem": {
			in: &tfprotov5.ApplyResourceChangeResponse{
				UnsafeToUseLegacyTypeSystem: true,
			},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics:      []*tfplugin5.Diagnostic{},
				LegacyTypeSystem: true,
			},
		},
		"NewIdentity": {
			in: &tfprotov5.ApplyResourceChangeResponse{
				NewIdentity: testTfprotov5ResourceIdentityData(),
			},
			expected: &tfplugin5.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				NewIdentity: testTfplugin5ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ApplyResourceChange_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.AttributePath{},
				tfplugin5.AttributePath_Step{},
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.ApplyResourceChange_Response{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestGetMetadata_ResourceMetadata(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ResourceMetadata
		expected *tfplugin5.GetMetadata_ResourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ResourceMetadata{},
			expected: &tfplugin5.GetMetadata_ResourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov5.ResourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin5.GetMetadata_ResourceMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_ResourceMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.GetMetadata_ResourceMetadata{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestImportResourceState_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ImportResourceStateResponse
		expected *tfplugin5.ImportResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ImportResourceStateResponse{},
			expected: &tfplugin5.ImportResourceState_Response{
				Diagnostics:       []*tfplugin5.Diagnostic{},
				ImportedResources: []*tfplugin5.ImportResourceState_ImportedResource{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ImportResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ImportResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				ImportedResources: []*tfplugin5.ImportResourceState_ImportedResource{},
			},
		},
		"ImportedResources": {
			in: &tfprotov5.ImportResourceStateResponse{
				ImportedResources: []*tfprotov5.ImportedResource{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin5.ImportResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				ImportedResources: []*tfplugin5.ImportResourceState_ImportedResource{
					{
						TypeName: "test",
					},
				},
			},
		},
		"Deferred": {
			in: &tfprotov5.ImportResourceStateResponse{
				Deferred: &tfprotov5.Deferred{
					Reason: tfprotov5.DeferredReasonResourceConfigUnknown,
				},
			},
			expected: &tfplugin5.ImportResourceState_Response{
				Diagnostics:       []*tfplugin5.Diagnostic{},
				ImportedResources: []*tfplugin5.ImportResourceState_ImportedResource{},
				Deferred: &tfplugin5.Deferred{
					Reason: tfplugin5.Deferred_RESOURCE_CONFIG_UNKNOWN,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ImportResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.ImportResourceState_ImportedResource{},
				tfplugin5.ImportResourceState_Response{},
				tfplugin5.Deferred{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestImportResourceState_ImportedResource(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ImportedResource
		expected *tfplugin5.ImportResourceState_ImportedResource
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ImportedResource{},
			expected: &tfplugin5.ImportResourceState_ImportedResource{},
		},
		"Private": {
			in: &tfprotov5.ImportedResource{
				Private: []byte("{}"),
			},
			expected: &tfplugin5.ImportResourceState_ImportedResource{
				Private: []byte("{}"),
			},
		},
		"State": {
			in: &tfprotov5.ImportedResource{
				State: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.ImportResourceState_ImportedResource{
				State: testTfplugin5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfprotov5.ImportedResource{
				TypeName: "test",
			},
			expected: &tfplugin5.ImportResourceState_ImportedResource{
				TypeName: "test",
			},
		},
		"Identity": {
			in: &tfprotov5.ImportedResource{
				Identity: testTfprotov5ResourceIdentityData(),
			},
			expected: &tfplugin5.ImportResourceState_ImportedResource{
				Identity: testTfplugin5ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ImportResourceState_ImportedResource(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.DynamicValue{},
				tfplugin5.ImportResourceState_ImportedResource{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestImportResourceState_ImportedResources(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tfprotov5.ImportedResource
		expected []*tfplugin5.ImportResourceState_ImportedResource
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin5.ImportResourceState_ImportedResource{},
		},
		"zero": {
			in:       []*tfprotov5.ImportedResource{},
			expected: []*tfplugin5.ImportResourceState_ImportedResource{},
		},
		"one": {
			in: []*tfprotov5.ImportedResource{
				{
					TypeName: "test",
				},
			},
			expected: []*tfplugin5.ImportResourceState_ImportedResource{
				{
					TypeName: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov5.ImportedResource{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
				},
			},
			expected: []*tfplugin5.ImportResourceState_ImportedResource{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ImportResourceState_ImportedResources(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ImportResourceState_ImportedResource{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestMoveResourceState_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.MoveResourceStateResponse
		expected *tfplugin5.MoveResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.MoveResourceStateResponse{},
			expected: &tfplugin5.MoveResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.MoveResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.MoveResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"TargetPrivate": {
			in: &tfprotov5.MoveResourceStateResponse{
				TargetPrivate: []byte(`{}`),
			},
			expected: &tfplugin5.MoveResourceState_Response{
				Diagnostics:   []*tfplugin5.Diagnostic{},
				TargetPrivate: []byte(`{}`),
			},
		},
		"TargetState": {
			in: &tfprotov5.MoveResourceStateResponse{
				TargetState: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.MoveResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				TargetState: testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.MoveResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.MoveResourceState_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanResourceChange_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.PlanResourceChangeResponse
		expected *tfplugin5.PlanResourceChange_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.PlanResourceChangeResponse{},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				RequiresReplace: []*tfplugin5.AttributePath{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.PlanResourceChangeResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
				RequiresReplace: []*tfplugin5.AttributePath{},
			},
		},
		"PlannedPrivate": {
			in: &tfprotov5.PlanResourceChangeResponse{
				PlannedPrivate: []byte("{}"),
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				PlannedPrivate:  []byte("{}"),
				RequiresReplace: []*tfplugin5.AttributePath{},
			},
		},
		"PlannedState": {
			in: &tfprotov5.PlanResourceChangeResponse{
				PlannedState: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				PlannedState:    testTfplugin5DynamicValue(),
				RequiresReplace: []*tfplugin5.AttributePath{},
			},
		},
		"RequiresReplace": {
			in: &tfprotov5.PlanResourceChangeResponse{
				RequiresReplace: []*tftypes.AttributePath{
					tftypes.NewAttributePath().WithAttributeName("test"),
				},
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				RequiresReplace: []*tfplugin5.AttributePath{
					{
						Steps: []*tfplugin5.AttributePath_Step{
							{
								Selector: &tfplugin5.AttributePath_Step_AttributeName{
									AttributeName: "test",
								},
							},
						},
					},
				},
			},
		},
		"UnsafeToUseLegacyTypeSystem": {
			in: &tfprotov5.PlanResourceChangeResponse{
				UnsafeToUseLegacyTypeSystem: true,
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:      []*tfplugin5.Diagnostic{},
				LegacyTypeSystem: true,
				RequiresReplace:  []*tfplugin5.AttributePath{},
			},
		},
		"Deferred": {
			in: &tfprotov5.PlanResourceChangeResponse{
				Deferred: &tfprotov5.Deferred{
					Reason: tfprotov5.DeferredReasonProviderConfigUnknown,
				},
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				RequiresReplace: []*tfplugin5.AttributePath{},
				Deferred: &tfplugin5.Deferred{
					Reason: tfplugin5.Deferred_PROVIDER_CONFIG_UNKNOWN,
				},
			},
		},
		"PlannedIdentity": {
			in: &tfprotov5.PlanResourceChangeResponse{
				PlannedIdentity: testTfprotov5ResourceIdentityData(),
			},
			expected: &tfplugin5.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin5.Diagnostic{},
				RequiresReplace: []*tfplugin5.AttributePath{},
				PlannedIdentity: testTfplugin5ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.PlanResourceChange_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.AttributePath{},
				tfplugin5.AttributePath_Step{},
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.PlanResourceChange_Response{},
				tfplugin5.Deferred{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadResource_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ReadResourceResponse
		expected *tfplugin5.ReadResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ReadResourceResponse{},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ReadResourceResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"NewState": {
			in: &tfprotov5.ReadResourceResponse{
				NewState: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				NewState:    testTfplugin5DynamicValue(),
			},
		},
		"Private": {
			in: &tfprotov5.ReadResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"Deferred": {
			in: &tfprotov5.ReadResourceResponse{
				Deferred: &tfprotov5.Deferred{
					Reason: tfprotov5.DeferredReasonAbsentPrereq,
				},
			},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				Deferred: &tfplugin5.Deferred{
					Reason: tfplugin5.Deferred_ABSENT_PREREQ,
				},
			},
		},
		"NewIdentity": {
			in: &tfprotov5.ReadResourceResponse{
				NewIdentity: testTfprotov5ResourceIdentityData(),
			},
			expected: &tfplugin5.ReadResource_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
				NewIdentity: testTfplugin5ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ReadResource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.ReadResource_Response{},
				tfplugin5.Deferred{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestUpgradeResourceState_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.UpgradeResourceStateResponse
		expected *tfplugin5.UpgradeResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.UpgradeResourceStateResponse{},
			expected: &tfplugin5.UpgradeResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.UpgradeResourceState_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"UpgradedState": {
			in: &tfprotov5.UpgradeResourceStateResponse{
				UpgradedState: testTfprotov5DynamicValue(),
			},
			expected: &tfplugin5.UpgradeResourceState_Response{
				Diagnostics:   []*tfplugin5.Diagnostic{},
				UpgradedState: testTfplugin5DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.UpgradeResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.UpgradeResourceState_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestUpgradeResourceIdentity_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.UpgradeResourceIdentityResponse
		expected *tfplugin5.UpgradeResourceIdentity_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.UpgradeResourceIdentityResponse{},
			expected: &tfplugin5.UpgradeResourceIdentity_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.UpgradeResourceIdentityResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.UpgradeResourceIdentity_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
		"UpgradedIdentity": {
			in: &tfprotov5.UpgradeResourceIdentityResponse{
				UpgradedIdentity: testTfprotov5ResourceIdentityData(),
			},
			expected: &tfplugin5.UpgradeResourceIdentity_Response{
				Diagnostics:      []*tfplugin5.Diagnostic{},
				UpgradedIdentity: testTfplugin5ResourceIdentityData(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.UpgradeResourceIdentity_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.DynamicValue{},
				tfplugin5.UpgradeResourceIdentity_Response{},
				tfplugin5.ResourceIdentityData{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateResourceTypeConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ValidateResourceTypeConfigResponse
		expected *tfplugin5.ValidateResourceTypeConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov5.ValidateResourceTypeConfigResponse{},
			expected: &tfplugin5.ValidateResourceTypeConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov5.ValidateResourceTypeConfigResponse{
				Diagnostics: []*tfprotov5.Diagnostic{
					testTfprotov5Diagnostic,
				},
			},
			expected: &tfplugin5.ValidateResourceTypeConfig_Response{
				Diagnostics: []*tfplugin5.Diagnostic{
					testTfplugin5Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ValidateResourceTypeConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.Diagnostic{},
				tfplugin5.ValidateResourceTypeConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
