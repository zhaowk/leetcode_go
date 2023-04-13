package _4xx

func mostFrequentEven(nums []int) int {
	target, mx := -1, 0
	m := make(map[int]int)

	for _, num := range nums {
		if (num & 1) == 0 {
			m[num]++
			if m[num] > mx {
				target = num
				mx = m[num]
			} else if m[num] == mx && num < target {
				target = num
			}
		}
	}

	return target
}
