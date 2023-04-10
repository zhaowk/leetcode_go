package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func minCut(s string) int {
	if len(s) < 2 {
		return 0
	}

	t := make([][]bool, len(s))

	for i := 0; i < len(s); i++ {
		t[i] = make([]bool, len(s)+1)
		t[i][i] = false
		for j := i + 1; j <= len(s); j++ {
			t[i][j] = isV(s[i:j])
		}
	}

	res := make([]int, len(s)+1)
	res[0] = 0
	res[1] = 0

	for i := 2; i <= len(s); i++ {
		res[i] = res[i-1] + 1
		if t[0][i] {
			res[i] = 0
		} else {
			for j := 1; j < i; j++ {
				if t[j][i] {
					res[i] = MyMin(res[i], res[j]+1)
				}
			}
		}
	}

	return res[len(s)]
}
