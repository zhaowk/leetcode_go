package leetcode_go

import "fmt"

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	times := minutesToTest / minutesToDie
	matrix := make([][]int, times + 1)
	// init
	for idx := range matrix {
		matrix[idx] = make([]int, 11)  // max
	}

	for idx := range matrix[0] {
		matrix[0][idx] = 1
	}

	for i := 1; i < len(matrix); i++ {
		matrix[i][0] = 1
		for j := 1; j < len(matrix[i]); j++ {
			matrix[i][j] = compCoefficient(matrix[i-1], j)
			if i == times && matrix[i][j] >= buckets {
				fmt.Println(matrix)
				return j
			}
		}
	}

	return 0
}

func getCoefficient(n int) []int {
	res := []int{1, 1}
	for i := 2; i <= n; i++ {
		tmp := []int{1}

		for j := 1; j < len(res); j ++ {
			tmp = append(tmp, res[j-1] + res[j])
		}
		tmp = append(tmp, 1)
		res = tmp
	}
	return res
}

func compCoefficient(a []int, n int) int {
	res := 0
	for i, c := range getCoefficient(n) {
		res += c * a[i]
	}
	return res
}
