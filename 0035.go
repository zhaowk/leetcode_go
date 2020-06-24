package leetcode_go

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	start, end := 0, len(nums)-1

	for start <= end {
		mid := (start + end) / 2
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return start
}
