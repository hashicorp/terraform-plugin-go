// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfprotov5_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestDynamicValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		dynamicValue  tfprotov5.DynamicValue
		expected      bool
		expectedError error
	}{
		"empty-dynamic-value": {
			dynamicValue:  tfprotov5.DynamicValue{},
			expected:      false,
			expectedError: fmt.Errorf("unable to read DynamicValue: DynamicValue had no JSON or msgpack data set"),
		},
		"null": {
			dynamicValue: testNewDynamicValueMust(t,
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_string_attribute": tftypes.String,
					},
				},
				tftypes.NewValue(
					tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"test_string_attribute": tftypes.String,
						},
					},
					nil,
				),
			),
			expected: true,
		},
		"known": {
			dynamicValue: testNewDynamicValueMust(t,
				tftypes.Object{
					AttributeTypes: map[string]tftypes.Type{
						"test_string_attribute": tftypes.String,
					},
				},
				tftypes.NewValue(
					tftypes.Object{
						AttributeTypes: map[string]tftypes.Type{
							"test_string_attribute": tftypes.String,
						},
					},
					map[string]tftypes.Value{
						"test_string_attribute": tftypes.NewValue(tftypes.String, "test-value"),
					},
				),
			),
			expected: false,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := testCase.dynamicValue.IsNull()

			if err != nil {
				if testCase.expectedError == nil {
					t.Fatalf("wanted no error, got error: %s", err)
				}

				if !strings.Contains(err.Error(), testCase.expectedError.Error()) {
					t.Fatalf("wanted error %q, got error: %s", testCase.expectedError.Error(), err.Error())
				}
			}

			if err == nil && testCase.expectedError != nil {
				t.Fatalf("got no error, wanted err: %s", testCase.expectedError)
			}

			if got != testCase.expected {
				t.Errorf("expected %t, got %t", testCase.expected, got)
			}
		})
	}
}

func testNewDynamicValueMust(t *testing.T, typ tftypes.Type, value tftypes.Value) tfprotov5.DynamicValue {
	t.Helper()

	dynamicValue, err := tfprotov5.NewDynamicValue(typ, value)

	if err != nil {
		t.Fatalf("unable to create DynamicValue: %s", err)
	}

	return dynamicValue
}
