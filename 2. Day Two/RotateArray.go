package day_two

func RotateArray(nums []int, k int) []int {
	n := len(nums)
	for range k {
		arr := []int{nums[n-1]}
		nums = append(arr, nums[0:n-1]...)
	}
	return nums
}
