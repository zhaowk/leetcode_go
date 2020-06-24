package leetcode_go

//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
// 超时
func isMatch2(s string, p string) bool {
	ls, lp := len(s), len(p)

	if ls == 0 && lp == 0 { // 匹配完成
		return true
	}

	if lp == 0 { //
		return false
	}

	if lp > 0 && p[0] == '*' {
		pos := 1
		for ; pos < len(p) && p[pos] == '*'; pos++ {
		}
		if pos == len(p) {
			return true
		}
		for i := 0; i <= ls; i++ {
			if isMatch2(s[i:], p[pos:]) {
				return true
			}
		}
		return false
	}

	if ls == 0 {
		return false
	}

	if s[0] == p[0] || p[0] == '?' {
		return isMatch2(s[1:], p[1:])
	}

	return false
}

// ac
func isMatch3(s string, p string) bool {
	ls, lp := len(s), len(p)

	if ls == 0 && lp == 0 { // 匹配完成
		return true
	}

	if lp == 0 { //
		return false
	}

	p = subStar(p)

	matches := make([][]bool, 0)
	match := make([]bool, len(s)+1)
	for i := 0; i <= len(s); i++ {
		match[i] = false
	}
	matches = append(matches, match)
	matches[0][0] = true // empty

	for i := 0; i < len(p); i++ {
		tmp := make([]bool, len(s)+1)
		tmp[0] = matches[i][0] && p[i] == '*'
		for j := 0; j < len(s); j++ {

			if p[i] == '*' {
				tmp[j+1] = tmp[j] || matches[i][j] || matches[i][j+1]
			} else if p[i] == '?' {
				tmp[j+1] = matches[i][j]
			} else {
				tmp[j+1] = matches[i][j] && p[i] == s[j]
			}
		}
		matches = append(matches, tmp)
	}

	return matches[len(p)][len(s)]
}

func subStar(p string) string {
	if len(p) < 2 {
		return p
	}

	a := make([]byte, 1)
	a[0] = p[0]
	for i := 1; i < len(p); i++ {
		if '*' != a[len(a)-1] || '*' != p[i] {
			a = append(a, p[i])
		}
	}
	return string(a)
}
