package leetcode_go

import "math/rand"

type Solution struct {
	idx []int
	nums []int
}

func Constructor2(nums []int) Solution {
	idx := make([]int, len(nums))
	for i := range idx {
		idx[i] = i
	}
	return Solution{
		nums: nums,
		idx: idx,
	}
}


func (this *Solution) Reset() []int {
	return this.nums
}


func (this *Solution) Shuffle() []int {
	rand.Shuffle(len(this.idx), func(i, j int) {
		this.idx[i], this.idx[j] = this.idx[j], this.idx[i]
	})

	result := make([]int, len(this.nums))

	for i := range result {
		result[i] = this.nums[this.idx[i]]
	}
	return result
}
