package _0xx

import "sort"

var result2 [][]int

func combinationSum2(candidates []int, target int) [][]int {
	result2 = make([][]int, 0)
	stack := make([]int, 0)
	sort.Ints(candidates)
	combinationSums2(candidates, target, stack)
	return result2
}

func combinationSums2(candidates []int, target int, solve []int) {
	for idx, value := range candidates {
		if target > 0 {
			if idx == len(candidates)-1 {
				if target == value {
					appends(append(solve, value))
					//tmp := make([]int, len(solve))
					//tmp = append(tmp, value)
					//copy(tmp,solve)
					//result2 = append(result2, tmp)
				}
				return
			}
			combinationSums2(candidates[idx+1:], target-value, append(solve, value))
			continue
		} else if target == 0 {
			appends(solve)
			//tmp := make([]int, len(solve))
			//copy(tmp,solve)
			//result2 = append(result2, tmp)
		}
		return
	}
}

func appends(n []int) {
	for _, v := range result2 {
		if len(v) == len(n) {
			equal := true
			for idx, vv := range v {
				if n[idx] != vv {
					equal = false
					break
				}
			}
			if equal {
				return
			}
		}
	}

	tmp := make([]int, len(n))
	copy(tmp, n)
	result2 = append(result2, tmp)
}
