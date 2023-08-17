package main

import "testing"

func Test_removeDuplicates(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "example-1",
			input:    "abbaca",
			expected: "ca",
		},
		{
			name:     "example-2",
			input:    "azxxzy",
			expected: "ay",
		},
		{
			name:     "example-3",
			input:    "aaaaaaaaa",
			expected: "a",
		},
	}

	for _, el := range testCases {
		result := removeDuplicates(el.input)
		if result != el.expected {
			t.Errorf("%s: expected %s but got %s", el.name, el.expected, result)
		}
	}
}
