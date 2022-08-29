package leetcode_go

func nthUglyNumber(n int) int {
	if n == 1 {
		return 1
	}

	var insertOrAppend = func(a []int, idx, t int) []int {
		for i := idx + 1; i < len(a); i++ {
			if a[i] > t {
				// insert
				tmp := make([]int, len(a) + 1)
				copy(tmp, a[:i])
				tmp[i] = t
				copy(tmp[i+1:], a[i:])
				return tmp
			} else if a[i] == t {
				// exists
				return a
			}
		}

		// the last
		return append(a, t)
	}

	tmp := make([]int, 1)
	tmp[0] = 1
	for i := 0; i < n; i++ {
		tmp = insertOrAppend(tmp, i, tmp[i] * 2)
		tmp = insertOrAppend(tmp, i, tmp[i] * 3)
		tmp = insertOrAppend(tmp, i, tmp[i] * 5)
	}

	return tmp[n-1]
}
