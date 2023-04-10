package _5xx

import "sort"

func rangeSum(nums []int, n int, left int, right int) int {
	res := make([]int, 0)
	result := 0
	for i := 0; i < len(nums); i++ {
		n := 0
		for j := i; j < len(nums); j++ {
			n += nums[j]
			res = append(res, n)
		}
	}
	sort.Ints(res)
	for i := left - 1; i < right; i++ {
		result = (result + res[i]) % 1000000009
	}
	return result
}
