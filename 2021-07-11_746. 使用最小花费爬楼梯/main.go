package main

import "fmt"

// dp[n] = min(dp[n-1] + value[n-1],  dp[n-2] + value[n-2])

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}

	a, b := 0, 0

	for i := 2; i <= len(cost); i++ {
		a, b = b, min(a+cost[i-2], b+cost[i-1])
	}

	return b
}

func main() {
	// fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
}
