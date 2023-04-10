package _3xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func integerReplacement(n int) int {
	return integerReplacementX(n, 0)
}

func integerReplacementX(n, k int) int {
	if n == 1 {
		return k
	}

	if n&1 == 0 {
		return integerReplacementX(n/2, k+1)
	} else {
		return MyMin(integerReplacementX(n+1, k+1), integerReplacementX(n-1, k+1))
	}
}
