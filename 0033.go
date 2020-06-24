package leetcode_go

func search(nums []int, target int) int {

	ln := len(nums)

	if ln == 0 {
		return -1
	}

	start, end := 0, ln-1

	for start <= end {
		first, last := nums[start], nums[end]

		if first == target {
			return start
		}

		if last == target {
			return end
		}

		mid := (start + end) / 2
		left := true
		if target > nums[mid] {
			if nums[mid] < first {
				if target < first { // Right
					left = false
				} else if target > first { // Left
					left = true
				} else { // Found
					return start
				}
			} else {
				left = false
			}
		} else if target < nums[mid] {
			if nums[mid] < first {
				left = true
			} else {
				if target < first { // Right
					left = false
				} else if target > first {
					left = true
				} else {
					return start
				}
			}
		} else {
			return mid
		}

		if left {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
