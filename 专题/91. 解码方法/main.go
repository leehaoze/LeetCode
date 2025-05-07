package main

import (
	"fmt"
)

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] == '0' {
		dp[1] = 0
	} else {
		dp[1] = 1
	}

	for i := 1; i < len(s); i++ {
		if s[i] == '0' && s[i-1] == '0' {
			return 0
		} else if s[i] == '0' {
			if !canCombine(s, i) {
				return 0
			}
			dp[i+1] = dp[i-1]
		} else if canCombine(s, i) {
			dp[i+1] = dp[i] + dp[i-1]
		} else {
			dp[i+1] = dp[i]
		}
	}

	// fmt.Println(dp)

	return dp[len(s)]
}

func canCombine(s string, i int) bool {
	if s[i-1] == '0' {
		return false
	}

	if s[i-1] == '1' {
		return true
	}

	if s[i-1] == '2' && s[i] < '7' {
		return true
	}

	return false
}

func main() {
	fmt.Println(numDecodings("230"))
}
