package leetcode_go

func colorBorder(grid [][]int, row int, col int, color int) [][]int {
	if grid[row][col] == color {
		return grid
	}

	visited := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]int, len(grid[i]))
	}

	pc := grid[row][col]
	var colorBorderInternal func(i, j int)
	colorBorderInternal = func(i, j int) {
		visited[i][j] = 1  // 已访问标记
		if i == 0 || j == 0 || i == len(grid) - 1 || j == len(grid[i]) - 1 ||
			(grid[i-1][j] != pc || grid[i][j-1] != pc || grid[i+1][j] != pc || grid[i][j+1] != pc) {  // 边界
			visited[i][j] = 2  // 预定着色
		}

		if i > 0 && visited[i-1][j] == 0 && grid[i-1][j] == pc {  // up
			colorBorderInternal(i-1, j)
		}

		if j > 0 && visited[i][j-1] == 0 && grid[i][j-1] == pc {  // left
			colorBorderInternal(i, j-1)
		}

		if i < len(grid) - 1 && visited[i+1][j] == 0 && grid[i + 1][j] == pc {  // down
			colorBorderInternal(i + 1, j)
		}

		if j < len(grid[i]) - 1 && visited[i][j+1] == 0 && grid[i][j + 1] == pc { // right
			colorBorderInternal(i, j + 1)
		}
	}

	colorBorderInternal(row, col)
	for i := 0; i < len(visited); i++ {
		for j := 0; j < len(visited[i]); j++ {
			if visited[i][j] == 2 {
				grid[i][j] = color  // 着色
			}
		}
	}
	return grid
}
