package main

import "testing"

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
			input:    "red turtle   ",
			expected: []string{"red", "turtle"},
		},
		{
			input:    "  bluelizard  knight SWORD",
			expected: []string{"bluelizard", "knight", "sword"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("The length of actual slice %d does not match the length of the expected slice %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Actual work: %s does not match expected word: %s", word, expectedWord)
			}
		}
	}
}
