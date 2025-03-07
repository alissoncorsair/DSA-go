package main

import "testing"

func TestBinarySearch(t *testing.T) {
	testCases := []struct {
		name     string
		array    []int
		target   int
		expected int
	}{
		{"Empty array", []int{}, 5, -1},
		{"Single element found", []int{5}, 5, 0},
		{"Single element not found", []int{5}, 10, -1},
		{"First element", []int{1, 3, 5, 7, 9}, 1, 0},
		{"Last element", []int{1, 3, 5, 7, 9}, 9, 4},
		{"Middle element", []int{1, 3, 5, 7, 9}, 5, 2},
		{"Element not present", []int{1, 3, 5, 7, 9}, 4, -1},
		{"Larger array", []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}, 13, 6},
		{"Negative numbers", []int{-10, -5, 0, 5, 10}, -5, 1},
		{"All negative numbers", []int{-10, -8, -6, -4, -2}, -6, 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := binarySearch(tc.array, tc.target)
			if result != tc.expected {
				t.Errorf("binarySearch(%v, %d) = %d; expected %d",
					tc.array, tc.target, result, tc.expected)
			}
		})
	}
}
