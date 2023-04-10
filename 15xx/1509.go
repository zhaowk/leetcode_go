package _5xx

import "sort"
import (
	. "github.com/zhaowk/leetcode_go"
)

func minDifference(nums []int) int {
	sort.Ints(nums)

	if len(nums) < 5 {
		return 0
	}

	n := len(nums) - 1
	mi := nums[n] - nums[0]

	mi = MyMin(mi, nums[n-3]-nums[0])
	mi = MyMin(mi, nums[n-2]-nums[1])
	mi = MyMin(mi, nums[n-1]-nums[2])
	mi = MyMin(mi, nums[n]-nums[3])
	return mi
}
