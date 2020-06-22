package leetcode_go

func removeDuplicates(nums []int) int {
	nl := len(nums)

	if nl < 2 {
		return nl
	}

	k := 0
	for i := 1; i < nl; i++ {
		if nums[i] != nums[k] {
			k++
			if i > k {
				nums[k] = nums[i]
			}
		}
	}
	return k + 1
}
