package leetcode_go

func findMinStep(board string, hand string) int {
	ball := map[rune]int{}
	boardB := map[rune]int{}
	for _, b := range hand {
		ball[b]++
	}

	for _, b := range board {
		boardB[b]++
	}

	for k := range boardB {
		if boardB[k]+ball[k] < 3 {
			return -1
		}
	}

	if len(board) == 1 {
		if ball[rune(board[0])] > 1 {
			return 2
		} else {
			return -1
		}
	}

	return findStep1([]byte(board), ball, 0)
}

func minOrInvalid(prev, curr int) int {
	if curr == -1 {
		return prev
	}

	if prev == -1 {
		return curr
	}
	return min(prev, curr)
}

func findStep1(b []byte, m map[rune]int, s int) int {
	r := -1
	for i := 1; i < len(b); i++ {
		if r > 0 && s >= r {
			return r
		}

		for k, v := range m {
			if v > 0 {
				m[k]--
				tmp := make([]byte, i)
				copy(tmp, b[:i])
				tmp = append(tmp, byte(k))
				r = minOrInvalid(r, findStep2(append(tmp, b[i:]...), m, s+1))
				m[k]++
			}
		}
	}

	return r
}

func findStep2(board []byte, m map[rune]int, s int) int {
	// doMerge
	target := make([]byte, 0)
	b := make([]byte, len(board))
	copy(b, board)
	for {
		target = make([]byte, 0)
		prevItem, prevCnt := b[0], 1
		for i := 1; i < len(b); i++ {
			if b[i] == prevItem {
				prevCnt++
			} else {
				if prevCnt < 3 {
					for j := 0; j < prevCnt; j++ {
						target = append(target, prevItem)
					}
				} else {
					// skip
				}
				prevItem, prevCnt = b[i], 1
			}
		}

		if prevCnt < 3 {
			for j := 0; j < prevCnt; j++ {
				target = append(target, prevItem)
			}
		}
		if len(target) == len(b) {
			break
		} else if len(target) == 0 {
			return s
		}
		b = target
	}

	if len(target) == 0 {
		return s
	}

	boardB := map[rune]int{}

	for _, i := range b {
		boardB[rune(i)]++
	}

	for k := range boardB {
		if boardB[k]+m[k] < 3 {
			return -1
		}
	}

	return findStep1(target, m, s)
}
