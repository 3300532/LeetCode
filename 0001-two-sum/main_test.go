package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	testCases := []struct {
		name      string
		nums      []int
		target    int
		expected1 []int
		expected2 []int
	}{
		{
			name:      "example-1",
			nums:      []int{2, 7, 11, 15},
			target:    9,
			expected1: []int{0, 1},
			expected2: []int{1, 0},
		},
		{
			name:      "example-2",
			nums:      []int{3, 2, 4},
			target:    6,
			expected1: []int{1, 2},
			expected2: []int{2, 1},
		},
		{
			name:      "example-3",
			nums:      []int{3, 3},
			target:    6,
			expected1: []int{0, 1},
			expected2: []int{1, 0},
		},
	}

	for _, el := range testCases {
		result := twoSum(el.nums, el.target)

		if !reflect.DeepEqual(result, el.expected1) && !reflect.DeepEqual(result, el.expected2) {
			t.Errorf("%s: expected %v but got %v", el.name, el.expected1, result)
		}

	}
}
