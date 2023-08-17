package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "example-1",
			input:    []int{1, 2, 3, 1},
			expected: true,
		},
		{
			name:     "example-2",
			input:    []int{1, 2, 3, 4},
			expected: false,
		},
		{
			name:     "example-3",
			input:    []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			expected: true,
		},
	}

	for _, el := range testCases {
		result := containsDuplicate(el.input)
		if result != el.expected {
			t.Errorf("%s: expected %t but got %t", el.name, el.expected, result)
		}
	}
}
