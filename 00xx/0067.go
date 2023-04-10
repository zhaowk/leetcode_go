package _0xx

import (
	"bytes"
)

func addBinary(a string, b string) string {
	la, lb := len(a), len(b)
	overflow, current := 0, 0

	buf := bytes.Buffer{}

	lastIndex := maxL(la, lb)

	for i := 1; i <= lastIndex; i++ {
		ca, cb := 0, 0
		if la-i >= 0 {
			ca = int(a[la-i] - '0')
		}

		if lb-i >= 0 {
			cb = int(b[lb-i] - '0')
		}

		current = ca + cb + overflow
		overflow = (2 & current) >> 1
		current = 1 & current
		buf.WriteByte(byte(current + '0'))
	}

	if overflow == 1 {
		buf.WriteByte('1')
	}

	tmp := buf.Bytes()
	lt := len(tmp)

	for i := 0; i < lt/2; i++ {
		tmp[i], tmp[lt-1-i] = tmp[lt-1-i], tmp[i]
	}

	return string(tmp)
}

func maxL(x, y int) int {
	if x < y {
		return y
	}
	return x
}
