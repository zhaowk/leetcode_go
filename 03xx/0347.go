package _3xx

import "sort"

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	x := make([]struct{ num, cnt int }, 0)

	for _, n := range nums {
		m[n]++
	}

	for i, v := range m {
		x = append(x, struct{ num, cnt int }{num: i, cnt: v})
	}

	sort.Slice(x, func(i, j int) bool {
		return x[i].cnt > x[j].cnt
	})

	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = x[i].num
	}
	return res
}
