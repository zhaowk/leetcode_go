package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := make(map[int]int, 0)

	for _, v := range nums {
		m[v] = 1
	}

	mx := 1

	for _, v := range nums {
		if m[v] != 0 {
			res := 1
			i := 1
			for ; m[v+i] == 1; i++ {
				delete(m, v+i)
			}
			res += i - 1
			i = 1
			for ; m[v-i] == 1; i++ {
				delete(m, v-i)
			}
			res += i - 1
			mx = MyMax(mx, res)
			delete(m, v)
		}

		if len(m) == 0 {
			break
		}

		//fmt.Println(m)
	}
	return mx
}
