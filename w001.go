package leetcode_go

func average(salary []int) float64 {
	ls := len(salary)
	if ls < 3 {
		return 0
	}

	minimum, maximum, total := salary[0], salary[0], 0

	for _, value := range salary {
		if value < minimum {
			minimum = value
		}

		if value > maximum {
			maximum = value
		}

		total += value
	}

	return float64(total-maximum-minimum) / float64(ls-2)
}
