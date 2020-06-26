package leetcode_go

func generateMatrix(n int) [][]int {
	result := make([][]int, n)

	if n == 0 {
		return result
	} else if n == 1 {
		return [][]int{{1}}
	}

	for i := 0; i < n; i++ {
		tmp := make([]int, n)
		for j := 0; j < n; j++ {
			tmp[j] = 0
		}
		result[i] = tmp
	}

	l := 0
	for k := 0; k < n/2; k++ {
		start, end := k, n-1-k

		for i, j := start, start; j < end; j++ {
			l++
			result[i][j] = l
		}

		for i, j := start, end; i < end; i++ {
			l++
			result[i][j] = l
		}

		for i, j := end, end; j > start; j-- {
			l++
			result[i][j] = l
		}

		for i, j := end, start; i > start; i-- {
			l++
			result[i][j] = l
		}
		if l == n*n {
			return result
		}
	}

	if n%2 != 0 {
		l++
		result[n/2][n/2] = l
	}

	return result
}
