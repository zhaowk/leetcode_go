package _1xx

import "sort"
import (
	. "github.com/zhaowk/leetcode_go"
)

func maximumGap(nums []int) int {
	mx := 0

	if len(nums) < 2 {
		return 0
	}

	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		mx = MyMax(mx, nums[i]-nums[i-1])
	}

	return mx
}
