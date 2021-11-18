package leetcode_go

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return true
	}
	xLeft, xRight, yDown, yUp := rectangles[0][0], rectangles[0][2], rectangles[0][1], rectangles[0][3]
	area := (xRight - xLeft) * (yUp - yDown)

	for i := 1; i < len(rectangles); i++ {
		area += (rectangles[i][2] - rectangles[i][0]) * (rectangles[i][3] - rectangles[i][1])
		xLeft, xRight, yDown, yUp = min(xLeft, rectangles[i][0]), max(xRight, rectangles[i][2]), min(yDown, rectangles[i][1]), max(yUp, rectangles[i][3])
	}

	if area != (xRight-xLeft)*(yUp-yDown) {
		return false
	}

	for i := 0; i < len(rectangles)-1; i++ {
		for j := i + 1; j < len(rectangles); j++ {
			if cross(rectangles[i], rectangles[j]) {
				return false
			}
		}
	}
	return true
}

func cross(x, y []int) bool {
	if y[0] < x[0] {
		x, y = y, x
	}

	if x[2] <= y[0] || x[3] <= y[1] || x[1] >= y[3] {
		return false
	}
	return true
}
