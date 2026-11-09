package day_three

//O(n^2)
func MajorityElement_BruteForce(nums []int) int {
	for i := range nums {
		count := 0
		for j := range nums {
			if nums[i] == nums[j] {
				count++
			}
		}
		if count > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}

//O(N log N) + O(N)
func MajorityElement_Better(nums []int) int {
	numMap := make(map[int]int, 0)
	for _, num := range nums {
		numMap[num]++
	}
	for num, count := range numMap {
		if count > len(nums)/2 {
			return num
		}
	}
	return -1
}

//Moore's Voting Algorithm
//O(N) + O(N)
func MajorityElement_Optimal(nums []int) int {
	val, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			val = num
			count++
			continue
		} else if num != val {
			count--
		} else {
			count++
		}
	}
	count1 := 0
	for _, num := range nums {
		if num == val {
			count1++
		}
	}
	if count1 > len(nums)/2 {
		return val
	}
	return -1
}
