package _0xx

var permuteUniqueResult [][]int

func permuteUnique(nums []int) [][]int {
	permuteUniqueResult = make([][]int, 0)
	t := make([]int, 0)
	permuteHelper(nums, t)
	return permuteUniqueResult
}

func permuteHelper(nums, res []int) {
	if len(res) == len(nums) {
		tmp := make([]int, len(res))
		for idx, v := range res {
			tmp[idx] = nums[v]
		}
		permuteUniqueResult = appendUnique(permuteUniqueResult, tmp)
		//fmt.Println(res,tmp)
		return
	}
	for idxN := range nums {
		found := false
		for _, idxR := range res {
			if idxN == idxR {
				found = true
			}
		}
		if !found {
			permuteHelper(nums, append(res, idxN))
		}
	}
}

func appendUnique(res [][]int, p []int) [][]int {
	for _, v := range res {
		if arrayEq(v, p) {
			return res
		}
	}
	return append(res, p)
}

func arrayEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
