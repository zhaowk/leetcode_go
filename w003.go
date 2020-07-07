package leetcode_go

//1 <= nums.length <= 10^5
//nums[i] 要么是 0 要么是 1 。
func longestSubarray(nums []int) int {

	ln := len(nums)
	if ln < 2 {
		return 0
	}

	mx, start, end := 0, 0, 0
	one, zero := false, false

	for i := 0; i < ln; i++ {
		if nums[i] == 0 {
			if one {
				if zero {
					if nums[i-1] == 1 {
						end = i - start - 1
						start = i
					} else {
						end = 0
						start = i
						zero = false
					}

					one = false
				} else {
					zero = true
					start = i
				}
			} else {
				end = 0
				start = i
				zero = false
			}
		} else if nums[i] == 1 {
			one = true
			end++
			mx = max(mx, end)
		}
	}

	if mx == ln {
		return mx - 1
	}

	return mx
}

//func max(x,y int) int {
//	if x<y {
//		return y
//	}
//	return x
//}
