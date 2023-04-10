package _0xx

func divide(dividend int, divisor int) int {
	const INT_MAX = 2147483647

	neg := false
	c := 0

	if dividend > 0 {
		dividend = -dividend
		neg = !neg
	} else if dividend == 0 {
		return 0
	}

	// 溢出？
	if dividend == -INT_MAX-1 && divisor == -1 {
		return INT_MAX
	}
	if divisor > 0 {
		divisor = -divisor
		neg = !neg
	}

	for dividend <= divisor {

		if dividend == divisor {
			c += 1
			break
		} else {
			count := 1
			d := divisor
			for dividend <= d {
				dividend -= d
				c += count
				d += d
				count += count
			}
		}
	}

	if neg {
		return -c
	} else {
		return c
	}
}
