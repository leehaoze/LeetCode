package main

func finalPrices(prices []int) []int {
	ret := make([]int, len(prices))
	for i := 0; i < len(prices); i++ {
		price := prices[i]
		for j := i + 1; j < len(prices); j++ {
			if prices[i] >= prices[j] {
				price = price - prices[j]
				break
			}
		}
		ret[i] = price
	}

	return ret
}
