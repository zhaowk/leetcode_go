package leetcode_go

func maxProfit1(k int, prices []int) int {
	if len(prices) < 2 || k < 1 {
		return 0
	}

	if k > len(prices)/2 {
		dp_s := make([]int, len(prices))
		dp_m := make([]int, len(prices))
		dp_s[0] = 0
		dp_m[0] = -prices[0]
		for i := 1; i < len(prices); i++ {
			dp_s[i] = max(dp_s[i-1], dp_m[i-1]+prices[i])
			dp_m[i] = max(dp_m[i-1], dp_s[i-1]-prices[i])
		}
		return dp_s[len(prices)-1]
	}

	dp_s := make([][]int, len(prices))
	dp_m := make([][]int, len(prices))

	for i := 0; i < len(dp_s); i++ {
		dp_s[i] = make([]int, k+1)
		dp_m[i] = make([]int, k+1)
	}

	for i := 1; i <= k; i++ {
		dp_s[0][i] = 0
		dp_m[0][i] = -prices[0]
	}

	for i := 1; i < len(prices); i++ {
		for j := k; j >= 1; j-- {
			dp_s[i][j] = max(dp_s[i-1][j], dp_m[i-1][j]+prices[i])
			dp_m[i][j] = max(dp_m[i-1][j], dp_s[i-1][j-1]-prices[i])
		}
	}

	return dp_s[len(prices)-1][k]
}
