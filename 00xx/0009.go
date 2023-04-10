package _0xx

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	y := x
	d := 0
	z := 0
	for y != 0 {
		d = y % 10
		y = y / 10
		z = z*10 + d
	}
	return z == x
}
