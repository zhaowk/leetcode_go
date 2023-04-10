package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	s := make([]int, len(word1)+1)

	for i := 0; i <= len(word1); i++ {
		s[i] = i
	}

	for i := 1; i <= len(word2); i++ {
		prev := s[0]
		s[0]++

		for j := 1; j <= len(word1); j++ {
			curr := s[j]

			if word1[j-1] != word2[i-1] {
				s[j] = MyMin(prev, s[j-1], s[j]) + 1
			} else {
				s[j] = MyMin(prev, s[j-1]+1, s[j]+1)
			}

			prev = curr
		}
		//fmt.Println(s)
	}

	return s[len(s)-1]
}
