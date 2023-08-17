package main

import "testing"

func Test_isAnagram(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "example-1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "example-2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for _, el := range testCases {
		result := isAnagram(el.s, el.t)
		if result != el.expected {
			t.Errorf("%s: expected %t but got %t", el.name, el.expected, result)
		}
	}
}
