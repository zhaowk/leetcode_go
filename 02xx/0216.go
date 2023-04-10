package _2xx

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return nil
	}

	if k == 1 && n > 0 && n < 10 {
		return [][]int{{n}}
	}

	res := make([][]int, 0)

	var findSum3 func(start, k, n int, tmp []int)

	findSum3 = func(start, k, n int, tmp []int) {
		if n < start {
			return
		}
		if k == 1 {
			if n <= 9 {
				t := make([]int, len(tmp)+1)
				copy(t, tmp)
				t[len(tmp)] = n
				res = append(res, t)
			}
			return
		}
		for i := start; i <= 9; i++ {
			tmp = append(tmp, i)
			findSum3(i+1, k-1, n-i, tmp)
			tmp = tmp[:len(tmp)-1]
		}
	}

	findSum3(1, k, n, make([]int, 0))

	return res
}
