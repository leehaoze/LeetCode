package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	dpPre := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		dpPre = max(dpPre, prices[i]-minPrice)
		minPrice = min(minPrice, prices[i])
	}

	return dpPre
}
