package leetcode_go

func firstMissingPositive(nums []int) int {
	ln := len(nums)

	bitmap := make([]uint8, ((ln+1)/8)+1)

	for _, n := range nums {
		if n > 0 && (n-1)/8 < len(bitmap) {
			bitmap[(n-1)/8] |= 1 << ((n - 1) % 8)
		}
	}

	for idx, v := range bitmap {
		if v != 255 {
			j := 0
			for (v>>(j%8))&1 != 0 {
				j++
			}
			return idx*8 + j + 1
		}
	}
	return 0
}
