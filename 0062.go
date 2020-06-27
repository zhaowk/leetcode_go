package leetcode_go

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m == 1 || n == 1 {
		return 1
	}

	if m > n {
		m, n = n, m
	}

	tmp := make([]int, m)
	for idx := range tmp {
		tmp[idx] = 1
	}

	for i := 1; i < n; i++ {
		top := 1
		for j := 1; j < m; j++ {
			tmp[j] = tmp[j] + top
			top = tmp[j]
		}
	}

	return tmp[m-1]
}
