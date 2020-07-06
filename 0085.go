package leetcode_go

func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}

	mx := 0
	dp := make([]int, len(matrix[0]))
	for i, m := range matrix {
		for j := range m {
			if matrix[i][j] == '1' {
				dp[j] = dp[j] + 1
			} else {
				dp[j] = 0
			}
		}
		mx = max(mx, largestRectangleArea(dp))
	}
	return mx
}
