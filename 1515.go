package leetcode_go

import (
	"math"
)

func getMinDistSum(positions [][]int) float64 {

	l, r := 0.0, 100.0
	mid, midr := 0.0, 0.0

	for r-l > 0.00000001 {
		mid = (l + r) / 2.0
		midr = (mid + r) / 2.0
		if dist2(mid, positions) > dist2(midr, positions) {
			l = mid
		} else {
			r = midr
		}
	}

	return dist2(l, positions)
}

func dist(x, y float64, posi [][]int) float64 {
	res := 0.0
	for idx := range posi {
		res += math.Sqrt((x-float64(posi[idx][0]))*(x-float64(posi[idx][0])) + (y-float64(posi[idx][1]))*(y-float64(posi[idx][1])))
	}
	return res
}

func dist2(x float64, posi [][]int) float64 {
	l, r := 0.0, 100.0 // , 100.0/3.0, 200.0/3.0
	mid, midr := 0.0, 0.0

	for r-l > 0.00000001 {
		mid = (l + r) / 2.0
		midr = (mid + r) / 2.0
		if dist(x, mid, posi) > dist(x, midr, posi) {
			l = mid
		} else {
			r = midr
		}
	}

	return dist(x, l, posi)
}
