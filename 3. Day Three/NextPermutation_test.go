package day_three

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test 1",
			nums:     []int{1, 2, 3},
			expected: []int{1, 3, 2},
		},
		// {
		// 	name:     "Test 2",
		// 	nums:     []int{3, 2, 1},
		// 	expected: []int{1, 2, 3},
		// },
		{
			name:     "Test 3",
			nums:     []int{1, 1, 5},
			expected: []int{1, 5, 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := NextPermutation_BruteForce(test.nums)
			if !reflect.DeepEqual(output, test.expected) {
				t.Errorf("NextPermutation_BruteForce(%v) = %v, expected %v", test.nums, output, test.expected)
			}
		})
	}
}
