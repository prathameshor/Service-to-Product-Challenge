package day_two

import "testing"

func TestRotateArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Test case 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Test case 2",
			nums:     []int{1, 2, 3, 4, 5, 6},
			k:        2,
			expected: []int{5, 6, 1, 2, 3, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			beforeRotate := tc.nums
			tc.nums = RotateArray(tc.nums, tc.k)
			for i := 0; i < len(tc.nums); i++ {
				if tc.nums[i] != tc.expected[i] {
					t.Errorf("RotateArray(%v, %d) = %v, want %v", beforeRotate, tc.k, tc.nums, tc.expected)
				}
			}
		})
	}
}
