package _2xx

//func maxSlidingWindow(nums []int, k int) []int {
//	res := make([]int, 0)
//	for i := 0; i <= len(nums) - k; i++ {
//		maxi := nums[i]
//		for j := i + 1; j < i + k; j++ {
//			if maxi < nums[j] {
//				maxi = nums[j]
//			}
//		}
//		res = append(res, maxi)
//	}
//	return res
//}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	maxi, idx := nums[0], 0
	for i := 1; i < k; i++ {
		if maxi <= nums[i] {
			maxi, idx = nums[i], i
		}
	}
	res = append(res, maxi)

	for i := 1; i <= len(nums)-k; i++ {
		if nums[i+k-1] >= maxi {
			maxi, idx = nums[i+k-1], i+k-1
		}

		if i > idx {
			maxi, idx = nums[i], i
			for j := 1; j < k; j++ {
				if maxi <= nums[i+j] {
					maxi, idx = nums[i+j], i+j
				}
			}
		}

		res = append(res, maxi)
	}
	return res
}
