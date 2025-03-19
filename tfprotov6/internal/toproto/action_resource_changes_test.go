// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package toproto_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/tfplugin6"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/toproto"
)

func testTfprotov6ActionResourceChanges() map[string]*tfprotov6.DynamicValue {
	return map[string]*tfprotov6.DynamicValue{
		"cty.GetAttrPath(\"object\")": testTfprotov6DynamicValue(),
	}
}

func testTfplugin6ActionResourceChanges() map[string]*tfplugin6.DynamicValue {
	return toproto.ActionResourceChanges(testTfprotov6ActionResourceChanges())
}

func TestActionResourceChanges(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		in       map[string]*tfprotov6.DynamicValue
		expected map[string]*tfplugin6.DynamicValue
	}{
		"nil": {
			in:       nil,
			expected: map[string]*tfplugin6.DynamicValue{},
		},
		"zero": {
			in:       map[string]*tfprotov6.DynamicValue{},
			expected: map[string]*tfplugin6.DynamicValue{},
		},
		"one": {
			in: map[string]*tfprotov6.DynamicValue{
				"test": testTfprotov6DynamicValue(),
			},
			expected: map[string]*tfplugin6.DynamicValue{
				"test": testTfplugin6DynamicValue(),
			},
		},
		"two": {
			in: map[string]*tfprotov6.DynamicValue{
				"testA": testTfprotov6DynamicValue(),
				"testB": testTfprotov6DynamicValue(),
			},
			expected: map[string]*tfplugin6.DynamicValue{
				"testA": testTfplugin6DynamicValue(),
				"testB": testTfplugin6DynamicValue(),
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := toproto.ActionResourceChanges(testCase.in)

			// Protocol Buffers generated types must have unexported fields
			// ignored or cmp.Diff() will raise an error. This is easier than
			// writing a custom Comparer for each type, which would have no
			// benefits.
			diffOpts := cmpopts.IgnoreUnexported(
				tfplugin6.DynamicValue{},
			)

			if diff := cmp.Diff(got, testCase.expected, diffOpts); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
