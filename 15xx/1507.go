package _5xx

import (
	"fmt"
	"strconv"
	"strings"
)

func reformatDate(date string) string {
	d := strings.Split(date, " ")

	var m = map[string]int{
		"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4,
		"May": 5, "Jun": 6, "Jul": 7, "Aug": 8,
		"Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,
	}

	if len(d) != 3 {
		return ""
	}

	day := myAtoi1(d[0])

	mon := m[d[1]]

	year, _ := strconv.Atoi(d[2])

	return fmt.Sprintf("%4d-%02d-%02d", year, mon, day)
}

func myAtoi1(s string) int {
	r := 0
	if len(s) > 0 {
		r = int(s[0]) - '0'
	}
	if len(s) > 1 && s[1] >= '0' && s[1] <= '9' {
		r = 10*r + int(s[1]) - '0'
	}
	return r
}
