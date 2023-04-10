package m

import (
	. "github.com/zhaowk/leetcode_go"
)

func respace(dictionary []string, sentence string) int {

	if len(sentence) == 0 {
		return 0
	}

	if len(dictionary) == 0 {
		return len(sentence)
	}

	m := make(map[string]bool, 0)

	for _, word := range dictionary {
		m[word] = true
	}

	mi := len(sentence)
	s := make([]int, len(sentence)+1)

	s[0] = 0

	for i := 1; i <= len(sentence); i++ {
		if m[sentence[0:i]] {
			s[i] = 0
		} else {
			s[i] = s[i-1] + 1
		}
	}

	for i := 1; i < len(sentence); i++ {
		base := s[i]
		for j := i + 1; j <= len(sentence); j++ {
			if m[sentence[i:j]] {
				s[j] = MyMin(s[j], base)
			} else {
				s[j] = MyMin(s[j], s[j-1]+1)
			}
		}
		if s[len(sentence)] < mi {
			mi = s[len(sentence)]
		}
	}
	return mi
}
