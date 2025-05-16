package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices {
		if price < minPrice {
			// 今天买入
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// 更新最大利润
			maxProfit = profit
		}
	}

	return maxProfit
}
