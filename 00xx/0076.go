package _0xx

func minWindow(s string, t string) string {

	if len(t) == 0 || len(s) == 0 {
		return ""
	}
	mi := len(s) + 1
	start, end := 0, 0
	tLetter := make(map[byte]int, 0)

	for _, c := range t {
		tLetter[byte(c)]++
	}

	letters := make(map[byte]int, 0)
	positions := make([]int, 0)

	for idx, char := range s {
		if tLetter[byte(char)] > 0 {
			letters[byte(char)]++
			positions = append(positions, idx)
			for validWindow(letters, tLetter) {
				if idx-positions[0] <= mi {
					mi = idx - positions[0]
					start, end = positions[0], idx
				}

				//fmt.Println(mi, start, end)
				letters[s[positions[0]]]--
				if letters[s[positions[0]]] == 0 {
					delete(letters, s[positions[0]])
				}

				positions = positions[1:]
			}
		}
	}

	if mi > len(s) {
		return ""
	}

	return s[start : end+1]
}

func validWindow(m, t map[byte]int) bool {
	for k, v := range t {
		if m[k] < v {
			return false
		}
	}
	return true
}
