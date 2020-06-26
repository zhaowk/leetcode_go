package leetcode_go

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	ls := len(strs)
	if ls == 0 {
		return result
	}

	tmp := make([]string, 0)
	tmp = append(tmp, strs[0])
	result = append(result, tmp)

	for i := 1; i < ls; i++ {
		found := false
		for idx, value := range result {
			if len(value) > 0 && isAnagrams(strs[i], value[0]) {
				result[idx] = append(result[idx], strs[i])
				found = true
			}
		}
		if !found {
			tmp := make([]string, 0)
			tmp = append(tmp, strs[i])
			result = append(result, tmp)
		}
	}

	return result
}

func isAnagrams(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	c := []byte(b)

	for idx, v := range a {
		found := false
		for i := idx; i < len(c); i++ {
			if v == int32(c[i]) {
				c[idx], c[i] = c[i], c[idx]
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}

	return true
}
