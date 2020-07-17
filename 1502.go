package leetcode_go

import "sort"

func canMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 2 {
		return true
	}

	sort.Ints(arr)

	n := arr[1] - arr[0]

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != n {
			return false
		}
	}
	return true
}
