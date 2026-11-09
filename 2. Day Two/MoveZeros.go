package day_two

func MoveZeros(nums []int) {
	n := len(nums)
	non_zero := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[non_zero], nums[i] = nums[i], nums[non_zero]
			non_zero++
		}
	}
}
