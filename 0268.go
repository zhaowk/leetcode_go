package leetcode_go

func missingNumber(nums []int) int {

	for i := 0; i < len(nums); i++ {
		j := nums[i]
		if j != i {
			for {
				if j < len(nums) {
					nums[j], j = j, nums[j]
					if j <= i || j < len(nums) && nums[j] == j {
						nums[j] = j
						break
					}
				} else { // nums[i] >= len(nums)
					break
				}
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i {
			return i
		}
	}

	return len(nums)
}
