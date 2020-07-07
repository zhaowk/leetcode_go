package leetcode_go

// 暴力
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}

	if len(s1) == 0 && s2 == s3 {
		return true
	}

	if len(s2) == 0 && s1 == s3 {
		return true
	}

	if len(s1) > 0 && len(s3) > 0 && s1[0] == s3[0] {
		if isInterleave(s1[1:], s2, s3[1:]) {
			return true
		}
	}

	if len(s2) > 0 && len(s3) > 0 && s2[0] == s3[0] {
		if isInterleave(s1, s2[1:], s3[1:]) {
			return true
		}
	}

	return false
}
