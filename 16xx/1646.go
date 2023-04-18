package _6xx

func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}

	results := make([]int, n+1)
	max := 1
	results[0] = 0
	results[1] = 1
	for i := 2; i <= n; i++ {
		if i&1 == 0 { // odd
			results[i] = results[i/2] + results[i/2+1]
		} else {
			results[i] = results[i/2]
		}
		if max < results[i] {
			max = results[i]
		}
	}
	return max
}
