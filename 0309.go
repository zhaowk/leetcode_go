package leetcode_go

func maxProfit4(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	s, b, p := 0, -prices[0], 0

	for i := 0; i < len(prices); i++ {
		s, b, p = max(s, b+prices[i]), max(b, p-prices[i]), s
	}

	return s
}
