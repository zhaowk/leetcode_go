package leetcode_go

func isPathCrossing(path string) bool {

	lp := len(path)

	if lp < 2 {
		return false
	}

	paths := make([][]int, 0)
	curr := []int{0, 0}
	paths = append(paths, []int{0, 0})
	for _, p := range path {
		switch p {
		case 'N':
			curr[0]--
		case 'S':
			curr[0]++
		case 'E':
			curr[1]--
		case 'W':
			curr[1]++
		}
		for _, pa := range paths {
			if len(pa) == 2 && pa[0] == curr[0] && pa[1] == curr[1] {
				return true
			}
		}

		paths = append(paths, []int{curr[0], curr[1]})
	}
	return false
}
