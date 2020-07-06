package leetcode_go

func grayCode(n int) []int {
	result := make([]int, 1)
	result[0] = 0

	if n == 0 {
		return result
	}

	for i := 0; i < n; i++ {

		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, 1<<i+result[j])
		}
	}

	return result
}
