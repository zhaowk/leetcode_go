package _1xx

func singleNumber1(nums []int) int {

	res := 0

	for _, n := range nums {
		res ^= n
	}

	return res
}
