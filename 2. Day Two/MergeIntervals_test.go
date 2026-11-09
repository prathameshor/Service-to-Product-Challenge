package day_two

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	testCases := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "Test case 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Test case 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeIntervals(tc.intervals)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("MergeIntervals(%v) = %v, want %v", tc.intervals, result, tc.expected)
			}
		})
	}
}
