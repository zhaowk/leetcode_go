package _0xx

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

func intToRoman(num int) string {
	var s string
	n := num / 1000
	s += romanNum(n, "M", "", "")

	n = num % 1000 / 100
	s += romanNum(n, "C", "D", "M")

	n = num % 100 / 10
	s += romanNum(n, "X", "L", "C")

	n = num % 10
	s += romanNum(n, "I", "V", "X")

	return s
}

func romanNum(n int, ca1, ca2, ca3 string) string {

	if n == 9 {
		return ca1 + ca3
	}
	if n-5 >= 0 {
		return ca2 + romanNum(n-5, ca1, ca2, "")
	}

	if n == 4 {
		return ca1 + ca2
	} else {
		var s string
		for i := 0; i < n; i++ {
			s += ca1
		}
		return s
	}
}
