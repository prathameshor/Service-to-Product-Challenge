package day_two

import (
	"reflect"
	"testing"
)

func TestMaximumSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		subArray []int
		want     int
	}{
		{
			name:     "Test case 1",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			subArray: []int{4, -1, 2, 1},
			want:     6,
		},
		{
			name:     "Test case 2",
			nums:     []int{1},
			subArray: []int{1},
			want:     1,
		},
		{
			name:     "Test case 3",
			nums:     []int{5, 4, -1, 7, 8},
			subArray: []int{5, 4, -1, 7, 8},
			want:     23,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := MaximumSumSubarray_BruteForce(test.nums)
			if result != test.want {
				t.Errorf("MaximumSumSubarray(%v) = %d, want %d", test.nums, result, test.want)
			}
		})

		t.Run(test.name, func(t *testing.T) {
			result := MaximumSumSubarray_Better(test.nums)
			if result != test.want {
				t.Errorf("MaximumSumSubarray(%v) = %d, want %d", test.nums, result, test.want)
			}
		})

		t.Run(test.name, func(t *testing.T) {
			subArray, sum := MaximumSumSubarray_Optimal(test.nums)
			if sum != test.want && !reflect.DeepEqual(subArray, test.subArray) {
				t.Errorf("MaximumSumSubarray(%v) = %d, want %d and %v", test.nums, sum, test.want, subArray)
			}
		})
	}
}
