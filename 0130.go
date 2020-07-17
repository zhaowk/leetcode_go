package leetcode_go

func solve(board [][]byte) {

	m := len(board)
	if m < 3 {
		return
	}
	n := len(board[0])
	if n < 3 {
		return
	}

	stat := make(map[int]map[int]bool)
	for i := 0; i < m; i++ {
		stat[i] = make(map[int]bool)
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			stat[i][0] = true
		}
		if board[i][n-1] == 'O' {
			stat[i][n-1] = true
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			stat[0][j] = true
		}
		if board[m-1][j] == 'O' {
			stat[m-1][j] = true
		}
	}

	for {
		bl := 0

		for _, v := range stat {
			bl += len(v)
		}

		for i, v := range stat {
			for j := range v {
				if i-1 > 0 && board[i-1][j] == 'O' {
					stat[i-1][j] = true
				}

				if i+1 < m && board[i+1][j] == 'O' {
					stat[i+1][j] = true
				}

				if j-1 > 0 && board[i][j-1] == 'O' {
					stat[i][j-1] = true
				}

				if j+1 < n && board[i][j+1] == 'O' {
					stat[i][j+1] = true
				}
			}
		}

		al := 0
		for _, v := range stat {
			al += len(v)
		}
		if al == bl {
			break
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !stat[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}
