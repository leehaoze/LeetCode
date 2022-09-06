package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	for i := 2; i < len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return min(dp[len(dp)-1]+cost[len(dp)-1], dp[len(dp)-2]+cost[len(dp)-2])
}
