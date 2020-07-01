package leetcode_go

func reverseBits(num uint32) uint32 {
	var target uint32 = 0
	for i := 0; i < 32; i++ {
		target <<= 1
		target += num & 1
		num >>= 1
	}
	return target
}
