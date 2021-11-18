package leetcode_go

func trailingZeroes(n int) int {
	c := 0
	for n > 0 {
		n /= 5 // 5的个数 + 25的个数 + 125的个数 + ...
		c += n
	}
	return c
}
