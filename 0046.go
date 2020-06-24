package leetcode_go

var resultPermute [][]int

func permute(nums []int) [][]int {
	resultPermute = make([][]int, 0)
	t := make([]int, 0)
	p(nums, t)
	return resultPermute
}

func p(nums, res []int) {
	if len(res) == len(nums) {
		resultPermute = append(resultPermute, res)
		//fmt.Println(res)
		return
	}
	for _, v := range nums {
		found := false
		for _, r := range res {
			if v == r {
				found = true
			}
		}
		if !found {
			p(nums, append(res, v))
		}
	}
}
