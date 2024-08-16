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

func TestValidateEphemeralResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.ValidateEphemeralResourceConfig_Request
		expected *tfprotov6.ValidateEphemeralResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.ValidateEphemeralResourceConfig_Request{},
			expected: &tfprotov6.ValidateEphemeralResourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin6.ValidateEphemeralResourceConfig_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.ValidateEphemeralResourceConfigRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.ValidateEphemeralResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.ValidateEphemeralResourceConfigRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.ValidateEphemeralResourceConfigRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestOpenEphemeralResourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.OpenEphemeralResource_Request
		expected *tfprotov6.OpenEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.OpenEphemeralResource_Request{},
			expected: &tfprotov6.OpenEphemeralResourceRequest{},
		},
		"Config": {
			in: &tfplugin6.OpenEphemeralResource_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.OpenEphemeralResourceRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.OpenEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.OpenEphemeralResourceRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.OpenEphemeralResourceRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestRenewEphemeralResourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.RenewEphemeralResource_Request
		expected *tfprotov6.RenewEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.RenewEphemeralResource_Request{},
			expected: &tfprotov6.RenewEphemeralResourceRequest{},
		},
		"Config": {
			in: &tfplugin6.RenewEphemeralResource_Request{
				Config: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.RenewEphemeralResourceRequest{
				Config: testTfprotov6DynamicValue(),
			},
		},
		"Private": {
			in: &tfplugin6.RenewEphemeralResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov6.RenewEphemeralResourceRequest{
				Private: []byte("{}"),
			},
		},
		"PriorState": {
			in: &tfplugin6.RenewEphemeralResource_Request{
				PriorState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.RenewEphemeralResourceRequest{
				PriorState: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.RenewEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.RenewEphemeralResourceRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.RenewEphemeralResourceRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestCloseEphemeralResourceRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin6.CloseEphemeralResource_Request
		expected *tfprotov6.CloseEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin6.CloseEphemeralResource_Request{},
			expected: &tfprotov6.CloseEphemeralResourceRequest{},
		},
		"Private": {
			in: &tfplugin6.CloseEphemeralResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov6.CloseEphemeralResourceRequest{
				Private: []byte("{}"),
			},
		},
		"PriorState": {
			in: &tfplugin6.CloseEphemeralResource_Request{
				PriorState: testTfplugin6DynamicValue(),
			},
			expected: &tfprotov6.CloseEphemeralResourceRequest{
				PriorState: testTfprotov6DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin6.CloseEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov6.CloseEphemeralResourceRequest{
				TypeName: "test",
			},
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := fromproto.CloseEphemeralResourceRequest(testCase.in)

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
