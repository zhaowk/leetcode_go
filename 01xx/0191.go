package _1xx

func hammingWeight(num uint32) int {
	cnt := 0

	for num > 0 {
		if 1&num == 1 {
			cnt++
		}
		num >>= 1
	}

	return cnt
}
