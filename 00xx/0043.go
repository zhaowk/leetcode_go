package _0xx

func multiply(num1 string, num2 string) string {

	ln1, ln2 := len(num1), len(num2)

	if ln1 == 0 || ln2 == 0 {
		return "0"
	}

	over, curr := 0, 0
	idx := 0
	num := make([]byte, 1)
	num[0] = '0'

	for i := ln1 - 1; i >= 0; i-- {
		idx = ln1 - 1 - i
		for j := ln2 - 1; j >= 0; j-- {
			over, curr = mul(int(num1[i]-'0'), int(num2[j]-'0'), int(num[idx]-'0'))
			num[idx] = byte(curr + '0')
			idx++
			if idx < len(num) {
				num[idx] = byte(over) + num[idx]
			} else {
				num = append(num, byte(over+'0'))
			}
		}
	}

	reverse(num)
	for len(num) > 1 && num[0] == '0' {
		num = num[1:]
	}
	return string(num)
}

func reverse(s []byte) {
	nl := len(s)
	for i := 0; i < nl/2; i++ {
		s[i], s[nl-1-i] = s[nl-1-i], s[i]
	}
}

func mul(x, y, z int) (int, int) {
	return (x*y + z) / 10, (x*y + z) % 10
}
