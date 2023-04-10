package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func minPathSum(grid [][]int) int {

	m := len(grid)

	if m == 0 {
		return 0
	}

	n := len(grid[0])

	if n == 0 {
		return 0
	}

	left := make([]int, m)
	left[0] = grid[0][0]
	for i := 1; i < m; i++ {
		left[i] += left[i-1] + grid[i][0]
	}

	tops := make([]int, n)
	tops[0] = grid[0][0]
	for i := 1; i < n; i++ {
		tops[i] = tops[i-1] + grid[0][i]
	}

	for i := 1; i < n; i++ {
		top := tops[i]
		left[0] = top
		for j := 1; j < m; j++ {
			left[j] = MyMin(top, left[j]) + grid[j][i]
			top = left[j]
		}
	}

	return left[m-1]
}
