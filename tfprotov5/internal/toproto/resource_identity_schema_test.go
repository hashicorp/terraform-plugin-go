// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/internal/tfplugin5"
)

func TestResourceIdentitySchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ResourceIdentitySchema
		expected *tfplugin5.ResourceIdentitySchema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ResourceIdentitySchema{},
			expected: &tfplugin5.ResourceIdentitySchema{},
		},
		"IdentityAttributes": {
			in: &tfprotov5.ResourceIdentitySchema{
				IdentityAttributes: []*tfprotov5.ResourceIdentitySchemaAttribute{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin5.ResourceIdentitySchema{
				IdentityAttributes: []*tfplugin5.ResourceIdentitySchema_IdentityAttribute{
					{
						Name: "test",
					},
				},
			},
		},
		"Version": {
			in: &tfprotov5.ResourceIdentitySchema{
				Version: 123,
			},
			expected: &tfplugin5.ResourceIdentitySchema{
				Version: 123,
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ResourceIdentitySchema(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ResourceIdentitySchema{},
				tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestResourceIdentitySchema_IdentityAttribute(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov5.ResourceIdentitySchemaAttribute
		expected *tfplugin5.ResourceIdentitySchema_IdentityAttribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov5.ResourceIdentitySchemaAttribute{},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
		},
		"Name": {
			in: &tfprotov5.ResourceIdentitySchemaAttribute{
				Name: "test",
			},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				Name: "test",
			},
		},
		"Type": {
			in: &tfprotov5.ResourceIdentitySchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				Type: []byte(`"bool"`),
			},
		},
		"RequiredForImport": {
			in: &tfprotov5.ResourceIdentitySchemaAttribute{
				RequiredForImport: true,
			},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				RequiredForImport: true,
			},
		},
		"OptionalForImport": {
			in: &tfprotov5.ResourceIdentitySchemaAttribute{
				OptionalForImport: true,
			},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				OptionalForImport: true,
			},
		},
		"Description": {
			in: &tfprotov5.ResourceIdentitySchemaAttribute{
				Description: "test",
			},
			expected: &tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				Description: "test",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ResourceIdentitySchema_IdentityAttribute(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestResourceIdentitySchema_IdentityAttributes(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       []*tfprotov5.ResourceIdentitySchemaAttribute
		expected []*tfplugin5.ResourceIdentitySchema_IdentityAttribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       []*tfprotov5.ResourceIdentitySchemaAttribute{},
			expected: []*tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
		},
		"one": {
			in: []*tfprotov5.ResourceIdentitySchemaAttribute{
				{
					Name: "test",
				},
			},
			expected: []*tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				{
					Name: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov5.ResourceIdentitySchemaAttribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
			expected: []*tfplugin5.ResourceIdentitySchema_IdentityAttribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ResourceIdentitySchema_IdentityAttributes(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin5.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
