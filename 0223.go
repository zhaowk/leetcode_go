package leetcode_go

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	area := (ax2-ax1)*(ay2-ay1) + (bx2-bx1)*(by2-by1)
	if ax1 >= bx2 || ay1 >= by2 || ax2 <= bx1 || ay2 <= by1 {
		return area
	}

	x, y := 0, 0

	if ax1 < bx1 {
		x = min(ax2, bx2) - bx1
	} else {
		x = min(ax2, bx2) - ax1
	}

	if ay1 < by1 {
		y = min(ay2, by2) - by1
	} else {
		y = min(ay2, by2) - ay1
	}

	return area - x*y
}
