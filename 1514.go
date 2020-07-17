package leetcode_go

import (
	"reflect"
)

func isHeap(a interface{}, l int, less func(i, j int) bool) bool {
	rl := reflect.ValueOf(a)
	left, right := 2*l+1, 2*l+2
	if left < rl.Len() && less(l, left) {
		return false
	} else if right < rl.Len() && less(l, right) {
		return false
	} else {
		if left >= rl.Len() {
			return true
		}
		if right >= rl.Len() {
			return isHeap(a, left, less)
		}
		return isHeap(a, left, less) && isHeap(a, right, less)
	}
}

func heapAdd(slice interface{}, item interface{}, less func(int, int, interface{}) bool) interface{} {

	rl := reflect.Append(reflect.ValueOf(slice), reflect.ValueOf(item))

	res := rl.Interface()
	l := rl.Len() - 1
	m := (l - 1) / 2

	for l > 0 {
		if less(l, m, res) {
			reflect.Swapper(res)(l, m)
			l, m = m, (m-1)/2
		} else {
			break
		}
	}
	return res
}

func heapDel(x interface{}, less func(i, j int, r interface{}) bool) (interface{}, interface{}) {
	rl := reflect.ValueOf(x)
	if rl.Len() == 0 {
		return x, 0
	}
	l := 0

	item := rl.Index(0)

	rl = rl.Slice(1, rl.Len())
	res := rl.Interface()
	length := rl.Len()

	for {
		left, right := 2*l+1, 2*l+2
		if right < length {
			if less(l, left, res) && less(l, right, res) {
				break
			} else if less(left, right, res) {
				reflect.Swapper(res)(l, left)
				l = left
			} else if !less(left, right, res) {
				reflect.Swapper(res)(l, right)
				l = right
			} else {
				break
			}
		} else if left < length && less(left, l, res) {
			reflect.Swapper(res)(l, left)
			break
		} else {
			break
		}
	}
	return res, item.Interface()
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	type h struct {
		Prob  float64
		Point int
	}

	up := func(x []h) {
		if len(x) == 0 {
			return
		}
		l := len(x) - 1
		m := (l - 1) / 2

		for l > 0 {
			if x[l].Prob > x[m].Prob {
				x[l], x[m] = x[m], x[l]
				l, m = m, (m-1)/2
			} else {
				break
			}
		}
	}
	down := func(x []h) {
		if len(x) == 0 {
			return
		}
		l := 0

		for {
			left, right := 2*l+1, 2*l+2
			if right < len(x) {
				if x[l].Prob > x[left].Prob && x[l].Prob > x[right].Prob {
					break
				} else if x[left].Prob > x[right].Prob {
					x[l], x[left] = x[left], x[l]
					l = left
				} else if x[left].Prob <= x[right].Prob {
					x[l], x[right] = x[right], x[l]
					l = right
				} else {
					break
				}
			} else if left < len(x) && x[l].Prob < x[left].Prob {
				x[l], x[left] = x[left], x[l]
				break
			} else {
				break
			}
		}
	}

	e := make([][]h, n+1)

	for i := 0; i < len(edges); i++ {
		e[edges[i][0]] = append(e[edges[i][0]], h{succProb[i], edges[i][1]})
		e[edges[i][1]] = append(e[edges[i][1]], h{succProb[i], edges[i][0]})
	}

	m := make(map[int]bool, 0)
	q := make([]h, 0)
	q = append(q, h{1.0, start})
	for len(q) > 0 {

		x := q[0]
		q[0] = q[len(q)-1]
		q = q[:len(q)-1]
		down(q)

		if x.Point == end {
			return x.Prob
		}

		if !m[x.Point] {
			m[x.Point] = true
			for _, y := range e[x.Point] {
				if !m[y.Point] {
					y.Prob *= x.Prob
					q = append(q, y)
					up(q)
				}
			}
		}
	}
	return 0
}
