// Copyright IBM Corp. 2020, 2026
// SPDX-License-Identifier: MPL-2.0

package tfprotov6_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

func TestDynamicValueIsNull(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		dynamicValue  tfprotov6.DynamicValue
		expected      bool
		expectedError error
	}{
		"empty-dynamic-value": {
			dynamicValue:  tfprotov6.DynamicValue{},
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

func testNewDynamicValueMust(t *testing.T, typ tftypes.Type, value tftypes.Value) tfprotov6.DynamicValue {
	t.Helper()

	dynamicValue, err := tfprotov6.NewDynamicValue(typ, value)

	if err != nil {
		t.Fatalf("unable to create DynamicValue: %s", err)
	}

	return dynamicValue
}
