package leetcode_go

func insert(intervals [][]int, newInterval []int) [][]int {

	li, ln := len(intervals), len(newInterval)

	if li == 0 {
		return append(intervals, newInterval)
	}

	result := make([][]int, 0)
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < li; i++ {
		if intervals[i][0] <= end {
			end = maxI(end, intervals[i][1])
		} else {
			result = append(result, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		}
	}

	result = append(result, []int{start, end})

	if ln == 0 {
		return result
	}

	intervals = result
	result = make([][]int, 0)

	start, end = intervals[0][0], intervals[0][1]
	if newInterval[1] < intervals[0][0] {
		result = append(result, newInterval)
		return append(result, intervals...)
	}

	for idx, interval := range intervals {

		if newInterval[0] > interval[1] { // not merge
			result = append(result, interval)
		} else {
			start = minI(interval[0], newInterval[0])
			for i := idx; i < li; i++ {
				if newInterval[1] < intervals[i][0] {
					end = maxI(end, newInterval[1])
					result = append(result, []int{start, end})
					break
				} else if newInterval[1] <= intervals[i][1] {
					end = maxI(end, intervals[i][1])
					result = append(result, []int{start, end})
					idx++
					if idx == li {
						return result
					}
					break
				} else {
					idx++
				}
			}
			if idx >= li {
				end = maxI(newInterval[1], intervals[li-1][1])
				result = append(result, []int{start, end})
			} else {
				result = append(result, intervals[idx:]...)
			}
			return result
		}
	}

	return append(intervals, newInterval)
}

func maxI(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func minI(x, y int) int {
	if x < y {
		return x
	}
	return y
}
