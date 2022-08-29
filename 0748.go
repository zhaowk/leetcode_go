package leetcode_go

import "math"

func shortestCompletingWord(licensePlate string, words []string) string {
	m := make(map[int32]int)
	var idx = -1
	var mi = math.MaxInt32

	var isCover = func(r map[int32]int) bool {
		for k, v := range m {
			if r[k] < v {
				return false
			}
		}
		return true
	}

	for _, l := range licensePlate {
		if l >= 'a' && l <= 'z' {
			m[l - 'a']++
		} else if l >= 'A' && l <= 'Z' {
			m[l - 'A']++
		}
	}

	for i, word := range words {
		r := make(map[int32]int)
		for _, l := range word {
			if l >= 'a' && l <= 'z' {
				r[l - 'a']++
			} else if l >= 'A' && l <= 'Z' {
				r[l - 'A']++
			}
		}

		if isCover(r) && len(word) < mi {
			idx = i
			mi = len(word)
		}
	}
	if idx < 0 {
		return ""
	}
	return words[idx]
}
