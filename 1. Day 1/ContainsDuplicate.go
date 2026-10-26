package day_one

func ContainsDuplicate(nums []int) bool {
	countNums := make(map[int]int)
	for _, num := range nums {
		if _, ok := countNums[num]; ok {
			return false
		}
		countNums[num] = 1
	}
	return true
}
