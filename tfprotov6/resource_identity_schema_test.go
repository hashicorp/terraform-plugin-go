// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package tfprotov6_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestResourceIdentitySchemaValueType(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		identitySchema *tfprotov6.ResourceIdentitySchema
		expected       tftypes.Type
	}{
		"nil": {
			identitySchema: nil,
			expected: tftypes.Object{
				AttributeTypes: map[string]tftypes.Type{},
			},
		},
		"Attribute-String": {
			identitySchema: &tfprotov6.ResourceIdentitySchema{
				IdentityAttributes: []*tfprotov6.ResourceIdentitySchemaAttribute{
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
		identitySchemaAttribute *tfprotov6.ResourceIdentitySchemaAttribute
		expected                tftypes.Type
	}{
		"nil": {
			identitySchemaAttribute: nil,
			expected:                nil,
		},
		"Bool": {
			identitySchemaAttribute: &tfprotov6.ResourceIdentitySchemaAttribute{
				Type: tftypes.Bool,
			},
			expected: tftypes.Bool,
		},
		"List-String": {
			identitySchemaAttribute: &tfprotov6.ResourceIdentitySchemaAttribute{
				Type: tftypes.List{
					ElementType: tftypes.String,
				},
			},
			expected: tftypes.List{
				ElementType: tftypes.String,
			},
		},
		"Number": {
			identitySchemaAttribute: &tfprotov6.ResourceIdentitySchemaAttribute{
				Type: tftypes.Number,
			},
			expected: tftypes.Number,
		},
		"String": {
			identitySchemaAttribute: &tfprotov6.ResourceIdentitySchemaAttribute{
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
