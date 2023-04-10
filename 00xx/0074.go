package _0xx

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}

	n := len(matrix[0])
	if n == 0 {
		return false
	}

	ms, me, h := 0, m-1, (m-1)/2

	if target <= matrix[0][n-1] {
		h = 0
	} else if target >= matrix[m-1][0] {
		h = m - 1
	} else {
		for ms < me {
			if target == matrix[ms][0] || target == matrix[me][0] {
				return true
			}
			h = (me + ms) / 2
			if target < matrix[h][0] {
				if h > 0 {
					if target > matrix[h-1][0] { // h - 1 行
						h = h - 1
						break
					} else if target == matrix[h-1][0] {
						return true
					}
				} else if h == 0 {
					break
				}

				me = h - 1
			} else if target == matrix[h][0] {
				return true
			} else { // target > matrix[h][0]
				if h < m-1 {
					if target < matrix[h+1][0] { // h 行
						break
					} else if target == matrix[h+1][0] {
						return true
					}
				}
				ms = h + 1
			}
		}
	}

	ns, ne := 0, n-1

	for ns <= ne {
		if target == matrix[h][ns] || target == matrix[h][ne] {
			return true
		}

		ha := (ns + ne) / 2
		if target < matrix[h][ha] {
			ne = ha - 1
		} else if target == matrix[h][ha] {
			return true
		} else { // target > matrix[h][ha]
			ns = ha + 1
		}
	}

	return false
}
