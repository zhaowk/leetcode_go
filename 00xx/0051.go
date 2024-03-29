package _0xx

var nQueueResult [][]string

func solveNQueens(n int) [][]string {
	nQueueResult = make([][]string, 0)

	if n == 1 {
		tmp := []string{"Q"}
		nQueueResult = append(nQueueResult, tmp)
		return nQueueResult
	}
	if n < 4 {
		return nQueueResult
	}

	t := make([][]byte, 0)

	for i := 0; i < n; i++ {
		tmp := make([]byte, 0)
		for j := 0; j < n; j++ {
			tmp = append(tmp, '.')
		}
		t = append(t, tmp)
	}
	nQueue(n, t, 0, 0)
	return nQueueResult
}

func nQueue(n int, q [][]byte, row, col int) {
	if row >= n {
		tmp := make([]string, 0)
		for _, bs := range q {
			tmp = append(tmp, string(bs))
		}
		nQueueResult = append(nQueueResult, tmp)
		return
	}
	for i := row; i < n; i++ {
		found := false
		for j := col; j < n; j++ {
			q[i][j] = 'Q'
			if validQueue(q) {
				found = true
				nQueue(n, q, row+1, 0)
			}
			q[i][j] = '.'
			found = false
		}
		if !found {
			return
		}
	}
}

func validQueue(q [][]byte) bool {

	l := len(q)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if q[i][j] == 'Q' {
				for k := 0; k < l; k++ {
					if k != j && q[i][k] == 'Q' {
						return false
					}
					if k != i && q[k][j] == 'Q' {
						return false
					}

					if i-k-1 >= 0 && j-k-1 >= 0 && q[i-k-1][j-k-1] == 'Q' {
						return false
					}

					if i-k-1 >= 0 && j+k+1 < l && q[i-k-1][j+k+1] == 'Q' {
						return false
					}
				}
			}
		}
	}
	return true
}
