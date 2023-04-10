package _0xx

func maxSubArray(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	mx, prev := nums[0], nums[0]

	for i := 1; i < ln; i++ {
		prev = maxS(nums[i], prev+nums[i])
		mx = maxS(mx, prev)
	}

	return mx
}

func maxS(x, y int) int {
	if x < y {
		return y
	}
	return x
}
