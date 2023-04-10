package _5xx

func winnerSquareGame(n int) bool {
	if n == 1 {
		return true
	}

	win := make([]int, n+1)
	num := make([]int, 0)

	for i := 1; i < n; i++ {
		pow := i * i
		if pow < n {
			num = append(num, pow)
			win[pow] = 1
		} else if pow == n {
			return true
		} else {
			break
		}
	}

	win[0] = 0
	win[1] = 1

	for i := 2; i <= n; i++ {
		if win[i] == 0 {
			for j := 0; j < len(num) && num[j] < i; j++ {
				if win[i-num[j]] == 0 {
					win[i] = 1
					break
				}
			}
		}
	}

	return win[n] == 1
}
