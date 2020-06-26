package leetcode_go

import "sort"

func merge(intervals [][]int) [][]int {

	result := make([][]int, 0)

	li := len(intervals)
	if li < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < li; i++ {
		if intervals[i][0] <= end {
			end = maxM(end, intervals[i][1])
		} else {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}

	result = append(result, []int{start, end})

	return result
}

func maxM(x, y int) int {
	if x < y {
		return y
	}
	return x
}
