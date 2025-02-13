// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
)

func TestResourceIdentitySchema(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       *tfprotov6.ResourceIdentitySchema
		expected *tfplugin6.ResourceIdentitySchema
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ResourceIdentitySchema{},
			expected: &tfplugin6.ResourceIdentitySchema{},
		},
		"IdentityAttributes": {
			in: &tfprotov6.ResourceIdentitySchema{
				IdentityAttributes: []*tfprotov6.ResourceIdentitySchemaAttribute{
					{
						Name: "test",
					},
				},
			},
			expected: &tfplugin6.ResourceIdentitySchema{
				IdentityAttributes: []*tfplugin6.ResourceIdentitySchema_IdentityAttribute{
					{
						Name: "test",
					},
				},
			},
		},
		"Version": {
			in: &tfprotov6.ResourceIdentitySchema{
				Version: 123,
			},
			expected: &tfplugin6.ResourceIdentitySchema{
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
				tfplugin6.ResourceIdentitySchema{},
				tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
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
		in       *tfprotov6.ResourceIdentitySchemaAttribute
		expected *tfplugin6.ResourceIdentitySchema_IdentityAttribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       &tfprotov6.ResourceIdentitySchemaAttribute{},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
		},
		"Name": {
			in: &tfprotov6.ResourceIdentitySchemaAttribute{
				Name: "test",
			},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{
				Name: "test",
			},
		},
		"Type": {
			in: &tfprotov6.ResourceIdentitySchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{
				Type: []byte(`"bool"`),
			},
		},
		"RequiredForImport": {
			in: &tfprotov6.ResourceIdentitySchemaAttribute{
				RequiredForImport: true,
			},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{
				RequiredForImport: true,
			},
		},
		"OptionalForImport": {
			in: &tfprotov6.ResourceIdentitySchemaAttribute{
				OptionalForImport: true,
			},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{
				OptionalForImport: true,
			},
		},
		"Description": {
			in: &tfprotov6.ResourceIdentitySchemaAttribute{
				Description: "test",
			},
			expected: &tfplugin6.ResourceIdentitySchema_IdentityAttribute{
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
				tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
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
		in       []*tfprotov6.ResourceIdentitySchemaAttribute
		expected []*tfplugin6.ResourceIdentitySchema_IdentityAttribute
	}{
		"nil": {
			in:       nil,
			expected: nil,
		},
		"zero": {
			in:       []*tfprotov6.ResourceIdentitySchemaAttribute{},
			expected: []*tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
		},
		"one": {
			in: []*tfprotov6.ResourceIdentitySchemaAttribute{
				{
					Name: "test",
				},
			},
			expected: []*tfplugin6.ResourceIdentitySchema_IdentityAttribute{
				{
					Name: "test",
				},
			},
		},
		"two": {
			in: []*tfprotov6.ResourceIdentitySchemaAttribute{
				{
					Name: "test1",
				},
				{
					Name: "test2",
				},
			},
			expected: []*tfplugin6.ResourceIdentitySchema_IdentityAttribute{
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
				tfplugin6.ResourceIdentitySchema_IdentityAttribute{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
