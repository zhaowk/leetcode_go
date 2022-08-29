package leetcode_go

func searchMatrix1(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	var searchM func(l, u, r, d int) bool
	searchM = func(l, u, r, d int) bool {
		left, up, right, down := l, u, r, d
		if matrix[left][up] == target || matrix[right][down] == target {
			return true
		}

		if matrix[left][up] > target || matrix[right][down] < target {
			return false
		}

		if left == right && up == down {
			return false
		}

		for left <= right && up <= down {
			x, y := (left + right) / 2, (up + down) / 2

			if matrix[x][y] == target {
				return true
			} else if matrix[x][y] < target {
				left, up = x, y
			} else {
				right, down = x, y
			}

			if right - left <= 1 && down - up <= 1 {
				break
			}
		}

		return searchM(l, down, left, d) || searchM(right, u, r, up)
	}

	return searchM(0, 0, len(matrix) - 1, len(matrix[0]) - 1)

	//left, right, up, down := 0, len(matrix), 0, len(matrix[0])
	//
	//for left <= right && up <= down {
	//	x, y := (left + right) / 2, (up + down) / 2
	//
	//	if matrix[x][y] == target {
	//		return true
	//	} else if matrix[x][y] < target {
	//		left, up = x, y
	//	} else {
	//		right, down = x, y
	//	}
	//
	//	if right - left <= 1 && down - up <= 1 {
	//		break
	//	}
	//}
	//
	//var searchBin = func(start, end int, data func(idx int) int) bool {
	//	if data(start) == target || data(end) == target {
	//		return true
	//	}
	//
	//	for start < end {
	//		x := (start + end) / 2
	//		if data(x) == target {
	//			return true
	//		} else if data(x) < target {
	//			start = x
	//		} else {
	//			end = x
	//		}
	//		//
	//		if start + 1 == end {
	//			return false
	//		}
	//	}
	//	return false
	//}
	//
	//if left == right && up == down {
	//	return false
	//} else if left != right && up != down {
	//	return searchBin(0, left, func(idx int) int { return matrix[idx][down]}) ||
	//		searchBin(0, up, func(idx int)int {return matrix[right][idx]})
	//} else if left != right {  // left + 1 == right && up == down
	//	return searchBin(0, up, func(idx int)int {return matrix[right][idx]})
	//} else {  // left == right && up + 1 == down
	//	return searchBin(0, left, func(idx int) int { return matrix[idx][down]})
	//}
}
