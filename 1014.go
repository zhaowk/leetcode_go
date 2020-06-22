package leetcode_go

func MaxScoreSightseeingPair(A []int) int {
	ans, mx := 0, A[0]+0

	for i := 1; i < len(A); i++ {
		ans = max(ans, mx+A[i]-i)
		mx = max(mx, A[i]+i)
	}
	return ans
}
