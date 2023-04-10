package _5xx

func numIdenticalPairs(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] == nums[i] {
				res++
			}
		}
	}
	return res
}
