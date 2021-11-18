package leetcode_go

func maxProduct_(words []string) int {
	letter := make([]uint32, len(words))
	m := 0

	for idx, word := range words {
		for _, l := range word {
			letter[idx] |= 1 << (l - 'a')
		}

		for i := 0; i < idx; i++ {
			if letter[i]&letter[idx] == 0 {
				m = max(m, len(words[i])*len(word))
			}
		}
	}

	return m
}
