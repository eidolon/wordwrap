package wordwrap_test

import (
	"testing"

	"github.com/eidolon/wordwrap"
)

func TestWrapper(t *testing.T) {
	tests := []struct {
		limit      uint
		breakWords bool
		input      string
		expected   string
	}{
		{
			10,
			false,
			"Test text Test text Test text Test text",
			"Test text\nTest text\nTest text\nTest text",
		},
	}

	for _, test := range tests {
		wrapper := wordwrap.Wrapper(test.limit, test.breakWords)
		wrapped := wrapper(test.input)

		if wrapped != test.expected {
			t.Fatalf(
				"Wrapper did not return expected result.\nExpected:\n%s\nActual:\n%s\n",
				test.expected,
				wrapped,
			)
		}
	}
}