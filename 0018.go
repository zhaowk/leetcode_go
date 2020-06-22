package leetcode_go

import "sort"

func fourSum(nums []int, target int) [][]int {
	var result [][]int

	ln := len(nums)
	if ln < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < ln-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < ln-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}

			for k := j + 1; k < ln-1; k++ {
				if k > j+1 && nums[k-1] == nums[k] {
					continue
				}

				for l := ln - 1; l > k; l-- {
					n := nums[i] + nums[j] + nums[k] + nums[l]
					if n == target {
						result = append(result, []int{nums[i], nums[j], nums[k], nums[l]})
						break
					} else if n < target {
						break
					}
				}
			}
		}
	}
	return result
}
