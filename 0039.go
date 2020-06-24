package leetcode_go

var result [][]int

func combinationSum(candidates []int, target int) [][]int {
	result = make([][]int, 0)
	stack := make([]int, 0)
	combinationSums(candidates, target, stack)
	return result
}

func combinationSums(candidates []int, target int, solve []int) {
	for idx, value := range candidates {
		if target > 0 {
			combinationSums(candidates[idx:], target-value, append(solve, value))
			continue
		} else if target == 0 {
			tmp := make([]int, len(solve))
			copy(tmp, solve)
			result = append(result, tmp)
		}
		return
	}
	return
}
