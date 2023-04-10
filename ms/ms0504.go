package ms

import "math"

func findClosedNumbers(num int) []int {
	// 下限
	if num == 1 {
		return []int{2, -1}
	}

	if num == 2147483647 {
		return []int{-1, -1}
	}

	large, small := -1, -1
	lPos, num1, num0 := 0, 0, 0

	for i := 1; i < math.MaxUint32; i <<= 1 {
		if num&i == 0 {
			num0++
			if lPos > 0 && large < 0 {
				large = setLarge(num, i, num0)
			}

		} else {
			num1++
			if num0 > 0 && small < 0 {
				small = setSmall(num, i, num1)
			}
			lPos = i
		}
		if large > 0 && small > 0 {
			break
		}
	}

	if large > math.MaxInt32 {
		large = -1
	}

	return []int{large, small}
}

func setLarge(num, i, num0 int) int {
	num = num | i
	for j := i; j > 0; j >>= 1 {
		num = num | j
	}

	for j := 1; j <= num0; j++ {
		num &= ^(i >> j)
	}

	return num
}

func setSmall(num, i, num1 int) int {
	num = num & ^i
	for j := i; j > 0; j >>= 1 {
		num = num & ^j
	}

	for j := 1; j <= num1; j++ {
		num |= i >> j
	}

	return num
}
