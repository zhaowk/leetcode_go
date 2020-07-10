package leetcode_go

func generate(numRows int) [][]int {

	res := make([][]int, 0)

	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})

	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		tmp := make([]int, i+1)
		tmp[0] = 1
		for j := 1; j < i; j++ {
			tmp[j] = res[i-1][j-1] + res[i-1][j]
		}

		tmp[i] = 1
		res = append(res, tmp)
	}
	return res

}
