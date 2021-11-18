package leetcode_go

import "sort"

func maximumGap(nums []int) int {
	mx := 0

	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		mx = max(mx, nums[i]-nums[i-1])
	}

	return mx
}
