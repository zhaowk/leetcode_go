package _1xx

import "math"
import (
	. "github.com/zhaowk/leetcode_go"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mx := math.MinInt32

	for i := 0; i < len(nums); i++ {
		fact := nums[i]
		mx = MyMax(mx, fact)
		for j := i + 1; j < len(nums); j++ {
			fact *= nums[j]
			mx = MyMax(mx, fact)
		}
	}
	return mx
}

func maxProduct2(nums []int) int {
	mx1, mi1, r := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mi := mx1, mi1
		mx1 = MyMax(nums[i], mx*nums[i], mi*nums[i])
		mi1 = MyMin(nums[i], mx*nums[i], mi*nums[i])
		r = MyMax(mx1, r)
	}
	return r
}
