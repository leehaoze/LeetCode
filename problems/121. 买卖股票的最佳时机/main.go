package main

import "fmt"

/*
dp[i] = 第i天卖出时的最大收益
dp[i] = max(dp[i-1], prices[i] - dp[i-1] + max)

dp[1] = 0
dp[2] = 0
dp[3] = 5, max = 5
*/

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	minPrices := prices[0]
	for i := 1; i < len(prices); i++ {
		minPrices = min(minPrices, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrices)
	}

	fmt.Println(dp)

	return dp[len(prices)-1]
}

func main() {
	fmt.Println(maxProfit([]int{2, 1, 2, 1, 0, 1, 2}))
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2}))
}
