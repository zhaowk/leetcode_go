package leetcode_go

func myAtoi(str string) int {
	var tmp = 0
	var i = 0
	var ne = 1
	for ; i < len(str); i++ {
		if str[i] != ' ' {
			break
		}
	}

	if i >= len(str) {
		return 0
	}

	if str[i] == '+' {
		i++
	} else if str[i] == '-' {
		ne = -1
		i++
	}

	for ; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			// 溢出判断
			if tmp > 214748364 {
				if ne > 0 {
					return 2147483647
				} else {
					return -2147483648
				}

			} else if tmp == 214748364 {
				if ne > 0 && str[i] > '7' {
					return 2147483647
				}

				if ne < 0 && str[i] >= '8' {
					return -2147483648
				}
			}

			tmp = tmp*10 + int(str[i]-'0')

		} else {
			break
		}
	}

	return ne * tmp
}
