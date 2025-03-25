// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestResourceIdentitySchemaValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identitySchema *tfprotov5.ResourceIdentitySchema
		expected       tftypes.Type
	}{
		"nil": {
			identitySchema: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			identitySchema: &tfprotov5.ResourceIdentitySchema{
				IdentityAttributes: []*tfprotov5.ResourceIdentitySchemaAttribute{
					{
						Name: "test_string_attribute",
						Type: tftypes.String,
					},
				},
			},
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{
					"test_string_attribute": tftypes.String,
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identitySchema.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}

func TestResourceIdentitySchemaAttributeValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identitySchemaAttribute *tfprotov5.ResourceIdentitySchemaAttribute
		expected                tftypes.Type
	}{
		"nil": {
			identitySchemaAttribute: nil,
			expected:                nil,
		},
		"Bool": {
			identitySchemaAttribute: &tfprotov5.ResourceIdentitySchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: tftypes.Bool,
		},
		"List-String": {
			identitySchemaAttribute: &tfprotov5.ResourceIdentitySchemaAttribute{
				Type: tftypes.List{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.String,
			},
		},
		"Number": {
			identitySchemaAttribute: &tfprotov5.ResourceIdentitySchemaAttribute{
				Type: tftypes.Number,
			},
			expected: tftypes.Number,
		},
		"String": {
			identitySchemaAttribute: &tfprotov5.ResourceIdentitySchemaAttribute{
				Type: tftypes.String,
			},
			expected: tftypes.String,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.identitySchemaAttribute.ValueType()

			if testCase.expected == nil {
				if got == nil {
					return
				}

				t.Fatalf("expected nil, got: %s", got)
			}

			if !testCase.expected.Equal(got) {
				t.Errorf("expected %s, got: %s", testCase.expected, got)
			}
		})
	}
}
