package _0xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func longestValidParentheses(s string) int {

	mx := 0
	ls := len(s)
	start, end := 0, ls-1

	for start < ls && s[start] != '(' {
		start++
	}

	for end >= 0 && s[end] != ')' {
		end--
	}

	for i := start; i <= end; i++ {
		left, right := 0, 0
		for j := i; j <= end; j++ {
			switch s[j] {
			case '(':
				left++
			case ')':
				right++
			}
			if left == right {
				mx = MyMax(mx, j-i+1)
			} else if left < right {
				break
			}
		}
	}
	return mx
}
