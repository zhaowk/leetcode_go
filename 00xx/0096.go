package _0xx

func numTrees(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	res := make([]int, n+1)
	res[0] = 0
	res[1] = 1
	res[2] = 2

	for i := 3; i <= n; i++ {
		sum := res[i-1] * 2
		for j := 1; j < i-1; j++ {
			sum += res[j] * res[i-j-1]
		}
		res[i] = sum
	}
	return res[n]
}
