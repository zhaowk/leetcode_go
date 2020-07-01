package leetcode_go

// k > 1 n > 1

var combineResult [][]int

func combine(n int, k int) [][]int {
	combineResult = make([][]int, 0)

	if k > n {
		return combineResult
	}

	tmp := make([]int, 0)

	combineR(n, k, 0, tmp)

	return combineResult
}

func combineR(n, k, t int, r []int) {
	if k == 0 {
		t := make([]int, len(r))
		copy(t, r)
		combineResult = append(combineResult, t)
	} else {
		for i := t + 1; i <= n; i++ {
			combineR(n, k-1, i, append(r, i))
		}
	}
}
