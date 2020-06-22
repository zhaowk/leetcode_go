package leetcode_go

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	ln := len(nums)
	closestN := nums[0] + nums[1] + nums[2]
	closest := myAbs(closestN - target)

	sort.Ints(nums)

	for i := 0; i < ln-2; i++ {

		j := i + 1
		k := ln - 1

		for j < k {
			t := nums[i] + nums[j] + nums[k]
			v := t - target
			vAbs := myAbs(v)
			if t == target {
				return target
			}

			if vAbs < closest {
				closest = vAbs
				closestN = t
			}

			if v < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return closestN
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
