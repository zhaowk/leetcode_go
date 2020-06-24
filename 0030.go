package leetcode_go

func findSubstring(s string, words []string) []int {
	result := make([]int, 0)
	matches := make([]bool, len(words))
	count := 0
	makeFalse(matches)

	ls := len(s)
	lws := len(words)

	if lws == 0 {
		// TODO 边界条件
		return result
	}

	lw := len(words[0])

	if lw == 0 {
		// TODO 边界条件
		return result
	}

	if ls < lw*lws {
		return result
	}

	for i := 0; i < ls; i++ {
		for j := 0; j < lws; j++ {
			startPos := i + j*lw
			if startPos+lw > ls {
				break
			}
			tmp := s[startPos : startPos+lw]
			match := false
			for k := 0; k < lws; k++ {
				if tmp == words[k] && !matches[k] {
					count++
					matches[k] = true
					match = true
					break
				}
			}

			if !match { // 重新开始匹配
				makeFalse(matches)
				count = 0
				break
			}
		}
		if count == lws {
			result = append(result, i)
		}
		makeFalse(matches)
		count = 0
	}

	return result
}

func makeFalse(b []bool) {
	for idx := range b {
		b[idx] = false
	}
}
