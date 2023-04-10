package _2xx

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	start := 0

	for idx, num := range nums {
		if idx-start > k {
			start = idx - k
		}
		for i := start; i < idx; i++ {
			x, y := nums[i], num
			if x > y {
				x, y = y, x
			}
			if y-x <= t {
				return true
			}
		}

	}
	return false
}
