package _1xx

import (
	. "github.com/zhaowk/leetcode_go"
)

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
			pr = MyMax(pr, mx-mi)
		}
	}

	return pr
}
