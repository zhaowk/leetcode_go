package leetcode_go

func lengthOfLIS(nums []int) int {
	lis := make([]int, len(nums))
	res := 1

	lis[0] = 1

	for i := 1; i < len(nums); i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				lis[i] = max(lis[i], lis[j] + 1)
				res = max(res, lis[i])
			}
		}
	}

	return res
}
