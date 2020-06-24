package leetcode_go

import "strconv"

func countAndSay(n int) string {

	if n == 0 {
		return ""
	}

	s, p := "1", ""

	for i := 1; i < n; i++ {
		before := s[0]
		count := 1
		idx := 1
		for ; idx < len(s); idx++ {
			if before == s[idx] {
				count++
			} else {
				p += strconv.Itoa(count) + string(before)
				before = s[idx]
				count = 1
			}
		}
		p += strconv.Itoa(count) + string(before)
		s, p = p, ""
	}
	return s
}
