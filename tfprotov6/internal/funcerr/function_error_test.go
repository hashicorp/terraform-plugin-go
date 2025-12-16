// Copyright IBM Corp. 2020, 2025
// SPDX-License-Identifier: MPL-2.0

package funcerr_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
	"github.com/hashicorp/terraform-plugin-log/tfsdklogtest"

	"github.com/hashicorp/terraform-plugin-go/internal/logging"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6/internal/funcerr"
)

func TestFunctionErrorHasError(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		functionError *funcerr.FunctionError
		expected      bool
	}{
		"nil": {
			functionError: nil,
			expected:      false,
		},
		"empty": {
			functionError: &funcerr.FunctionError{},
			expected:      false,
		},
		"argument": {
			functionError: &funcerr.FunctionError{
				FunctionArgument: pointer(int64(0)),
			},
			expected: true,
		},
		"message": {
			functionError: &funcerr.FunctionError{
				Text: "test function error",
			},
			expected: true,
		},
		"argument-and-message": {
			functionError: &funcerr.FunctionError{
				FunctionArgument: pointer(int64(0)),
				Text:             "test function error",
			},
			expected: true,
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.functionError.HasError()

			if diff := cmp.Diff(got, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}

func TestFunctionErrorLog(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		functionError *funcerr.FunctionError
		expected      []map[string]interface{}
	}{
		"nil": {
			functionError: nil,
			expected:      nil,
		},
		"empty": {
			functionError: &funcerr.FunctionError{},
			expected:      nil,
		},
		"argument": {
			functionError: &funcerr.FunctionError{
				FunctionArgument: pointer(int64(0)),
			},
			expected: []map[string]interface{}{
				{
					"@level":                  "error",
					"@message":                "Response contains function error",
					"@module":                 "sdk.proto",
					"function_error_argument": float64(0),
				},
			},
		},
		"message": {
			functionError: &funcerr.FunctionError{
				Text: "test function error",
			},
			expected: []map[string]interface{}{
				{
					"@level":              "error",
					"@message":            "Response contains function error",
					"@module":             "sdk.proto",
					"function_error_text": "test function error",
				},
			},
		},
		"argument-and-message": {
			functionError: &funcerr.FunctionError{
				FunctionArgument: pointer(int64(0)),
				Text:             "test function error",
			},
			expected: []map[string]interface{}{
				{
					"@level":                  "error",
					"@message":                "Response contains function error",
					"@module":                 "sdk.proto",
					"function_error_text":     "test function error",
					"function_error_argument": float64(0),
				},
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var output bytes.Buffer

			ctx := tfsdklogtest.RootLogger(context.Background(), &output)
			ctx = logging.ProtoSubsystemContext(ctx, tfsdklog.Options{})

			testCase.functionError.Log(ctx)

			entries, err := tfsdklogtest.MultilineJSONDecode(&output)

			if err != nil {
				t.Fatalf("unable to read multiple line JSON: %s", err)
			}

			if diff := cmp.Diff(entries, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
