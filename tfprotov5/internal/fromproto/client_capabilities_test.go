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

func TestValidateResourceTypeConfigClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.ValidateResourceTypeConfigClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.ValidateResourceTypeConfigClientCapabilities{},
		},
		"WriteOnlyAttributesAllowed": {
			in: &tfplugin5.ClientCapabilities{
				WriteOnlyAttributesAllowed: true,
			},
			expected: &tfprotov5.ValidateResourceTypeConfigClientCapabilities{
				WriteOnlyAttributesAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateResourceTypeConfigClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestConfigureProviderClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.ConfigureProviderClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.ConfigureProviderClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.ConfigureProviderClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ConfigureProviderClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadDataSourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.ReadDataSourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.ReadDataSourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.ReadDataSourceClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ReadDataSourceClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestReadResourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.ReadResourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.ReadResourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.ReadResourceClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ReadResourceClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanResourceChangeClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.PlanResourceChangeClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.PlanResourceChangeClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.PlanResourceChangeClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.PlanResourceChangeClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestImportResourceStateClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.ImportResourceStateClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.ImportResourceStateClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.ImportResourceStateClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ImportResourceStateClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestOpenEphemeralResourceClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.OpenEphemeralResourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.OpenEphemeralResourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.OpenEphemeralResourceClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.OpenEphemeralResourceClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestPlanActionClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ClientCapabilities
		expected *tfprotov5.PlanActionClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ClientCapabilities{},
			expected: &tfprotov5.PlanActionClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin5.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov5.PlanActionClientCapabilities{
				DeferralAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.PlanActionClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
