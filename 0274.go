package leetcode_go

func hIndex(citations []int) int {
	idx := make([]int, len(citations) + 1)

	for _, c := range citations {
		for i := 1; i < len(idx) && i <= c; i ++ {
			idx[i]++
		}
	}

	for i := len(idx) - 1; i >= 0; i-- {
		if idx[i] >= i {
			return i
		}
	}
	return 0
}
