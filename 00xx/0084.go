package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

// 暴力, AC
func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	mi, mx := heights[0], heights[0]

	for i := 1; i < len(heights); i++ {
		mi = heights[i]
		mx = MyMax(mx, heights[i])
		for j := i - 1; j >= 0; j-- {
			mi = MyMin(mi, heights[j])
			if mi == 0 {
				break
			}
			mx = MyMax(mx, mi*(i-j+1))
		}
	}

	return mx
}

// 暴力, AC
func largestRectangleArea2(heights []int) int {

	if len(heights) == 0 {
		return 0
	}

	mi, mx := heights[0], heights[0]

	for i := 1; i < len(heights); i++ {
		mi = heights[i]
		mx = MyMax(mx, heights[i])
		for j := i - 1; j >= 0; j-- {
			mi = MyMin(mi, heights[j])
			if mi == 0 {
				break
			}
			mx = MyMax(mx, mi*(i-j+1))
		}
	}

	return mx
}
