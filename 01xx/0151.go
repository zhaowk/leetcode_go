package _1xx

import (
	"strings"
)

func reverseWords(s string) string {

	a := make([]string, 0)

	start, end := 0, 0

	for start < len(s) && end <= len(s) {
		for start = end; start < len(s) && s[start] == ' '; start++ {
		}

		for end = start; end < len(s) && s[end] != ' '; end++ {
		}

		if start < end {
			a = append(a, s[start:end])
		}
	}

	for i := 0; i < len(a)/2; i++ {
		a[i], a[len(a)-1-i] = a[len(a)-1-i], a[i]
	}

	return strings.Join(a, " ")
}
