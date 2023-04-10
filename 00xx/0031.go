package _0xx

import "sort"

func nextPermutation(nums []int) {
	nl := len(nums)

	if nl < 2 {
		return
	}

	target, rep, pos := nums[0], 0, -1
	for i := 1; i < nl; i++ {
		if nums[i-1] < nums[i] {
			target, rep, pos = nums[i], i, i-1
		} else if pos >= 0 && target > nums[i] && nums[i] > nums[pos] {
			target, rep = nums[i], i
		}
	}

	if pos >= 0 {
		nums[pos], nums[rep] = nums[rep], nums[pos]
	}

	sort.Ints(nums[pos+1:])
}
