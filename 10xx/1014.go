package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func MaxScoreSightseeingPair(A []int) int {
	ans, mx := 0, A[0]+0

	for i := 1; i < len(A); i++ {
		ans = MyMax(ans, mx+A[i]-i)
		mx = MyMax(mx, A[i]+i)
	}
	return ans
}
