package leetcode_go

func fullJustify(words []string, maxWidth int) []string {

	totalLen := 0
	result := make([]string, 0)
	start := 0
	for idx, word := range words {
		if len(word)+totalLen <= maxWidth {
			totalLen += len(word) + 1
		} else {
			if idx == start {
				result = append(result, word)
				totalLen = 0
				start = idx
				continue
			} else if idx == start+1 {
				str := words[start]
				for i := 0; i < maxWidth-len(words[start]); i++ {
					str += " "
				}
				result = append(result, str)
				totalLen = len(word) + 1
				start = idx
			} else {
				totalLen -= idx - start
				spaces := (maxWidth - totalLen) / (idx - start - 1)
				moreSpaces := (maxWidth - totalLen) % (idx - start - 1)

				spacesStr := ""
				for i := 0; i < spaces; i++ {
					spacesStr += " "
				}
				moreSpacesStr := spacesStr + " "
				str := ""

				if idx == start+1 {
					result = append(result, word+spacesStr)
				} else {
					for i := start; i < start+moreSpaces; i++ {
						str += words[i] + moreSpacesStr
					}
					str += words[start+moreSpaces]
					for i := start + moreSpaces + 1; i < idx; i++ {
						str += spacesStr + words[i]
					}
					result = append(result, str)
				}

				totalLen = len(word) + 1
				start = idx
			}
		}
	}

	if start < len(words) {
		totalLen = 0
		str := ""
		for i := start; i < len(words); i++ {
			str += words[i]
			totalLen += len(words[i])
			if len(str) < maxWidth {
				str += " "
				totalLen++
			}
		}

		for i := 0; i < maxWidth-totalLen; i++ {
			str += " "
		}

		result = append(result, str)
	}

	return result
}
