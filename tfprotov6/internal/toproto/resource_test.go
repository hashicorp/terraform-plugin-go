// Copyright (c) HashiCorp, Inc.

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestApplyResourceChange_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ApplyResourceChangeResponse
		expected *tfplugin6.ApplyResourceChange_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ApplyResourceChangeResponse{},
			expected: &tfplugin6.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ApplyResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"Private": {
			in: &tfprotov6.ApplyResourceChangeResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin6.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
		"NewState": {
			in: &tfprotov6.ApplyResourceChangeResponse{
				NewState: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.ApplyResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				NewState:    testTfplugin6DynamicValue(),
			},
		},
		"UnsafeToUseLegacyTypeSystem": {
			in: &tfprotov6.ApplyResourceChangeResponse{
				UnsafeToUseLegacyTypeSystem: true,
			},
			expected: &tfplugin6.ApplyResourceChange_Response{
				Diagnostics:      []*tfplugin6.Diagnostic{},
				LegacyTypeSystem: true,
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ApplyResourceChange_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.ApplyResourceChange_Response{},
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
		in       *tfprotov6.ResourceMetadata
		expected *tfplugin6.GetMetadata_ResourceMetadata
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ResourceMetadata{},
			expected: &tfplugin6.GetMetadata_ResourceMetadata{},
		},
		"TypeName": {
			in: &tfprotov6.ResourceMetadata{
				TypeName: "test",
			},
			expected: &tfplugin6.GetMetadata_ResourceMetadata{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.GetMetadata_ResourceMetadata(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.GetMetadata_ResourceMetadata{},
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
		in       *tfprotov6.ImportResourceStateResponse
		expected *tfplugin6.ImportResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ImportResourceStateResponse{},
			expected: &tfplugin6.ImportResourceState_Response{
				Diagnostics:       []*tfplugin6.Diagnostic{},
				ImportedResources: []*tfplugin6.ImportResourceState_ImportedResource{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ImportResourceStateResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ImportResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				ImportedResources: []*tfplugin6.ImportResourceState_ImportedResource{},
			},
		},
		"ImportedResources": {
			in: &tfprotov6.ImportResourceStateResponse{
				ImportedResources: []*tfprotov6.ImportedResource{
					{
						TypeName: "test",
					},
				},
			},
			expected: &tfplugin6.ImportResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				ImportedResources: []*tfplugin6.ImportResourceState_ImportedResource{
					{
						TypeName: "test",
					},
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ImportResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ImportResourceState_ImportedResource{},
				tfplugin6.ImportResourceState_Response{},
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
		in       *tfprotov6.ImportedResource
		expected *tfplugin6.ImportResourceState_ImportedResource
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ImportedResource{},
			expected: &tfplugin6.ImportResourceState_ImportedResource{},
		},
		"Private": {
			in: &tfprotov6.ImportedResource{
				Private: []byte("{}"),
			},
			expected: &tfplugin6.ImportResourceState_ImportedResource{
				Private: []byte("{}"),
			},
		},
		"State": {
			in: &tfprotov6.ImportedResource{
				State: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.ImportResourceState_ImportedResource{
				State: testTfplugin6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfprotov6.ImportedResource{
				TypeName: "test",
			},
			expected: &tfplugin6.ImportResourceState_ImportedResource{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ImportResourceState_ImportedResource(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.DynamicValue{},
				tfplugin6.ImportResourceState_ImportedResource{},
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
		in       []*tfprotov6.ImportedResource
		expected []*tfplugin6.ImportResourceState_ImportedResource
	}{
		"nil": {
			in:       nil,
			expected: []*tfplugin6.ImportResourceState_ImportedResource{},
		},
		"zero": {
			in:       []*tfprotov6.ImportedResource{},
			expected: []*tfplugin6.ImportResourceState_ImportedResource{},
		},
		"one": {
			in: []*tfprotov6.ImportedResource{
				{
					TypeName: "test",
				},
			},
			expected: []*tfplugin6.ImportResourceState_ImportedResource{
				{
					TypeName: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov6.ImportedResource{
				{
					TypeName: "test1",
				},
				{
					TypeName: "test2",
				},
			},
			expected: []*tfplugin6.ImportResourceState_ImportedResource{
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
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ImportResourceState_ImportedResources(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.ImportResourceState_ImportedResource{},
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
		in       *tfprotov6.MoveResourceStateResponse
		expected *tfplugin6.MoveResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.MoveResourceStateResponse{},
			expected: &tfplugin6.MoveResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.MoveResourceStateResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.MoveResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"TargetState": {
			in: &tfprotov6.MoveResourceStateResponse{
				TargetState: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.MoveResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				TargetState: testTfplugin6DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.MoveResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.MoveResourceState_Response{},
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
		in       *tfprotov6.PlanResourceChangeResponse
		expected *tfplugin6.PlanResourceChange_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.PlanResourceChangeResponse{},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin6.Diagnostic{},
				RequiresReplace: []*tfplugin6.AttributePath{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.PlanResourceChangeResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
				RequiresReplace: []*tfplugin6.AttributePath{},
			},
		},
		"PlannedPrivate": {
			in: &tfprotov6.PlanResourceChangeResponse{
				PlannedPrivate: []byte("{}"),
			},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin6.Diagnostic{},
				PlannedPrivate:  []byte("{}"),
				RequiresReplace: []*tfplugin6.AttributePath{},
			},
		},
		"PlannedState": {
			in: &tfprotov6.PlanResourceChangeResponse{
				PlannedState: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics:     []*tfplugin6.Diagnostic{},
				PlannedState:    testTfplugin6DynamicValue(),
				RequiresReplace: []*tfplugin6.AttributePath{},
			},
		},
		"RequiresReplace": {
			in: &tfprotov6.PlanResourceChangeResponse{
				RequiresReplace: []*tftypes.AttributePath{
					tftypes.NewAttributePath().WithAttributeName("test"),
				},
			},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				RequiresReplace: []*tfplugin6.AttributePath{
					{
						Steps: []*tfplugin6.AttributePath_Step{
							{
								Selector: &tfplugin6.AttributePath_Step_AttributeName{
									AttributeName: "test",
								},
							},
						},
					},
				},
			},
		},
		"UnsafeToUseLegacyTypeSystem": {
			in: &tfprotov6.PlanResourceChangeResponse{
				UnsafeToUseLegacyTypeSystem: true,
			},
			expected: &tfplugin6.PlanResourceChange_Response{
				Diagnostics:      []*tfplugin6.Diagnostic{},
				LegacyTypeSystem: true,
				RequiresReplace:  []*tfplugin6.AttributePath{},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.PlanResourceChange_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.AttributePath{},
				tfplugin6.AttributePath_Step{},
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.PlanResourceChange_Response{},
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
		in       *tfprotov6.ReadResourceResponse
		expected *tfplugin6.ReadResource_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ReadResourceResponse{},
			expected: &tfplugin6.ReadResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ReadResourceResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ReadResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"NewState": {
			in: &tfprotov6.ReadResourceResponse{
				NewState: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.ReadResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				NewState:    testTfplugin6DynamicValue(),
			},
		},
		"Private": {
			in: &tfprotov6.ReadResourceResponse{
				Private: []byte("{}"),
			},
			expected: &tfplugin6.ReadResource_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
				Private:     []byte("{}"),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ReadResource_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.ReadResource_Response{},
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
		in       *tfprotov6.UpgradeResourceStateResponse
		expected *tfplugin6.UpgradeResourceState_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.UpgradeResourceStateResponse{},
			expected: &tfplugin6.UpgradeResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.UpgradeResourceStateResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.UpgradeResourceState_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
		"UpgradedState": {
			in: &tfprotov6.UpgradeResourceStateResponse{
				UpgradedState: testTfprotov6DynamicValue(),
			},
			expected: &tfplugin6.UpgradeResourceState_Response{
				Diagnostics:   []*tfplugin6.Diagnostic{},
				UpgradedState: testTfplugin6DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.UpgradeResourceState_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.DynamicValue{},
				tfplugin6.UpgradeResourceState_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestValidateResourceConfig_Response(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ValidateResourceConfigResponse
		expected *tfplugin6.ValidateResourceConfig_Response
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in: &tfprotov6.ValidateResourceConfigResponse{},
			expected: &tfplugin6.ValidateResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{},
			},
		},
		"Diagnostics": {
			in: &tfprotov6.ValidateResourceConfigResponse{
				Diagnostics: []*tfprotov6.Diagnostic{
					testTfprotov6Diagnostic,
				},
			},
			expected: &tfplugin6.ValidateResourceConfig_Response{
				Diagnostics: []*tfplugin6.Diagnostic{
					testTfplugin6Diagnostic,
				},
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Intentionally not checking the error return as it is impossible
			// to implement a test case which would raise an error. This return
			// will be removed in preference of a panic a future change.
			// Reference: https://github.com/hashicorp/terraform-plugin-go/issues/365
			got, _ := toproto.ValidateResourceConfig_Response(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.Diagnostic{},
				tfplugin6.ValidateResourceConfig_Response{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
