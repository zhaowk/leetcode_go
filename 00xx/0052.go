package _0xx

var nQueueResultNum = 0

func totalNQueens(n int) int {
	nQueueResultNum = 0

	if n == 1 {
		return 1
	}
	if n < 4 {
		return 0
	}

	t := make([][]byte, 0)

	for i := 0; i < n; i++ {
		tmp := make([]byte, 0)
		for j := 0; j < n; j++ {
			tmp = append(tmp, '.')
		}
		t = append(t, tmp)
	}
	nQueueNum(n, t, 0, 0)
	return nQueueResultNum
}

func nQueueNum(n int, q [][]byte, row, col int) {
	if row >= n {
		nQueueResultNum++
		return
	}
	for i := row; i < n; i++ {
		found := false
		for j := col; j < n; j++ {
			q[i][j] = 'Q'
			if validQueue(q) {
				found = true
				nQueueNum(n, q, row+1, 0)
			}
			q[i][j] = '.'
			found = false
		}
		if !found {
			return
		}
	}
}

//func validQueue(q [][]byte) bool {
//
//	l := len(q)
//
//	for i:=0;i<l;i++ {
//		for j:=0;j<l;j++ {
//			if q[i][j] == 'Q' {
//				for k:=0;k<l;k++ {
//					if k != j && q[i][k] == 'Q' {
//						return false
//					}
//					if k != i && q[k][j] == 'Q' {
//						return false
//					}
//
//					if i - k - 1 >= 0 && j - k - 1 >= 0 && q[i - k - 1][j - k - 1] == 'Q' {
//						return false
//					}
//
//					if i - k - 1 >= 0 && j + k + 1 < l && q[i - k - 1][j + k + 1] == 'Q' {
//						return false
//					}
//				}
//			}
//		}
//	}
//	return true
//}
