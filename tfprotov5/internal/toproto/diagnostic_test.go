package toproto

import (
	"testing"
)

func TestForceValidUTF8(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Input string
		Want  string
	}{
		{
			"hello",
			"hello",
		},
		{
			"こんにちは",
			"こんにちは",
		},
		{
			"baﬄe", // NOTE: "ﬄ" is a single-character ligature
			"baﬄe", // ligature is preserved exactly
		},
		{
			"wé́́é́́é́́!", // NOTE: These "e" have multiple combining diacritics
			"wé́́é́́é́́!", // diacritics are preserved exactly
		},
		{
			"😸😾", // Astral-plane characters
			"😸😾", // preserved exactly
		},
		{
			"\xff\xff",     // neither byte is valid UTF-8
			"\ufffd\ufffd", // both are replaced by replacement character
		},
		{
			"\xff\xff\xff\xff\xff",           // more than three invalid bytes
			"\ufffd\ufffd\ufffd\ufffd\ufffd", // still expanded even though it exceeds our initial slice capacity in the implementation
		},
		{
			"t\xffe\xffst",     // invalid bytes interleaved with other content
			"t\ufffde\ufffdst", // the valid content is preserved
		},
		{
			"\xffこんにちは\xffこんにちは",     // invalid bytes interacting with multibyte sequences
			"\ufffdこんにちは\ufffdこんにちは", // the valid content is preserved
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Input, func(t *testing.T) {
			t.Parallel()

			got := forceValidUTF8(test.Input)
			if got != test.Want {
				t.Errorf("wrong result\ngot:  %q\nwant: %q", got, test.Want)
			}
		})
	}
}
