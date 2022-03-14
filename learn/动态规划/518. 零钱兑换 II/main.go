package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for _, coin := range coins {
		dp[coin] = 1
	}

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin <= 0 {
				continue
			}
			dp[i] += dp[coin]
		}
	}

	// for _, d := range dp {
	// 	fmt.Println(d)
	// }
	return dp[amount]
}

func main() {
	// fmt.Println(change(5, []int{1, 2, 5}))
	fmt.Println(change(3, []int{2}))
}
