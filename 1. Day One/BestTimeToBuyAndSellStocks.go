package day_one

func BestTimeToBuyAndSellStocks(nums []int) int {
	profit, buying_price := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < buying_price {
			buying_price = nums[i]
		} else if nums[i]-buying_price > profit {
			profit = nums[i] - buying_price
		}
	}
	return profit
}
