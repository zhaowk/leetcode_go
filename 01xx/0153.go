package _1xx

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	mi := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < mi {
			mi = nums[i]
		}
	}
	return mi
}
