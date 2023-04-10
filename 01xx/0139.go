package _1xx

// 超时
func wordBreak(s string, wordDict []string) bool {
	ls := len(s)
	if ls == 0 {
		return true
	}

	for _, word := range wordDict {
		if len(word) <= ls && s[:len(word)] == word {
			if wordBreak(s[len(word):], wordDict) {
				return true
			}
		}
	}

	return false
}

func wordBreak2(s string, wordDict []string) bool {
	ls, lw := len(s), len(wordDict)

	if ls == 0 {
		return true
	}

	if lw == 0 {
		return false
	}

	resMap := make([]bool, ls+1)
	resMap[0] = true
	for i := 0; i < ls; i++ {
		if resMap[i] {
			for _, word := range wordDict {
				if i+len(word) <= ls && s[i:i+len(word)] == word {
					resMap[i+len(word)] = true
				}
			}
		}
	}

	return resMap[ls]
}
