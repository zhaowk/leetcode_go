package leetcode_go

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	mx, mi, pr := prices[0], prices[0], 0

	for _, price := range prices {
		if mi > price {
			mx, mi = price, price
		} else {
			mx = price
			pr = max(pr, mx-mi)
		}
	}

	return pr
}
