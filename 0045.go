package leetcode_go

func jump(nums []int) int {

	ln := len(nums)
	steps := make([]int, len(nums))

	steps[ln-1] = 0

	for i := ln - 2; i >= 0; i-- {
		mi := 2147483647
		for j := i + 1; j <= i+nums[i]; j++ {
			if j < ln && steps[j] < mi {
				mi = steps[j]
			}
		}
		steps[i] = mi + 1
	}
	return steps[0]
}
