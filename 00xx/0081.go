package _0xx

func search2(nums []int, target int) bool {
	ln := len(nums)

	if ln == 0 {
		return false
	}

	for i := 0; i < ln; i++ {
		if nums[i] == target {
			return true
		}
	}

	return false
}
