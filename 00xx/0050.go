package _0xx

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	}

	pwNeg := false
	if n < 0 {
		pwNeg = true
		n = -n
	}

	neg := 1.0
	if x < 0 {
		if n%2 == 1 {
			neg = -1.0
		}

		x = -x
	}

	count := 0
	y, z, r := 1, x, 1.0
	for count < n {
		y, z = 1, x
		for count+y <= n {
			count += y
			r *= z
			y += y
			z *= z
		}
	}

	//
	r = neg * r
	//
	if pwNeg {
		return 1.0 / r
	}
	return r
}
