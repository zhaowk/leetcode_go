package leetcode_go

func searchRange(nums []int, target int) []int {

	start, end := 0, len(nums)-1

	if end < 0 {
		return []int{-1, -1}
	}
	if start == end {
		if nums[0] == target {
			return []int{0, 0}
		} else {
			return []int{-1, -1}
		}
	}

	found := false

	for start <= end {
		mid := (start + end) / 2
		if target <= nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		}
		if target == nums[mid] { //
			found = true
		}
	}

	xstart, xend := 0, len(nums)-1
	for xstart <= xend {
		mid := (xstart + xend) / 2
		if target < nums[mid] {
			xend = mid - 1
		} else if target >= nums[mid] {
			xstart = mid + 1
		}
		if target == nums[mid] { //
			found = true
		}
	}

	if found {
		return []int{start, xend}
	} else {
		return []int{-1, -1}
	}
}
