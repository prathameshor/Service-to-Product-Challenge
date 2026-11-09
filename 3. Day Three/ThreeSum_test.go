package day_three

import "testing"

func TestThreeSum(t *testing.T) {
	test := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "Test case 1",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, 1, 0}, {-1, -1, 2}},
		},
	}

	for _, tc := range test {
		t.Run(tc.name, func(t *testing.T) {
			result := ThreeSum_BruteForce(tc.nums)
			if len(result) != len(tc.expected) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, result, tc.expected)
			}
		})

		t.Run(tc.name, func(t *testing.T) {
			result := ThreeSum_Better(tc.nums)
			if len(result) != len(tc.expected) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, result, tc.expected)
			}
		})

		t.Run(tc.name, func(t *testing.T) {
			result := ThreeSum_Optimal(tc.nums)
			if len(result) != len(tc.expected) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, result, tc.expected)
			}
		})
	}
}
