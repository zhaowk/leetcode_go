package leetcode_go

//I             1
//V             5
//X             10
//L             50
//C             100
//D             500
//M             1000

// I X C

func romanToInt(s string) int {
	p := s
	n := 0
	for len(p) > 0 {
		switch p[0] {
		case 'M':
			n += 1000
			p = p[1:]
		case 'D':
			n += 500
			p = p[1:]
		case 'C':
			if len(p) > 1 {
				if p[1] == 'D' {
					n += 400
					p = p[2:]
				} else if p[1] == 'M' {
					n += 900
					p = p[2:]
				} else {
					n += 100
					p = p[1:]
				}
			} else {
				n += 100
				p = p[1:]
			}
		case 'L':
			n += 50
			p = p[1:]
		case 'X':
			if len(p) > 1 {
				if p[1] == 'L' {
					n += 40
					p = p[2:]
				} else if p[1] == 'C' {
					n += 90
					p = p[2:]
				} else {
					n += 10
					p = p[1:]
				}
			} else {
				n += 10
				p = p[1:]
			}
		case 'V':
			n += 5
			p = p[1:]
		case 'I':
			if len(p) > 1 {
				if p[1] == 'V' {
					n += 4
					p = p[2:]
				} else if p[1] == 'X' {
					n += 9
					p = p[2:]
				} else {
					n += 1
					p = p[1:]
				}
			} else {
				n += 1
				p = p[1:]
			}
		}
	}
	return n
}
