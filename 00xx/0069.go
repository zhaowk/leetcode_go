package _0xx

func mySqrt(x int) int {
	// hack
	if x <= 0 {
		return 0
	}

	if x == 1 {
		return 1
	}

	i, j := 0, x
	for i < j {
		if i*i == x {
			return i
		}
		if j*j == x {
			return j
		}

		h := (i + j) / 2
		if h*h < x {
			i = h + 1
		} else if h*h == x {
			return h
		} else {
			j = h - 1
		}
	}

	if i*i < x {
		return i
	} else {
		return i - 1
	}
}
