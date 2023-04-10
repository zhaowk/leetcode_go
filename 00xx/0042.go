package _0xx

func trap(height []int) int {
	mx, waterSum, lineSum := 0, 0, 0
	start, mid, end := 0, 0, len(height)-1
	for i, x := range height {
		if x > mx {
			mx = x
			mid = i
		}
	}

	waterLine := 0
	for i := start; i < mid; i++ {
		waterLine = maxw(height[i], waterLine)
		waterSum += waterLine
		lineSum += height[i]
	}
	waterLine = 0
	for i := end; i > mid; i-- {
		waterLine = maxw(height[i], waterLine)
		waterSum += waterLine
		lineSum += height[i]
	}

	return waterSum - lineSum
}

func maxw(x, y int) int {
	if x < y {
		return y
	}
	return x
}
