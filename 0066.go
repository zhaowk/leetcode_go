package leetcode_go

func plusOne(digits []int) []int {
	ld := len(digits)
	if ld == 0 {
		return []int{1}
	}

	for i := ld - 1; i >= 0; i-- {
		switch digits[i] {
		case 9:
			digits[i] = 0
		default:
			digits[i]++
			return digits
		}
	}
	r := []int{1}
	return append(r, digits...)
}
