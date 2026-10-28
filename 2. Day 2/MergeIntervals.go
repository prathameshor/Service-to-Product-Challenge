package day_two

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	for _, interval := range intervals {
		if len(result) == 0 || result[len(result)-1][1] < interval[0] {
			result = append(result, interval)
		} else {
			result[len(result)-1][1] = interval[1]
		}
	}

	return result
}
