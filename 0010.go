package leetcode_go

// s a-z   p a-z . *
func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)

	if ls == 0 && lp == 0 { // 匹配完成
		return true
	}

	if lp == 0 { //
		return false
	}

	if lp > 1 && p[1] == '*' {
		for i := 0; i <= ls; i++ {
			if matchStar(p[0], s[0:i]) {
				if isMatch(s[i:], p[2:]) {
					return true
				}
			} else {
				break
			}
		}
		return false
	}

	if ls == 0 {
		return false
	}

	if s[0] == p[0] || p[0] == '.' {
		return isMatch(s[1:], p[1:])
	}

	return false
}

func matchStar(f uint8, s string) bool {
	if f == '.' {
		return true
	}

	for _, c := range s {
		if c != int32(f) {
			return false
		}
	}
	return true
}
