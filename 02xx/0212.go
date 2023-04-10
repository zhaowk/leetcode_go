package _2xx

func findWords(board [][]byte, words []string) []string {
	res := make([]string, 0)
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	var searchWord func(i, j int, word string) bool

	searchWord = func(i, j int, word string) bool {
		// set visited
		visited[i][j] = true

		// search end, true
		if len(word) == 0 {
			return true
		}

		// check range, =,  check visited, search next
		// up
		if i > 0 && word[0] == board[i-1][j] && !visited[i-1][j] && searchWord(i-1, j, word[1:]) {
			return true
		}

		// down
		if i < len(board)-1 && word[0] == board[i+1][j] && !visited[i+1][j] && searchWord(i+1, j, word[1:]) {
			return true
		}

		// left
		if j > 0 && word[0] == board[i][j-1] && !visited[i][j-1] && searchWord(i, j-1, word[1:]) {
			return true
		}

		// right
		if j < len(board[i])-1 && word[0] == board[i][j+1] && !visited[i][j+1] && searchWord(i, j+1, word[1:]) {
			return true
		}

		// search failed, rollback
		visited[i][j] = false
		return false
	}

	findWord := func(word string) bool {
		if len(word) == 0 {
			return true
		}

		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if board[i][j] == word[0] {
					initVisited(visited)
					if searchWord(i, j, word[1:]) {
						return true
					}
				}
			}
		}
		return false
	}

	for _, word := range words {
		if findWord(word) {
			res = append(res, word)
		}
	}
	return res
}

func initVisited(data [][]bool) {
	for i := range data {
		for j := range data[i] {
			data[i][j] = false
		}
	}
}
