package logging

import "testing"

func TestProviderLoggerName(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		providerAddress string
		expected        string
	}{
		"unparseable": {
			providerAddress: "example.com/test",
			expected:        "",
		},
		"basic": {
			providerAddress: "example.com/user/test",
			expected:        "test",
		},
		"hyphenated": {
			providerAddress: "example.com/user/test-test",
			expected:        "test_test",
		},
	}

	for name, testCase := range testCases {
		testCase := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := ProviderLoggerName(testCase.providerAddress)

			if testCase.expected != got {
				t.Errorf("wanted: %q, got: %q", testCase.expected, got)
			}
		})
	}
}
