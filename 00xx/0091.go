package _0xx

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	if s[0] == '0' {
		return 0
	}

	prev, curr := 1, 1 //

	for i := 1; i < len(s); i++ {
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] < '7') {
			if s[i] == '0' {
				curr = prev
			} else {
				prev, curr = curr, prev+curr
			}
		} else {
			if s[i] == '0' {
				return 0
			}
			prev = curr
		}
	}
	return curr
}
