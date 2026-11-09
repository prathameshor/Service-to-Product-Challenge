package day_three

import (
	"reflect"
	"slices"
)

func NextPermutation_BruteForce(nums []int) []int {
	var pnums []int
	for _, num := range nums {
		pnums = append(pnums, num)
	}
	slices.Sort(pnums)
	permutations := GetPermutation(pnums)
	// slices.SortFunc(permutations, func(a, b []int) int {
	// 	for i := range a {
	// 		if a[i] < b[i] {
	// 			return -1
	// 		}
	// 		if a[i] > b[i] {
	// 			return 1
	// 		}
	// 	}
	// 	return 0
	// })
	for i, perm := range permutations {
		if reflect.DeepEqual(perm, nums) {
			return permutations[(i+1)%len(permutations)]
		}
	}
	return nil
}

func GetPermutation(nums []int) [][]int {
	var permutations [][]int
	recurringPermute(0, nums, &permutations)
	return permutations
}

func recurringPermute(start int, nums []int, permutations *[][]int) {
	if start == len(nums)-1 {
		perm := make([]int, len(nums))
		copy(perm, nums)
		*permutations = append(*permutations, perm)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		recurringPermute(start+1, nums, permutations)
		nums[start], nums[i] = nums[i], nums[start]
	}
}
