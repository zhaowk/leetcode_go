package _0xx

func sumOfBeauties(nums []int) int {
	mxa := make([]int, len(nums))
	mia := make([]int, len(nums))
	mxa[0] = nums[0]
	mia[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		if mxa[i-1] < nums[i] {
			mxa[i] = nums[i]
		} else {
			mxa[i] = mxa[i-1]
		}
	}

	for i := len(nums) - 2; i >= 0; i-- {
		if mia[i+1] > nums[i] {
			mia[i] = nums[i]
		} else {
			mia[i] = mia[i+1]
		}
	}

	var result int
	for i := 1; i < len(nums)-1; i++ {
		if mxa[i-1] < nums[i] && nums[i] < mia[i+1] {
			result += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			result++
		}
	}
	return result
}
