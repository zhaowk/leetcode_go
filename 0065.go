package leetcode_go

func isNumber(s string) bool {
	ls := len(s)
	if ls == 0 {
		return false
	}

	bOp, bNum, bExp, bPoint, bSpace := false, false, false, false, false
	idx := 0
	for ; idx < ls && s[idx] == ' '; idx++ {
	}

	if idx == ls {
		return false
	}

	for i := idx; i < ls; i++ {
		if s[i] != ' ' && bSpace {
			return false
		}

		if s[i] == '+' || s[i] == '-' {
			if bOp || bPoint || bNum {
				return false
			}
			bOp = true
		} else if s[i] == 'e' {
			if !bNum || bExp {
				return false
			}
			bOp = false
			bNum = false
			bPoint = false // 不能有 .
			bExp = true
		} else if s[i] == '.' {
			if bPoint || bExp {
				return false
			}
			bPoint = true
			bOp = true // 不能有op
		} else if s[i] >= '0' && s[i] <= '9' {
			bNum = true
		} else if s[i] == ' ' {
			bSpace = true
		} else {
			return false
		}
	}

	return bNum
}
