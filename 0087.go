package leetcode_go

func isScramble(s1 string, s2 string) bool {

	if s1 == s2 || s1 == reverseStr(s2) {
		return true
	}

	if !isLike(s1, s2) {
		return false
	}

	if len(s1) == 2 {
		return true
	}

	for i := 1; i < len(s1); i++ {
		if isScramble(s1[:i], s2[:i]) && isScramble(s1[i:], s2[i:]) {
			return true
		}
		if isScramble(s1[:i], s2[len(s2)-i:]) && isScramble(s1[i:], s2[:len(s2)-i]) {
			return true
		}
	}

	return false
}

func isLike(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[byte]int, 0)

	for idx := range s1 {
		m[s1[idx]]++
		m[s2[idx]]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func reverseStr(s string) string {
	t := []byte(s)

	for i := 0; i < len(t)/2; i++ {
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	return string(t)
}
