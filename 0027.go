package leetcode_go

func removeElement(nums []int, val int) int {
	nl := len(nums)

	if nl == 0 {
		return 0
	}

	k := 0
	for i := 0; i < nl; i++ {
		if nums[i] != val {
			if i > k {
				nums[k] = nums[i]
			}
			k++
		}
	}
	return k
}
