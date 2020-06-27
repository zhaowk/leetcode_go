package leetcode_go

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	m := len(obstacleGrid)

	if m == 0 {
		return 0
	}

	n := len(obstacleGrid[0])

	if n == 0 {
		return 0
	}

	tmp := make([]int, m)
	tmpNum := 1
	for idx := range tmp {
		if obstacleGrid[idx][0] == 1 {
			tmpNum = 0
		}
		tmp[idx] = tmpNum
	}

	tops := make([]int, n)
	tmpNum = 1
	for idx := range tops {
		if obstacleGrid[0][idx] == 1 {
			tmpNum = 0
		}
		tops[idx] = tmpNum
	}

	for i := 1; i < n; i++ {
		top := tops[i]
		tmp[0] = top
		for j := 1; j < m; j++ {
			if obstacleGrid[j][i] == 1 {
				tmp[j] = 0
				top = 0
			} else {
				tmp[j] = tmp[j] + top
				top = tmp[j]
			}
		}
	}

	return tmp[m-1]

}
