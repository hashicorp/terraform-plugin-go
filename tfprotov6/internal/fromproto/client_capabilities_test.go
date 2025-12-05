// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestValidateResourceConfigClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ValidateResourceConfigClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ValidateResourceConfigClientCapabilities{},
		},
		"WriteOnlyAttributesAllowed": {
			in: &tfplugin6.ClientCapabilities{
				WriteOnlyAttributesAllowed: true,
			},
			expected: &tfprotov6.ValidateResourceConfigClientCapabilities{
				WriteOnlyAttributesAllowed: true,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateResourceConfigClientCapabilities(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestConfigureProviderClientCapabilities(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ConfigureProviderClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ConfigureProviderClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.ConfigureProviderClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ReadDataSourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ReadDataSourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.ReadDataSourceClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ReadResourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ReadResourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.ReadResourceClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.PlanResourceChangeClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.PlanResourceChangeClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.PlanResourceChangeClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.ImportResourceStateClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.ImportResourceStateClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.ImportResourceStateClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.OpenEphemeralResourceClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.OpenEphemeralResourceClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.OpenEphemeralResourceClientCapabilities{
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
		in       *tfplugin6.ClientCapabilities
		expected *tfprotov6.PlanActionClientCapabilities
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ClientCapabilities{},
			expected: &tfprotov6.PlanActionClientCapabilities{},
		},
		"DeferralAllowed": {
			in: &tfplugin6.ClientCapabilities{
				DeferralAllowed: true,
			},
			expected: &tfprotov6.PlanActionClientCapabilities{
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
