package leetcode_go

const (
	typNum = iota
	typAdd
	typMinus
	typMulti
	typDivide
	typOther
)

type calculator struct {
	typ int
	val int
}

var (
	op = map[int]func(x, y int) int {
		typAdd: func(x, y int) int {return x + y},
		typMinus: func(x, y int) int {return x - y},
		typMulti: func(x, y int) int {return x * y},
		typDivide: func(x, y int) int {return x / y},
	}
)

func calculate2(s string) int {
	tmp := make([]*calculator, 0)
	num := 0
	typ := typOther
	var calcu = func(c *calculator) {
		if typ == typNum {
			if len(tmp) > 1 && (tmp[len(tmp) - 1].typ == typMulti || tmp[len(tmp) - 1].typ == typDivide) {
				// calculate: * /
				r := op[tmp[len(tmp) - 1].typ](tmp[len(tmp) - 2].val, num)
				tmp = tmp[:len(tmp) - 2]
				tmp = append(tmp, &calculator{typ:typNum, val: r})
			} else {
				tmp = append(tmp, &calculator{typ: typNum, val: num})
			}
		}

		if c != nil {
			typ = c.typ

			if typ == typAdd || typ == typMinus {
				// calculate + -
				for len(tmp) > 2 {
					n1, ope, n2 := tmp[0], tmp[1], tmp[2]
					r := op[ope.typ](n1.val, n2.val)
					tmp = tmp[2:]
					tmp[0].val = r
				}
			}

			tmp = append(tmp, c)
		}
	}
	for _, c := range s {
		switch c {
		case '+':
			calcu(&calculator{typ: typAdd})
		case '-':
			calcu(&calculator{typ: typMinus})
		case '*':
			calcu(&calculator{typ: typMulti})
		case '/':
			calcu(&calculator{typ: typDivide})
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			typ = typNum
			num = 10 * num + int(c - '0')
		default:
			calcu(nil)
			typ = typOther
		}

		if typ != typNum {
			num = 0
		}
	}

	if typ == typNum {
		calcu(nil)
	}

	for len(tmp) > 2 {
		n1, ope, n2 := tmp[0], tmp[1], tmp[2]
		r := op[ope.typ](n1.val, n2.val)
		tmp = tmp[2:]
		tmp[0].val = r
	}

	if len(tmp) == 1 && tmp[0].typ == typNum {
		return tmp[0].val
	}
	return 0
}
