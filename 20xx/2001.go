package _0xx

func interchangeableRectangles(rectangles [][]int) int64 {
	index := make(map[float64]int64)

	for _, v := range rectangles {
		index[float64(v[0])/float64(v[1])]++
	}

	var result int64
	for _, v := range index {
		if v > 1 {
			result += v * (v - 1) / 2
		}
	}
	return result
}
