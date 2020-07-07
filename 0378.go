package leetcode_go

import (
	"sort"
)

func kthSmallest(matrix [][]int, k int) int {

	if k <= 0 || k > len(matrix)*len(matrix) {
		return 0
	}

	nums := make([]int, len(matrix)*len(matrix))

	for idx, v := range matrix {
		copy(nums[idx*len(matrix):], v)
	}

	sort.Ints(nums)

	return nums[k-1]
}
