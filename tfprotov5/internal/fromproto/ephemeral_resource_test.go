// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package fromproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/fromproto"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestValidateEphemeralResourceConfigRequest(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfplugin5.ValidateEphemeralResourceConfig_Request
		expected *tfprotov5.ValidateEphemeralResourceConfigRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.ValidateEphemeralResourceConfig_Request{},
			expected: &tfprotov5.ValidateEphemeralResourceConfigRequest{},
		},
		"Config": {
			in: &tfplugin5.ValidateEphemeralResourceConfig_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.ValidateEphemeralResourceConfigRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.ValidateEphemeralResourceConfig_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.ValidateEphemeralResourceConfigRequest{
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
		in       *tfplugin5.OpenEphemeralResource_Request
		expected *tfprotov5.OpenEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.OpenEphemeralResource_Request{},
			expected: &tfprotov5.OpenEphemeralResourceRequest{},
		},
		"Config": {
			in: &tfplugin5.OpenEphemeralResource_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.OpenEphemeralResourceRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.OpenEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.OpenEphemeralResourceRequest{
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
		in       *tfplugin5.RenewEphemeralResource_Request
		expected *tfprotov5.RenewEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.RenewEphemeralResource_Request{},
			expected: &tfprotov5.RenewEphemeralResourceRequest{},
		},
		"Config": {
			in: &tfplugin5.RenewEphemeralResource_Request{
				Config: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.RenewEphemeralResourceRequest{
				Config: testTfprotov5DynamicValue(),
			},
		},
		"Private": {
			in: &tfplugin5.RenewEphemeralResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov5.RenewEphemeralResourceRequest{
				Private: []byte("{}"),
			},
		},
		"State": {
			in: &tfplugin5.RenewEphemeralResource_Request{
				State: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.RenewEphemeralResourceRequest{
				State: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.RenewEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.RenewEphemeralResourceRequest{
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
		in       *tfplugin5.CloseEphemeralResource_Request
		expected *tfprotov5.CloseEphemeralResourceRequest
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfplugin5.CloseEphemeralResource_Request{},
			expected: &tfprotov5.CloseEphemeralResourceRequest{},
		},
		"Private": {
			in: &tfplugin5.CloseEphemeralResource_Request{
				Private: []byte("{}"),
			},
			expected: &tfprotov5.CloseEphemeralResourceRequest{
				Private: []byte("{}"),
			},
		},
		"State": {
			in: &tfplugin5.CloseEphemeralResource_Request{
				State: testTfplugin5DynamicValue(),
			},
			expected: &tfprotov5.CloseEphemeralResourceRequest{
				State: testTfprotov5DynamicValue(),
			},
		},
		"TypeName": {
			in: &tfplugin5.CloseEphemeralResource_Request{
				TypeName: "test",
			},
			expected: &tfprotov5.CloseEphemeralResourceRequest{
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