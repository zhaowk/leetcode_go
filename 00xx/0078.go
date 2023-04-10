package _0xx

var subsetsResult [][]int

func subsets(nums []int) [][]int {
	subsetsResult = make([][]int, 0)

	subset := make([]int, 0)

	subsetsR(nums, subset, 0, 0)

	return subsetsResult
}

func subsetsR(nums, sub []int, idx, k int) {

	if k > len(nums) {
		return
	}

	if len(sub) == k {
		tmp := make([]int, k)
		copy(tmp, sub)
		subsetsResult = append(subsetsResult, tmp)
	}

	for i := idx; i < len(nums); i++ {
		subsetsR(nums, append(sub, nums[i]), i+1, k+1)
	}
}
