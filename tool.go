package _0xx

import "math"

func MyMax(i ...int) int {
	m := math.MinInt32
	for _, v := range i {
		if v > m {
			m = v
		}
	}
	return m
}

func MyMin(i ...int) int {
	m := math.MaxInt32
	for _, v := range i {
		if v < m {
			m = v
		}
	}
	return m
}

func MyAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func FloatEq(a, b, c float64) bool {

	return math.Abs(a-b) < c
}

func FloatEqN5(a, b float64) bool {
	return FloatEq(a, b, 1e-5)
}

func FloatEqN6(a, b float64) bool {
	return FloatEq(a, b, 1e-6)
}
