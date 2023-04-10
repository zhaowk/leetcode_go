package _4xx

import (
	"bytes"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func addStrings(num1 string, num2 string) string {
	la, lb := len(num1), len(num2)
	overflow, current := 0, 0

	buf := bytes.Buffer{}

	lastIndex := MyMax(la, lb)

	for i := 1; i <= lastIndex; i++ {
		ca, cb := 0, 0
		if la-i >= 0 {
			ca = int(num1[la-i] - '0')
		}

		if lb-i >= 0 {
			cb = int(num2[lb-i] - '0')
		}

		current = ca + cb + overflow
		overflow = current / 10
		current = current % 10
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
