package day_one

import (
	"reflect"
	"testing"
)

func TestProductOfArrayExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Valid Product",
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "Valid Product",
			nums: []int{0, 2, 3, 4},
			want: []int{24, 0, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			res := ProductOfArrayExceptSelf(test.nums)
			if !reflect.DeepEqual(res, test.want) {
				t.Errorf("ProductOfArrayExceptSelf(%v) = %v, expected %v", test.nums, res, test.want)
			}
		})
	}
}
