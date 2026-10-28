package day_two

import "testing"

func TestMoveZeros(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test case 1",
			nums:     []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			name:     "Test case 2",
			nums:     []int{0, 0, 1},
			expected: []int{1, 0, 0},
		},
		{
			name:     "Test case 3",
			nums:     []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			beforeMove := make([]int, len(tc.nums))
			copy(beforeMove, tc.nums)
			MoveZeros(tc.nums)
			for i := 0; i < len(tc.nums); i++ {
				if tc.nums[i] != tc.expected[i] {
					t.Errorf("MoveZeros(%v) = %v, want %v", beforeMove, tc.nums, tc.expected)
				}
			}
		})
	}
}
