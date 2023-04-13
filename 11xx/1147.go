package _1xx

// 你会得到一个字符串text。你应该把它分成 k 个子字符串(subtext1, subtext2，…， subtextk)，要求满足:
//   - subtext[i] 是 非空字符串
//   - 所有子字符串的连接等于 text ( 即subtext[1] + subtext[2] + ... + subtext[k] == text)
//   - 对于所有 i的有效值( 即1 <= i<= k ) ，subtext[i] == subtext[k - i + 1] 均成立
//
// 返回k可能最大值。
func longestDecomposition(text string) int {
	count := 1
	frontStart, frontEnd := 0, 1
	backEnd := len(text)
	backStart := backEnd - 1

	for frontEnd < backStart {
		if text[frontStart:frontEnd] == text[backStart:backEnd] {
			frontStart = frontEnd
			backEnd = backStart
			count += 2
		}
		frontEnd++
		backStart--
	}

	// last equal
	if frontEnd == backStart && text[frontStart:frontEnd] == text[backStart:backEnd] {
		count++
	}

	return count
}
