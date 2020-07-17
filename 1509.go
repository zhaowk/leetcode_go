package leetcode_go

import "sort"

func minDifference(nums []int) int {
	sort.Ints(nums)

	if len(nums) < 5 {
		return 0
	}

	n := len(nums) - 1
	mi := nums[n] - nums[0]

	mi = min(mi, nums[n-3]-nums[0])
	mi = min(mi, nums[n-2]-nums[1])
	mi = min(mi, nums[n-1]-nums[2])
	mi = min(mi, nums[n]-nums[3])
	return mi
}
