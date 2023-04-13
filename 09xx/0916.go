package _9xx

func wordSubsets(words1 []string, words2 []string) []string {
	ms := make([]map[rune]int, len(words2))
	for idx, word := range words2 {
		ms[idx] = make(map[rune]int)
		for _, r := range word {
			ms[idx][r]++
		}
	}

	m := make(map[rune]int)
	r := make([]string, 0)

	for _, _m := range ms {
		for k, v := range _m {
			if m[k] < v {
				m[k] = v
			}
		}
	}

	for _, word := range words1 {
		if contains(m, word) {
			r = append(r, word)
		}
	}
	return r
}

func contains(m map[rune]int, s string) bool {
	tmp := make(map[rune]int)
	for k, v := range m {
		tmp[k] = v
	}

	for _, r := range s {
		if _, ok := tmp[r]; ok {
			tmp[r]--
			if tmp[r] == 0 {
				delete(tmp, r)
			}
		}
		if len(tmp) == 0 {
			return true
		}
	}
	return false
}
