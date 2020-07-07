package leetcode_go

func findMaxValueOfEquation(points [][]int, k int) int {
	mx := -2147483648

	prev := 0
	lp := len(points)
	for i := 1; i < lp; {
		if points[i][0]-points[prev][0] > k {
			prev++
			i = prev + 1
		} else {
			if prev < i {
				mx = maxF(mx, points[i][1]+points[prev][1]+points[i][0]-points[prev][0])
			}
			i++
		}
	}

	for prev < lp-1 {
		mx = maxF(mx, points[lp-1][1]+points[prev][1]+points[lp-1][0]-points[prev][0])
		prev++
	}

	return mx
}

func maxF(x, y int) int {
	if x < y {
		return y
	}
	return x
}
