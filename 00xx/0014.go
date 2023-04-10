package _0xx

func longestCommonPrefix(strs []string) string {
	var s string

	if len(strs) == 0 {
		return ""
	}

	if len(strs[0]) == 0 {
		return ""
	}

	for i := 1; i <= len(strs[0]); i++ {

		for _, v := range strs {
			if len(v) < i || v[i-1] != strs[0][i-1] {
				return s
			}
		}

		s = strs[0][:i]
	}
	return s
}
