package _1xx

func getRow(rowIndex int) []int {

	res := make([][]int, 0)

	if rowIndex == 0 {
		return []int{1}
	}

	res = append(res, []int{1})

	if rowIndex == 1 {
		return []int{1, 1}
	}

	for i := 1; i <= rowIndex; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		for j := 1; j < i; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}

		tmp[i] = 1
		res = append(res, tmp)
	}
	return res[rowIndex]
}
