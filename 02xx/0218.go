package _2xx

// TODO
func getSkyline(buildings [][]int) [][]int {
	const (
		start     = 0
		end       = 1
		height    = 2
		resHeight = 1
	)

	var checkAndAppend = func(res [][]int, data []int) [][]int {
		if len(res) == 0 || res[len(res)-1][resHeight] != data[resHeight] {
			return append(res, data)
		}
		return res
	}

	res := make([][]int, 0)
	// first: {start, height}, {end, 0}
	res = append(res, []int{buildings[0][start], buildings[0][height]}, []int{buildings[0][end], 0})
	idx := 0

	for i := 1; i < len(buildings); i++ {
		tmp := make([][]int, idx)
		copy(tmp, res[:idx])
		for j := idx; j < len(res); j++ {
			next := 0
			if j == len(res)-1 { // last
				next = buildings[i][end]
			} else {
				next = res[j+1][start]
				if res[j+1][start] < buildings[i][start] {
					tmp = append(tmp, res[j])
					idx++
					continue
				}
			}

			if buildings[i][height] > res[j][resHeight] { // higher
				if res[j][start] < buildings[i][start] {
					tmp = checkAndAppend(tmp, []int{res[j][start], res[j][resHeight]})
					tmp = checkAndAppend(tmp, []int{buildings[i][start], buildings[i][height]})
				} else {
					tmp = checkAndAppend(tmp, []int{res[j][start], buildings[i][height]})
				}

				if buildings[i][end] < next {
					tmp = checkAndAppend(tmp, []int{buildings[i][end], res[j][resHeight]})
				}
			} else { // lower
				tmp = checkAndAppend(tmp, res[j])
			}

			if buildings[i][end] <= next {
				if j == len(res)-1 { // last
					if tmp[len(tmp)-1][resHeight] != 0 {
						tmp = checkAndAppend(tmp, []int{buildings[i][end], 0})
					}
				} else {
					tmp = append(tmp, res[j+1:]...)
				}

				for idx < len(tmp) {
					if tmp[idx][start] < buildings[i][start] {
						idx++
					} else {
						break
					}
				}
				break
			}
		}

		res = tmp
	}
	return res
}
