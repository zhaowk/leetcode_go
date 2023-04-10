package _0xx

import (
	"unicode"
)
import (
	. "github.com/zhaowk/leetcode_go"
)

func recoverFromPreorder(S string) *TreeNode {
	root := &TreeNode{0, nil, nil}

	tmp := make([]*TreeNode, 1)

	for len(S) > 0 {
		value, depth, pos := getOneNode(S)
		//fmt.Println(value, depth, pos, S)
		if depth == 0 {
			root.Val = value
			tmp[depth] = root
		} else {
			item := &TreeNode{value, nil, nil}
			if depth >= len(tmp) {
				tmp = append(tmp, item)
			} else {
				tmp[depth] = item
			}
			if nil == tmp[depth-1].Left {
				tmp[depth-1].Left = item
			} else if nil == tmp[depth-1].Right {
				tmp[depth-1].Right = item
			}
		}
		S = S[pos:]
	}

	return root
}

func getOneNode(s string) (value, depth, pos int) {
	depth = 0
	pos = 0
	value = 0
	for ; s[pos] == '-'; pos++ {
		depth++
	}

	for ; pos < len(s) && unicode.IsDigit(rune(s[pos])); pos++ {
		value = value*10 + int(s[pos]-'0')
	}
	return
}
