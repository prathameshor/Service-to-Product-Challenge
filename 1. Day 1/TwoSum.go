package day_one

func TwoSum(nums []int, target int) []int {
	index_map := make(map[int]int, 0)
	for index, num := range nums {
		look_up := target - num
		if value, ok := index_map[look_up]; ok {
			return []int{value, index}
		}
		index_map[num] = index
	}

	return []int{}
}

// func main() {
// 	fmt.Println(TwoSum([]int{1, 2, 0, 6, 7, 9}, 6))
// }
