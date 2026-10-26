package day_one

import "testing"

func TestBestTimeToBuyAndSellStocks(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Best time",
			nums: []int{1, 2, 6, 7, 9, 1, 9},
			want: 8,
		},
		{
			name: "Best  time",
			nums: []int{1, 2, 6, 7},
			want: 6,
		},
		{
			name: "Best  time",
			nums: []int{7, 1, 5, 3, 6, 4},
			want: 5,
		},
		{
			name: "Best  time",
			nums: []int{7, 6, 4, 3, 1},
			want: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BestTimeToBuyAndSellStocks(test.nums)
			if result != test.want {
				t.Errorf("BestTimeTOBuyAndSellStocks(%v) = %d, wants %d", test.nums, result, test.want)
			}
		})
	}
}
