package main

func coinChange(coins []int, amount int) int {
	//memo := make([]int, amount+1)
	//
	//return dp(amount, coins, memo)

	dp := make([]int, amount+1)
	// base case
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i < amount+1; i++ {
		for _, coin := range coins {
			// 子问题误解
			if i-coin < 0 {
				continue
			}

			if dp[i] > 1+dp[i-coin] {
				dp[i] = 1 + dp[i-coin]
			}
		}
	}

	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func dp(n int, coins []int, memo []int) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return - 1
	}

	if memo[n] != 0 {
		return memo[n]
	}

	minVal := 1<<63 - 1
	for _, coin := range coins {
		temp := dp(n-coin, coins, memo)
		if temp == -1 {
			continue
		}

		if 1+temp < minVal {
			minVal = 1 + temp
		}
	}

	if minVal != 1<<63-1 {
		memo[n] = minVal
	} else {
		memo[n] = -1
	}

	return memo[n]
}
