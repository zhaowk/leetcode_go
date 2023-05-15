package _1xx

func lastSubstring(s string) string {
	indexes := make([]int, 1)
	indexes[0] = 0

	for i := 1; i < len(s); i++ {
		if s[i] > s[indexes[0]] {
			indexes = []int{i}
		} else if s[i] == s[indexes[0]] {
			indexes = append(indexes, i)
		}
	}

	pos := 1
	for len(indexes) > 1 {
		target := make([]int, 0)

		var mx uint8
		for _, k := range indexes {
			if k+pos >= len(s) {
				continue
			}

			if s[k+pos] > mx {
				target = []int{k}
				mx = s[k+pos]
			} else if s[k+pos] == mx {
				target = append(target, k)
			}
		}
		pos++
		indexes = target
	}

	return s[indexes[0]:]
}
