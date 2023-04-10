package _0xx

func spiralOrder(matrix [][]int) []int {
	result := make([]int, 0)

	lm := len(matrix)

	if lm == 0 {
		return result
	}
	ln := len(matrix[0])
	if ln == 0 {
		return result
	}

	if lm == 1 {
		return matrix[0]
	}

	if ln == 1 {
		for _, ma := range matrix {
			result = append(result, ma[0])
		}
		return result
	}

	cirNum := 0
	if lm%2 == 0 {
		cirNum = lm / 2
	} else {
		if ln%2 != 0 && ln < lm {
			cirNum = ln / 2
		} else {
			cirNum = lm / 2
		}
	}

	for n := 0; n < cirNum; n++ {
		start, end := n, ln-1-n
		for i, j := start, start; j < end; j++ {
			result = append(result, matrix[i][j])
		}
		start, end = n, lm-1-n
		for i, j := start, ln-1-n; i < end; i++ {
			result = append(result, matrix[i][j])
		}

		start, end = ln-1-n, n
		for i, j := lm-1-n, start; j > end; j-- {
			result = append(result, matrix[i][j])
		}
		start, end = lm-1-n, n
		for i, j := start, n; i > end; i-- {
			result = append(result, matrix[i][j])
		}
		if len(result) == lm*ln {
			return result
		}
	}

	if lm%2 != 0 {
		if lm <= ln {
			start, end := lm/2, ln-lm/2
			for j := start; j < end; j++ {
				result = append(result, matrix[start][j])
			}
		} else {
			if ln%2 != 0 {
				start, end := ln/2, lm-ln/2
				for i := start; i < end; i++ {
					result = append(result, matrix[i][start])
				}
			}
		}
	}

	return result
}
