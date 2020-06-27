package leetcode_go

func climbStairs(n int) int {
	if n < 4 {
		return n
	}

	prev, curr := 0, 1

	for i := 0; i < n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
