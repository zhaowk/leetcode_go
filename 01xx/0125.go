package _1xx

import "unicode"

func isPalindromeStr(s string) bool {

	p, q := 0, len(s)-1

	for p <= q {
		for p < len(s) && !unicode.IsDigit(rune(s[p])) && !unicode.IsLetter(rune(s[p])) {
			p++
		}
		for q >= 0 && !unicode.IsDigit(rune(s[q])) && !unicode.IsLetter(rune(s[q])) {
			q--
		}
		if p < len(s) && q >= 0 && unicode.ToLower(rune(s[p])) != unicode.ToLower(rune(s[q])) {
			return false
		}
		p++
		q--
	}
	return true
}
