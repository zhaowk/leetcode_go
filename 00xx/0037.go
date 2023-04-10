package _0xx

func solveSudoku(board [][]byte) {
	solveSudo(board, 0, 0)
}

func solveSudo(board [][]byte, row, col int) bool {

	for row < 9 && col < 9 {
		if board[row][col] == '.' {
			for k := '1'; k <= '9'; k++ {
				board[row][col] = byte(k)
				if isValidSudoku(board) {
					if solveSudo(board, row+(col+1)/9, (col+1)%9) {
						return true
					}
				}
			}
			board[row][col] = '.' // 回退
			return false
		} else {
			row, col = row+(col+1)/9, (col+1)%9
		}
	}
	return true
}
