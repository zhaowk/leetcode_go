package _0xx

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int

	ln := len(nums)
	if ln < 3 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < ln-2; i++ {
		for j := i + 1; j < ln-1; j++ {
			k := ln - 1
			for ; k > j; k-- {
				n := nums[i] + nums[j] + nums[k]
				if n == 0 {
					result = myAppend(result, nums[i], nums[j], nums[k])
					n = nums[k]
					break
				} else if n < 0 {
					break
				}
			}
		}
	}
	return result
}

func threeSum1(nums []int) [][]int {
	var result [][]int

	ln := len(nums)
	if ln < 3 {
		return result
	}

	for i := 0; i < ln-2; i++ {
		for j := i + 1; j < ln-1; j++ {
			for k := j + 1; k < ln; k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = myAppend(result, nums[i], nums[j], nums[k])
				}
			}
		}
	}

	return result
}

func myAppend(target [][]int, n ...int) [][]int {

	sort.Ints(n)

	for _, arr := range target {
		var m = true
		for idx, v := range arr {
			if n[idx] != v {
				m = false
				break
			}
		}
		if m {
			return target
		}
	}

	return append(target, n)
}
