package leetcode_go

func isHappy(n int) bool {

	m := make(map[int]bool)

	for {
		if n == 1 {
			return true
		}

		if m[n] {
			return false
		}

		m[n] = true

		t := 0
		for n > 0 {
			t += (n % 10) * (n % 10)
			n /= 10
		}
		n = t
	}

	return false
}
