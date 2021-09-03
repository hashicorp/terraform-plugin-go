package tftypes

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAttributePathErrorEqual(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ape      AttributePathError
		other    AttributePathError
		expected bool
	}{
		"equal": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			expected: true,
		},
		"err-different": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("different error"),
			},
			expected: false,
		},
		"err-nil": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  nil,
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  nil,
			},
			expected: true,
		},
		"err-nil-ape": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  nil,
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			expected: false,
		},
		"err-nil-other": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  nil,
			},
			expected: false,
		},
		"path-different": {
			ape: AttributePathError{
				Path: NewAttributePath().WithAttributeName("test"),
				err:  errors.New("test error"),
			},
			other: AttributePathError{
				Path: NewAttributePath().WithAttributeName("different"),
				err:  errors.New("test error"),
			},
			expected: false,
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := testCase.ape.Equal(testCase.other)

			if diff := cmp.Diff(testCase.expected, got); diff != "" {
				t.Errorf("Unexpected results (-wanted +got): %s", diff)
			}
		})
	}
}
