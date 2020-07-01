package leetcode_go

func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}

	idx := 2
	prev1, prev := nums[0], nums[1]

	for i := 2; i < len(nums); i++ {
		prev = nums[i-1]
		if nums[i] != prev1 {
			nums[idx] = nums[i]
			idx++
		}
		prev1, prev = prev, nums[i]
	}

	return idx
}
