package leetcode_go

func countBits(n int) []int {
	res := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		bits := 0

		for j := i; j != 0; j >>= 1 {
			if j & 1 == 1 {
				bits++
			}
		}
		res[i] = bits
	}
	return res
}
