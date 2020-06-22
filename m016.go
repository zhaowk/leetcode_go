package leetcode_go

func patternMatching(pattern string, value string) bool {

	lp, lv := len(pattern), len(value)

	if lp == 0 && lv == 0 {
		return true
	}

	if lp == 0 {
		return false
	}

	if lp == 1 {
		return true
	}

	a, b := 0, 0
	for _, v := range pattern {
		switch v {
		case 'a':
			a++
		case 'b':
			b++
		}
	}

	if b == 0 && lv%a != 0 || (a == 0 && lv%b != 0) {
		return false
	}

	if (b == 0 && lv%a == 0) || (a == 0 && lv%b == 0) {
		if lv == 0 {
			return true
		}
		base := max(a, b)
		s, b := 0, base
		vb := value[s:b]

		for b <= lv {
			if value[s:b] != vb {
				return false
			}
			s += base
			b += base
		}
		return true
	}

	maxLa := lv / a

	for i := 0; i <= maxLa; i++ {
		if (lv-a*i)%b != 0 {
			continue
		}

		la, lb := i, (lv-a*i)/b
		flag := true
		va, vb := "", ""
		pos := 0
		for _, v := range pattern {
			if len(va) == la && len(vb) == lb && va == vb {
				return false
			}
			if pos < lv {
				switch v {
				case 'a':
					if len(va) != la {
						va = value[pos : pos+la]
					} else {
						if va != value[pos:pos+la] {
							flag = false
							break
						}
					}
					pos += la
				case 'b':
					if len(vb) != lb {
						vb = value[pos : pos+lb]
					} else {
						if vb != value[pos:pos+lb] {
							flag = false
							break
						}
					}
					pos += lb
				}
			}
			if !flag {
				break
			}
		}
		if flag {
			return true
		}
	}

	return false
}
