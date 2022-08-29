package leetcode_go

// n / 3
func majorityElement2(nums []int) []int {
	m := make(map[int]int)
	res := make([]int, 0)
	thres := len(nums) / 3

	for _, v := range nums {
		m[v]++
	}

	for k, v := range m {
		if v > thres {
			res = append(res, k)
		}
	}
	return res
}
