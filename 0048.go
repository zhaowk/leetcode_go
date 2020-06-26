package leetcode_go

func rotate(matrix [][]int) {
	lm := len(matrix)
	if lm < 2 {
		return
	}

	for i := 0; i < lm/2; i++ {
		for j := i; j < lm-i-1; j++ {
			matrix[i][j], matrix[j][lm-i-1], matrix[lm-i-1][lm-j-1], matrix[lm-j-1][i] =
				matrix[lm-j-1][i], matrix[i][j], matrix[j][lm-i-1], matrix[lm-i-1][lm-j-1]
		}
	}
}
