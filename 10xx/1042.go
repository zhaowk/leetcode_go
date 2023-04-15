package _0xx

func gardenNoAdj(n int, paths [][]int) []int {
	maps := make([][]int, n+1)
	for _, path := range paths {
		// 双向图
		maps[path[0]] = append(maps[path[0]], path[1])
		maps[path[1]] = append(maps[path[1]], path[0])
	}

	result := make([]int, n+1)
	for i := 1; i < len(result); i++ {
		tmp := 0
		for _, index := range maps[i] {
			if result[index] > 0 {
				tmp |= 1 << (result[index] - 1)
			}
		}
		result[i] = findOne(tmp)
	}
	return result[1:]
}

func findOne(t int) int {
	if t&1 == 0 {
		return 1
	} else if t&2 == 0 {
		return 2
	} else if t&4 == 0 {
		return 3
	} else {
		return 4
	}
}
