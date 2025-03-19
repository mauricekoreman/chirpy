package main

import (
	"strings"
	"testing"
)

func TestReplaceProfanity(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    "something something kerfuffle joejoe",
			expected: "something something **** joejoe",
		},
		{
			input:    "something something kerfuffle kerfuffle",
			expected: "something something **** ****",
		},
		{
			input:    "something something kerfuffle sharbert",
			expected: "something something **** ****",
		},
		{
			input:    "something something kerfuffle Sharbert",
			expected: "something something **** ****",
		},
		{
			input:    "something something kerfuffle Sharbert!",
			expected: "something something **** Sharbert!",
		},
		{
			input:    "fornax something kerfuffle Sharbert!",
			expected: "**** something **** Sharbert!",
		},
	}

	for _, c := range cases {
		actual := clearProfanity(c.input)

		if strings.Compare(actual, c.expected) != 0 {
			t.Errorf("Expected: %s, but got: %s", c.expected, actual)
		}
	}
}
