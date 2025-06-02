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
			input:    "pOtaTo POTatO PoTatO",
			expected: []string{"potato", "potato", "potato"},
		},
		{
			input:    "       duck      horse      bratwurst",
			expected: []string{"duck", "horse", "bratwurst"},
		},
	}

	for _, c := range cases {
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to both print an error message
		// and fail the test
		actual := cleanInput(c.input)
		expected := c.expected
		if len(actual) != len(expected) {
			t.Errorf("actual length: %v | expected length: %v", len(actual), len(expected))
			continue
		}

		for i := range actual {
			// Check each word in the slice
			// if they don't match, use t.Errorf to both print an error message
			// and fail the test
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("actual word: %v | expected word: %v", word, expectedWord)
			}
		}
	}
}
