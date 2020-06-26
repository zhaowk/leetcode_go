package leetcode_go

func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	for i := 0; i < k-1; i++ {
		nextPermutation(nums)
	}

	r := ""
	for i := 0; i < n; i++ {
		r += string(nums[i] + '0')
	}

	return r
}
