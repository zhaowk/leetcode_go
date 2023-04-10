package _1xx

func titleToNumber(s string) int {
	c := 0

	for _, v := range s {
		c = c*26 + int(v) - '@'
	}
	return c
}
