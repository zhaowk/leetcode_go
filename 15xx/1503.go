package _5xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func getLastMoment(n int, left []int, right []int) int {
	mx := -1
	for i := 0; i < len(left); i++ {
		mx = MyMax(mx, left[i])
	}
	for i := 0; i < len(right); i++ {
		mx = MyMax(mx, n-right[i])
	}
	return mx
}
