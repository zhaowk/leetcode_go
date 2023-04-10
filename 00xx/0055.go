package _0xx

func canJump(nums []int) bool {
	idx, ln := 0, len(nums)
	if ln == 0 {
		return false
	}

	if ln == 1 {
		return true
	}

	for i, v := range nums {
		if i > idx {
			return false
		}
		idx = maxC(idx, i+v)
		if idx >= ln-1 {
			return true
		}
	}
	return idx >= ln-1
}

func maxC(x, y int) int {
	if x < y {
		return y
	}
	return x
}
