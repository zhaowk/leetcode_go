package _0xx

// 计数排序 两趟
func sortColors(nums []int) {
	zero, one, two := 0, 0, 0

	for _, v := range nums {
		if v == 0 {
			zero++
		} else if v == 1 {
			one++
		} else if v == 2 {
			two++
		}
	}

	for i := 0; i < zero; i++ {
		nums[i] = 0
	}

	for i := zero; i < zero+one; i++ {
		nums[i] = 1
	}

	for i := zero + one; i < zero+one+two; i++ {
		nums[i] = 2
	}
}
