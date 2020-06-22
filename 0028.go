package leetcode_go

func strStr(haystack string, needle string) int {
	hl := len(haystack)
	nl := len(needle)

	if nl == 0 {
		return 0
	}

	for i := 0; i < hl; i++ {
		found := true
		for j := 0; j < nl; j++ {
			if i+j >= hl {
				found = false
				break
			}
			if haystack[i+j] != needle[j] {
				found = false
				break
			}
		}
		if found {
			return i
		}
	}
	return -1
}
