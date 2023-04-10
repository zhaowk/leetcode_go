package _8xx

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}

	idx, end, res := -1, false, false
	m := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			if end {
				return false
			}
			if idx == -1 {
				idx = i
			} else {
				if s[i] != goal[idx] || s[idx] != goal[i] {
					return false
				}
				idx = -1
				end = true
			}
		} else {
			m[s[i]]++
			if m[s[i]] >= 2 {
				res = true
			}
		}
	}
	return end || (idx == -1 && res)
}
