package wordwrap_test

import (
	"testing"

	"github.com/eidolon/wordwrap"
)

func TestWrapper(t *testing.T) {
	tests := []struct {
		limit      int
		breakWords bool
		input      string
		expected   string
	}{
		// When not breaking words
		// Should wrap text so words fit on lines with the given limit
		{
			10,
			false,
			"Test text Test text Test text Test text",
			"Test text\nTest text\nTest text\nTest text",
		},
		// Should remove additional whitespace
		{
			10,
			false,
			"Test  text  Test  text  Test  text  Test  text",
			"Test text\nTest text\nTest text\nTest text",
		},
		// Should trim text
		{
			10,
			false,
			"    Test  text  Test  text  Test  text  Test  text    ",
			"Test text\nTest text\nTest text\nTest text",
		},
		// Should not break words if breakWords is false
		{
			4,
			false,
			"Testtext",
			"Testtext",
		},
		// When breaking words
		// Should break words, and insert a hyphen
		{
			4,
			true,
			"Testtext",
			"Test\ntext",
		},
		// Should break words into tiny pieces
		{
			1,
			true,
			"Testtext",
			"T\ne\ns\nt\nt\ne\nx\nt",
		},
		// Slightly more realistic data
		// Should be broken properly
		{
			40,
			false,
			"This is a bunch of text that should be split at somewhere near 40 characters.",
			"This is a bunch of text that should be\nsplit at somewhere near 40 characters.",
		},
	}

	for i, test := range tests {
		wrapper := wordwrap.Wrapper(test.limit, test.breakWords)
		wrapped := wrapper(test.input)

		if wrapped != test.expected {
			t.Fatalf(
				"Wrapper for case %d did not return expected result.\nExpected:\n%s\nActual:\n%s\n",
				i+1,
				test.expected,
				wrapped,
			)
		}
	}
}

func TestWrapperPanicsWithInvalidLimit(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("Wrapper did not panic on invalid input.")
		}
	}()

	// 0 is an invalid limit
	wordwrap.Wrapper(0, false)
}
