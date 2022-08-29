package leetcode_go

func gameOfLife(board [][]int)  {
	var findAlive = func(i, j int) int {
		cnt := 0
		if i > 0 {
			cnt += board[i - 1][j] & 1
			if j > 0 {
				cnt += board[i - 1][j - 1] & 1
			}

			if j < len(board[i]) - 1 {
				cnt += board[i - 1][j + 1] & 1
			}
		}

		if j > 0 {
			cnt += board[i][j - 1] & 1
			if i < len(board) - 1 {
				cnt += board[i + 1][j - 1] & 1
			}
		}

		if i < len(board) - 1 {
			cnt += board[i + 1][j] & 1
		}

		if j < len(board[i]) - 1 {
			cnt += board[i][j + 1] & 1
		}

		if i < len(board) - 1 && j < len(board[i]) - 1  {
			cnt += board[i + 1][j + 1] & 1
		}
		return cnt
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			alive := findAlive(i, j)
			if board[i][j] & 1 == 0 {  // death
				if alive == 3 {
					board[i][j] |= 2 //
				}
			} else {  // alive
				switch alive {
				case 2, 3:
					board[i][j] |= 2  // alive
				default:
					//  board[i][j] &= ^2
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] >>= 1
		}
	}
}
