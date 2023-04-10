package _0xx

func lengthOfLastWord(s string) int {
	ls := len(s)
	r := 0

	if ls == 0 {
		return 0
	}

	i, j := 0, 0
	for idx, val := range s {
		if val == ' ' {
			i, j = idx, idx
		} else {
			j++
			r = j - i
		}
	}

	return r
}
