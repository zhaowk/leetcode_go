package leetcode_go

func maximalSquare(matrix [][]byte) int {
	target := 0
	tmp := make([][]int, len(matrix))
	for idx, m := range matrix {
		tmp[idx] = make([]int, len(m))
		tmp[idx][0] = int(m[0] - '0')
		if m[0] == '1' {
			target = 1
		}
	}

	for idx, v := range matrix[0] {
		if v == '1' {
			target = 1
		}
		tmp[0][idx] = int(v - '0')
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				tmp[i][j] = 0
			} else {
				tmp[i][j] = min(tmp[i-1][j-1], tmp[i][j-1], tmp[i-1][j]) + 1
				target = max(target, tmp[i][j])
			}
		}
	}
	return target * target
}
