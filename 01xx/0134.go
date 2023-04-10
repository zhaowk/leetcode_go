package _1xx

func canCompleteCircuit(gas []int, cost []int) int {

	if len(gas) == 0 || len(cost) == 0 || len(gas) != len(cost) {
		return -1
	}

	res := -1

	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		curr := gas[i] - cost[i]
		for j := 1; j < len(gas); j++ {
			if curr < 0 {
				break
			}
			pos := (i + j) % len(gas)
			//fmt.Println("before:",pos, curr)
			curr += gas[pos] - cost[pos]
			//fmt.Println("after",pos, curr)
		}

		if curr >= 0 {
			return i
		}
	}
	return res
}
