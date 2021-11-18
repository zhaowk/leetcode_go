package leetcode_go

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	if s == t {
		return true
	}

	m := make(map[int]int)

	for i, v := range s {
		if m[int(v)] != 0 {
			if m[int(v)] != int(t[i]) {
				return false
			}
		} else {
			for _, w := range m {
				if w == int(t[i]) {
					return false
				}
			}
			m[int(v)] = int(t[i])
		}
	}
	return true
}
