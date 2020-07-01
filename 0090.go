package leetcode_go

import (
	"sort"
)

var subsetsWithDupResult [][]int

func subsetsWithDup(nums []int) [][]int {
	subsetsWithDupResult = make([][]int, 0)

	subset := make([]int, 0)

	sort.Ints(nums)

	subsetsWithDupR(nums, subset, 0, 0)

	return subsetsWithDupResult
}

func subsetsWithDupR(nums, sub []int, idx, k int) {

	//fmt.Println(sub, idx, k)

	if k > len(nums) {
		return
	}

	if len(sub) == k {
		tmp := make([]int, k)
		copy(tmp, sub)
		subsetsWithDupResult = append(subsetsWithDupResult, tmp)
	}

	for i := idx; i < len(nums); i++ {
		if i > idx && nums[i] == nums[i-1] {
			continue
		}
		subsetsWithDupR(nums, append(sub, nums[i]), i+1, k+1)
	}
}
