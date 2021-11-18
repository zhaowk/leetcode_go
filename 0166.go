package leetcode_go

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	neg := false

	if numerator < 0 {
		neg = !neg
		numerator = -numerator
	} else if numerator == 0 {
		return "0"
	}

	if denominator < 0 {
		neg = !neg
		denominator = -denominator
	} else if denominator == 0 {
		return ""
	}

	ns := ""
	if neg {
		ns = "-"
	}

	if numerator%denominator == 0 {
		return fmt.Sprintf("%s%d", ns, numerator/denominator)
	}

	result := fmt.Sprintf("%s%d.", ns, numerator/denominator)
	l := numerator % denominator

	a := make([]int, 0)
	b := make([]int, 0)

	for l != 0 {
		a = append(a, l)
		b = append(b, l*10/denominator)
		l = l * 10 % denominator
		if l == 0 {
			for _, v := range b {
				result += strconv.Itoa(v)
			}
			return result
		}

		for idx, v := range a {
			if l == v {
				for i := 0; i < idx; i++ {
					result += strconv.Itoa(b[i])
				}
				result += "("
				for i := idx; i < len(b); i++ {
					result += strconv.Itoa(b[i])
				}
				result += ")"
				return result
			}
		}
	}

	return result
}
