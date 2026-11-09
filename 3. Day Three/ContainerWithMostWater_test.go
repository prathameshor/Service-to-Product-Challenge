package day_three

import "testing"

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test 1",
			nums:     []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Test 2",
			nums:     []int{1, 1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := ContainerWithMostWater(test.nums)
			if output != test.expected {
				t.Errorf("ContainerWithMostWater(%v) = %d, expected %d", test.nums, output, test.expected)
			}
		})
	}
}
