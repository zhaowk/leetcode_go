package _2xx

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	vi := make([][]bool, len(grid))
	for i := 0; i < len(vi); i++ {
		vi[i] = make([]bool, len(grid[i]))
	}

	num := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && !vi[i][j] {
				num++
				//
				run(grid, vi, i, j)
			}
		}
	}

	return num
}

func run(grid [][]byte, vi [][]bool, i, j int) {
	vi[i][j] = true
	if i+1 < len(grid) && !vi[i+1][j] && grid[i+1][j] == '1' {
		run(grid, vi, i+1, j)
	}

	if i-1 >= 0 && !vi[i-1][j] && grid[i-1][j] == '1' {
		run(grid, vi, i-1, j)
	}

	if j+1 < len(grid[i]) && !vi[i][j+1] && grid[i][j+1] == '1' {
		run(grid, vi, i, j+1)
	}

	if j-1 >= 0 && !vi[i][j-1] && grid[i][j-1] == '1' {
		run(grid, vi, i, j-1)
	}
}
