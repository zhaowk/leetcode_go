package leetcode_go

func getSkyline(buildings [][]int) [][]int {
	res := make([][]int, 0)
	res = append(res, []int{buildings[0][0], buildings[0][2]}, []int{buildings[0][1], 0})
	idx := 0

	for i := 1; i < len(buildings); i++ {
		tmp := res[:idx]
		for j := idx; buildings[i][1] < res[j][0]; j++ {
			if res[j][1] < buildings[i][2] { // higher building

			}
		}

		res = tmp
	}
	return res
}
