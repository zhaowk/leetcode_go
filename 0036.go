package leetcode_go

func isValidSudoku(board [][]byte) bool {
	const SodokuSize = 9
	if len(board) != SodokuSize {
		return false
	}

	validArr := make([]bool, SodokuSize)
	makeFalse(validArr)

	for i := 0; i < SodokuSize; i++ {
		makeFalse(validArr)
		for j := 0; j < SodokuSize; j++ {
			n := board[i][j]
			pos := n - '1'

			if n >= '0' && n <= '9' {
				if validArr[pos] {
					return false
				}
				validArr[n-'1'] = true
			}
		}
		makeFalse(validArr)
		for j := 0; j < SodokuSize; j++ {
			n := board[j][i]
			pos := n - '1'

			if n >= '0' && n <= '9' {
				if validArr[pos] {
					return false
				}
				validArr[n-'1'] = true
			}
		}
		makeFalse(validArr)
		for j := 0; j < SodokuSize; j++ {
			n := board[i*3%9+j/3][i/3*3+j%3]
			pos := n - '1'
			if n >= '0' && n <= '9' {
				if validArr[pos] {
					return false
				}
				validArr[n-'1'] = true
			}
		}
	}
	return true
}

//func makeFalse(b []bool) {
//	for idx, _ := range b {
//		b[idx] = false
//	}
//}
