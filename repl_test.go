package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLength := len(actual)
		expectedLength := len(c.expected)

		if actualLength != expectedLength {
			t.Errorf("output length (%d) doesn't match expected (%d) for input '%s'", actualLength, expectedLength, c.input)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("word %s did not match expected %s", word, expectedWord)
			}
		}
	}
}
