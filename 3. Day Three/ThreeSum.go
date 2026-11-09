package day_three

import (
	"fmt"
	"slices"
)

func ThreeSum_BruteForce(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					slice := []int{nums[i], nums[j], nums[k]}
					slices.Sort(slice)
					result = append(result, slice)
				}
			}
		}
	}

	return getUniqueSlices(result)
}

func ThreeSum_Better(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	for i := 0; i < n; i++ {
		hashset := make(map[int]struct{}, n)
		for j := i + 1; j < n; j++ {
			val := -(nums[i] + nums[j])
			if _, exists := hashset[val]; exists {
				triplet := []int{nums[i], nums[j], val}
				slices.Sort(triplet)
				result = append(result, triplet)
			}
			hashset[nums[j]] = struct{}{}
		}
	}

	return getUniqueSlices(result)
}

func ThreeSum_Optimal(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	slices.Sort(nums)
	fmt.Println(nums)
	for i := range n {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := n - 1
		for j < k {
			val := nums[i] + nums[j] + nums[k]
			if val < 0 {
				j++
			} else if val > 0 {
				k--
			} else {
				triplet := []int{nums[i], nums[j], nums[k]}
				result = append(result, triplet)
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}

	return getUniqueSlices(result)
}

func getUniqueSlices(input [][]int) [][]int {
	seen := make(map[string]struct{})
	var result [][]int

	for _, innerSlice := range input {
		key := fmt.Sprint(innerSlice)

		if _, exists := seen[key]; !exists {
			seen[key] = struct{}{}
			result = append(result, innerSlice)
		}
	}

	return result
}
