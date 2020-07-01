package leetcode_go

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	m := len(board)

	if m == 0 {
		return false
	}

	n := len(board[0])

	if n == 0 {
		return false
	}

	b := make([][]bool, m)
	for i := 0; i < m; i++ {
		tmp := make([]bool, n)
		b[i] = tmp
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if realExists(board, b, word, i, j) {
					return true
				}
			}
		}
	}

	return realExists(board, b, word, 0, 0)
}

func realExists(board [][]byte, b [][]bool, s string, i, j int) bool {

	//fmt.Println(s, i, j)

	if len(s) == 0 { // 默认匹配
		return true
	}
	if board[i][j] == s[0] {
		if len(s) == 1 { // 最后一位匹配
			return true
		}
		b[i][j] = true

		if i-1 >= 0 && !b[i-1][j] && realExists(board, b, s[1:], i-1, j) { // Up
			return true
		}

		if j-1 >= 0 && !b[i][j-1] && realExists(board, b, s[1:], i, j-1) { // Left
			return true
		}

		if i+1 < len(board) && !b[i+1][j] && realExists(board, b, s[1:], i+1, j) { // Down
			return true
		}

		if j+1 < len(board[i]) && !b[i][j+1] && realExists(board, b, s[1:], i, j+1) { // Right
			return true
		}

		b[i][j] = false // Not found
	}

	return false
}
