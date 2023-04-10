package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func candy(ratings []int) int {
	if len(ratings) < 2 {
		return len(ratings)
	}

	cn := make([]int, len(ratings))

	for i := 0; i < len(cn); i++ {
		cn[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			cn[i] = cn[i-1] + 1
		}
	}

	r := cn[len(cn)-1]
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			cn[i] = MyMax(cn[i], cn[i+1]+1)
		}
		r += cn[i]
	}

	return r
}
