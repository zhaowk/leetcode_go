package leetcode_go

// 暴力
func maxProfit3(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	mx := 0

	for i := 0; i < len(prices); i++ {
		mx = max(mx, maxProfitOne(prices[0:i])+maxProfitOne(prices[i:]))
	}

	return mx
}

func maxProfitOne(prices []int) int {
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
