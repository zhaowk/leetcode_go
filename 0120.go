package leetcode_go

import "math"

func minimumTotal(triangle [][]int) int {
	prev := make([]int, len(triangle)+1)
	result := make([]int, len(triangle)+1)

	if len(triangle) == 0 {
		return 0
	} else if len(triangle) == 1 {
		return triangle[0][0]
	}

	prev[0] = triangle[0][0]
	result[0] = triangle[0][0]
	mi := math.MaxInt32

	for i := 1; i < len(triangle); i++ {
		result[0] = prev[0] + triangle[i][0]
		mi = result[0]
		for j := 1; j < len(triangle[i])-1; j++ {
			result[j] = triangle[i][j] + minM(prev[j-1], prev[j])
			if result[j] < mi {
				mi = result[j]
			}
		}
		result[len(triangle[i])-1] = prev[len(triangle[i])-2] + triangle[i][len(triangle[i])-1]
		if result[len(triangle[i])-1] < mi {
			mi = result[len(triangle[i])-1]
		}

		for idx := 0; idx < len(triangle[i]); idx++ {
			prev[idx] = result[idx]
		}
	}

	return mi
}

func minM(x, y int) int {
	if x < y {
		return x
	}
	return y
}
