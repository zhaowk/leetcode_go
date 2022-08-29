package leetcode_go

import (
	"fmt"
)

type pos struct {
	a, b int
}

type hp struct {
	data []pos
	comp func(a ,b pos) bool
}

func (h *hp)up() {
	if len(h.data) < 2 {
		return
	}

	idx := len(h.data) - 1

	for idx > 0 {
		if h.comp(h.data[idx], h.data[(idx-1)/2]) {
			h.data[idx], h.data[(idx-1)/2] = h.data[(idx-1)/2], h.data[idx]
			idx = (idx-1)/2
		} else {
			break
		}
	}
}

func (h *hp)down(idx int) {
	for idx * 2 + 1 < len(h.data) {
		if idx * 2 + 2 < len(h.data) {
			if h.comp(h.data[idx * 2 + 1], h.data[idx]) || h.comp(h.data[idx * 2 + 2], h.data[idx]) {
				if h.comp(h.data[idx * 2 + 1], h.data[idx * 2 + 2]) {
					h.data[idx * 2 + 1], h.data[idx] = h.data[idx], h.data[idx * 2 + 1]
					idx = idx * 2 + 1
				} else {
					h.data[idx * 2 + 2], h.data[idx] = h.data[idx], h.data[idx * 2 + 2]
					idx = idx * 2 + 2
				}
			} else {
				break
			}
		} else {
			if h.comp(h.data[idx * 2 + 1], h.data[idx]) {
				h.data[idx * 2 + 1], h.data[idx] = h.data[idx], h.data[idx * 2 + 1]
				idx = idx * 2 + 1
			} else {
				break
			}
		}
	}
}

func (h *hp)Insert(p pos) {
	if len(h.data) == 0 {
		h.data = []pos{p}
	} else {
		h.data = append(h.data, p)
		h.up()
	}
}

func (h *hp)Pop() pos {
	if len(h.data) == 0 {
		return pos{}
	}

	res := h.data[0]
	h.data[0] = h.data[len(h.data) - 1]
	h.data = h.data[:len(h.data) - 1]
	h.down(0)
	return res
}

func (h *hp)Range(f func(i int, p pos)bool) {
	tmp := &hp{
		comp: h.comp,
	}
	tmp.data = make([]pos, len(h.data))

	copy(tmp.data, h.data)

	for i:= 0; len(tmp.data) > 0; i++ {
		if !f(i, tmp.Pop()) {
			break
		}
	}
}

func (h *hp)Valid(comp func(a, b pos)bool) bool {
	tmp := []int{0}

	for i:= 0; len(tmp) > 0; i++ {
		if tmp[0] * 2 + 1 < len(h.data) {
			if !comp(h.data[tmp[0]], h.data[tmp[0] * 2 + 1]) {
				fmt.Println("1", tmp[0])
				return false
			}
			tmp = append(tmp, tmp[0] * 2 + 1)
		}

		if tmp[0] * 2 + 2 < len(h.data) {
			if !comp(h.data[tmp[0]], h.data[tmp[0] * 2 + 2]) {
				fmt.Println("2", tmp[0])

				return false
			}
			tmp = append(tmp, tmp[0] * 2 + 2)
		}

		if len(tmp) > 0 {
			tmp = tmp[1:]
		}
	}

	return true
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	r1, r2 := 0, len(arr) - 1
	if k == 1 {
		return []int{arr[r1], arr[r2]}
	}

	h := &hp {
		comp: func(a, b pos) bool {
			return arr[a.a] * arr[b.b] < arr[a.b] * arr[b.a]
		},
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			h.Insert(pos{i, j})
		}
	}

	fmt.Println("valid", h.Valid(func(a, b pos) bool {
		return arr[a.a] * arr[b.b] < arr[a.b] * arr[b.a]
	}))

	h.Range(func(i int, p pos) bool {
		fmt.Println(i, p.a, p.b, "=>", arr[p.a], "/", arr[p.b] , float64(arr[p.a]) / float64(arr[p.b]))
		return true
	})

	p := pos{}

	for i := 0; i < k; i++ {
		p = h.Pop()
	}

	return []int{arr[p.a], arr[p.b]}
}
