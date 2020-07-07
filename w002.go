package leetcode_go

func kthFactor(n int, k int) int {
	if k == 1 {
		return 1
	}

	f := 1

	for i := 2; i <= n; i++ {
		if n%i == 0 {
			f++
			if f == k {
				return i
			}
		}
	}
	return -1
}
