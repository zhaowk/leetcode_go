package _1xx

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	profit := 0

	mi, mx := prices[0], prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < mx {
			if mi < mx {
				profit += mx - mi
			}
			mx, mi = prices[i], prices[i]
		} else if prices[i] > mx {
			mx = prices[i]
		}

		if prices[i] < mi {
			mi = prices[i]
		}
	}

	if mi < mx {
		profit += mx - mi
	}

	return profit
}
