package day_two

//Complexity close to O(n^3)
func MaximumSumSubarray_BruteForce(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			sum := 0
			for k := i; k <= j; k++ {
				sum += nums[k]
			}
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

//Complexity close to O(n^2)
func MaximumSumSubarray_Better(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

//Complexity close to O(n) - Kaden's Algorithm
func MaximumSumSubarray_Optimal(nums []int) ([]int, int) {
	max, sum := 0, 0
	maxSumSubArray := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if (nums[i] + sum) > 0 {
			maxSumSubArray = append(maxSumSubArray, nums[i])
			sum += nums[i]
			if sum > max {
				max = sum
			}
		} else {
			sum = 0
			maxSumSubArray = []int{}
		}
	}
	return maxSumSubArray, max
}
