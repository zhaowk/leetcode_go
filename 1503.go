package leetcode_go

func getLastMoment(n int, left []int, right []int) int {
	mx := -1
	for i := 0; i < len(left); i++ {
		mx = max(mx, left[i])
	}
	for i := 0; i < len(right); i++ {
		mx = max(mx, n-right[i])
	}
	return mx
}
