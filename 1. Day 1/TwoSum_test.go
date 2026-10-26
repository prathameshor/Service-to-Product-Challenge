package day_one

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "pair found",
			nums:   []int{1, 2, 0, 6, 7, 9},
			target: 6,
			want:   []int{2, 3},
		},
		{
			name:   "no pair found",
			nums:   []int{1, 2, 3},
			target: 10,
			want:   []int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := TwoSum(test.nums, test.target)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", test.nums, test.target, got, test.want)
			}
		})
	}
}
