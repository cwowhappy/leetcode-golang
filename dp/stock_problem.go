package dp

func Problem121(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0
	maxBuyProfit := -prices[0]

	for i := 1; i < len(prices); i ++ {
		if maxProfit < prices[i] + maxBuyProfit {
			maxProfit = prices[i] + maxBuyProfit
		}
		if maxBuyProfit < -prices[i] {
			maxBuyProfit = -prices[i]
		}
	}

	return maxProfit
}

func Problem122(prices []int) int {
	maxProfit := 0

	if 0 < len(prices) {
		curMinPrice := prices[0]
		curMaxProfit := 0
		for i := 1; i < len(prices); i ++ {
			if curMaxProfit < prices[i] - curMinPrice {
				curMaxProfit = prices[i] - curMinPrice
			} else {
				maxProfit += curMaxProfit
				curMaxProfit = 0
				curMinPrice = prices[i]
			}
		}
	}

	return maxProfit
}