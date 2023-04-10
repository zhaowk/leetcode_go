package _1xx

import "math"
import (
	. "github.com/zhaowk/leetcode_go"
)

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon), len(dungeon[0])
	dp := make([][]int, n+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, m+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt32
		}
	}
	dp[n][m-1], dp[n-1][m] = 1, 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			mi := MyMin(dp[i+1][j], dp[i][j+1])
			dp[i][j] = MyMax(mi-dungeon[i][j], 1)
		}
	}
	return dp[0][0]
}
