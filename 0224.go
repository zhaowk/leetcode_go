package leetcode_go

const (
	typeNumber = iota // number
	typeAdd           // +
	typeMinus         // -
	typeStart         // (
	typeEnd           // )
	typeOther         // other
)

type calNode struct {
	typ int
	val int
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	calc := getCalNodes(s)
	q := buildCalcQ(calc, 0, len(calc))
	return calcQ(q)
}

func getCalNodes(s string) []*calNode {
	calc := make([]*calNode, 0)
	idx := make([]int, 0)
	typ := typeOther
	num := 0
	appendCalc := func() {
		if typ == typeNumber {
			if len(calc) > 1 && calc[len(calc)-2].typ == typeNumber && (calc[len(calc)-1].typ == typeAdd || calc[len(calc)-1].typ == typeMinus) {
				target := calc[len(calc)-2].val
				if len(calc) > 2 && calc[len(calc)-3].typ == typeMinus {
					calc[len(calc)-3].typ = typeAdd
					target = -target
				}
				if calc[len(calc)-1].typ == typeAdd {
					target += num
				} else if calc[len(calc)-1].typ == typeMinus {
					target -= num
				}
				calc = calc[:len(calc)-2]
				calc = append(calc, &calNode{typ: typeNumber, val: target})
			} else {
				calc = append(calc, &calNode{typ: typeNumber, val: num})
			}
		}
	}
	for _, c := range s {
		switch c {
		case '(':
			//if typ == typeNumber {
			//	calc = append(calc, &calNode{typ: typ, val: num })
			//}
			appendCalc()
			typ = typeStart
			idx = append(idx, len(calc))
			calc = append(calc, &calNode{typ: typ})
		case ')':
			//if typ == typeNumber {
			//	calc = append(calc, &calNode{typ: typ, val: num })
			//}
			appendCalc()

			typ = typeEnd
			i := idx[len(idx)-1]
			idx = idx[:len(idx)-1]
			calc[i].val = len(calc)
			calc = append(calc, &calNode{typ: typ, val: i})
		case '+':
			//if typ == typeNumber {
			//	calc = append(calc, &calNode{typ: typ, val: num })
			//}
			appendCalc()

			typ = typeAdd
			calc = append(calc, &calNode{typ: typ})
		case '-':
			//if typ == typeNumber {
			//	calc = append(calc, &calNode{typ: typ, val: num })
			//}
			appendCalc()

			if len(calc) == 0 || calc[len(calc)-1].typ == typeStart {
				calc = append(calc, &calNode{typ: typeNumber, val: 0})
				calc = append(calc, &calNode{typ: typeMinus})
			} else {
				typ = typeMinus
				calc = append(calc, &calNode{typ: typ})
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			typ = typeNumber
			num = 10*num + int(c-'0')
		default:
			//if typ == typeNumber {
			//	calc = append(calc, &calNode{typ: typ, val: num })
			//}

			appendCalc()

			typ = typeOther
		}

		if typ != typeNumber {
			num = 0
		}
	}

	if typ == typeNumber {
		calc = append(calc, &calNode{typ: typ, val: num})
	}
	return calc
}

func buildCalcQ(nodes []*calNode, start, end int) []*calNode {
	res := make([]*calNode, 0)

	for start < end {
		start, res = appendExp(res, nodes, start)
	}
	return res
}

func appendExp(res, nodes []*calNode, start int) (int, []*calNode) {
	if start > len(nodes)-1 {
		return start, res
	}
	if nodes[start].typ == typeNumber {
		res = append(res, nodes[start])
		return start + 1, res
	} else if nodes[start].typ == typeStart {
		idx := nodes[start].val
		tmp := buildCalcQ(nodes, start+1, idx)
		res = append(res, tmp...)
		return idx + 1, res
	} else if nodes[start].typ == typeAdd || nodes[start].typ == typeMinus {
		var idx int
		idx, res = appendExp(res, nodes, start+1)
		res = append(res, nodes[start])
		return idx, res
	}
	panic("error")
}

func calcQ(node []*calNode) int {
	if node == nil || len(node) == 0 {
		return 0
	}

	for len(node) > 1 {
		for i := 2; i < len(node); i++ {
			if node[i].typ == typeAdd {
				tmp := node[0 : i-2]
				tmp = append(tmp, &calNode{typ: typeNumber, val: node[i-2].val + node[i-1].val})
				if i+1 < len(node) {
					tmp = append(tmp, node[i+1:]...)
				}
				node = tmp
				break
			} else if node[i].typ == typeMinus {
				tmp := node[0 : i-2]
				tmp = append(tmp, &calNode{typ: typeNumber, val: node[i-2].val - node[i-1].val})
				if i+1 < len(node) {
					tmp = append(tmp, node[i+1:]...)
				}
				node = tmp
				break
			} else if node[i].typ != typeNumber {
				return -1
			}
		}
	}
	return node[0].val
}
