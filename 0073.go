package leetcode_go

func setZeroes(matrix [][]int) {
	m := len(matrix)
	if m == 0 {
		return
	}

	n := len(matrix[0])
	if n == 0 {
		return
	}

	mZero := make([]int, 0)
	nZero := make([]int, 0)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				mZero = append(mZero, i)
				nZero = append(nZero, j)
			}
		}
	}

	for i := 0; i < m; i++ {
		for _, idx := range nZero {
			matrix[i][idx] = 0
		}
	}

	for i := 0; i < n; i++ {
		for _, idx := range mZero {
			matrix[idx][i] = 0
		}
	}
}
