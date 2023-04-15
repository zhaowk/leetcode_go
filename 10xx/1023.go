package _0xx

// 如果我们可以将小写字母插入模式串pattern得到待查询项query，那么待查询项与给定模式串匹配。（我们可以在任何位置插入每个字符，也可以插入 0 个字符。）
// 给定待查询列表queries，和模式串pattern，返回由布尔值组成的答案列表answer。只有在待查项queries[i] 与模式串pattern 匹配时，answer[i]才为 true，否则为 false。
func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))

	ps := make([]string, 0)
	var i int
	for j := 0; j < len(pattern); j++ {
		if pattern[j] >= 'A' && pattern[j] <= 'Z' {
			if i < j {
				ps = append(ps, pattern[i:j])
			}
			i = j
		}
	}

	ps = append(ps, pattern[i:])
	for k, query := range queries {
		result[k] = match(query, ps)
	}
	return result
}

func match(s string, ps []string) bool {
	var (
		i int // string `s` index
	)

	for _, pattern := range ps {
		if len(pattern) == 0 {
			continue
		}
		var j int
		if pattern[0] >= 'A' && pattern[0] <= 'Z' { // 大写字母开头
			for ; s[i] >= 'a' && s[i] <= 'z'; i++ { // 跳过多余的小写字母
			}
			if i >= len(s) || s[i] != pattern[0] { // 匹配大写字母
				return false
			}
			// 匹配到大写字母， 指针移动
			i++
			j++
		}

		for ; j < len(pattern); j++ { // O(N)， 匹配小写字母串
			for ; i < len(s); i++ {
				if s[i] == pattern[j] {
					i++
					break
				} else if s[i] >= 'A' && s[i] <= 'Z' {
					return false
				}
			}
		}

		if j < len(pattern)-1 { // 未匹配完成
			return false
		}
		// 跳过小写字母
		for ; i < len(s) && s[i] >= 'a' && s[i] <= 'z'; i++ {
		}
	}

	// 未匹配的大写字母 ==> 失败
	for ; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			return false
		}
	}
	return true
}
