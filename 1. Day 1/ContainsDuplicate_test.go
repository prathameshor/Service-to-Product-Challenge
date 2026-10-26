package day_one

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Has Duplicates",
			nums: []int{1, 2, 0, 6, 7, 9, 1, 9},
			want: false,
		},
		{
			name: "Has Duplicates",
			nums: []int{1, 2, 0, 6, 7, 9},
			want: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ContainsDuplicate(test.nums)
			if result != test.want {
				t.Errorf("ContainsDuplicate(%v) = %t, wants %t", test.nums, result, test.want)
			}
		})
	}
}
