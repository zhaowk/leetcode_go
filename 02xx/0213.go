package _2xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var (
		result0, result1 = 0, 0
		no0res0, no0res1 = 0, 0
	)

	for idx, val := range nums {
		if idx == 0 {
			result1 = val
		} else if idx == 1 {
			result1, result0 = MyMax(result1, val), result1
			no0res1 = val
		} else if idx == len(nums)-1 {
			no0res1, no0res0 = MyMax(no0res1, no0res0+val), no0res1
		} else {
			result1, result0 = MyMax(result1, result0+val), result1
			no0res1, no0res0 = MyMax(no0res1, no0res0+val), no0res1
		}
	}

	return MyMax(result1, no0res1)
}
