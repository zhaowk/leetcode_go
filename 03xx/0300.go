package _3xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	res := 1

	lis[0] = 1

	for i := 1; i < len(nums); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = MyMax(lis[i], lis[j]+1)
				res = MyMax(res, lis[i])
			}
		}
	}

	return res
}
