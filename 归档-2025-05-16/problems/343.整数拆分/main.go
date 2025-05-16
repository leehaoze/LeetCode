package main

import "fmt"

// 正整数n，拆解为 k 个数，使其乘积最大

// dp[i] 为 正整数 i 对应的 最大乘积
// dp[i] = 最大乘积
/*

max(i, dp[i]) * max(j, dp[j]), i = [1, n/2]

dp[1] = 0
dp[2] = 1 * 1 = 1
dp[3] = 1 * 2 = 2
dp[4] = max(1 * 3, 2 * 2) = 4
dp[5] = max(1 * 4, 2 * 3) = 6
dp[6] = max(1 * 5, 2 * 4, 3 * 3) = 9
dp[7] = max(1 * 6, 2 * 6, 3 * 4) = 12
dp[8] = max(1 * 12, 2 * 9, 3 * 6, 4 * 4) = 18
dp[9] = max(1 * 18, 2 * 12, 3 * 9, 4 * 6) = 27
dp[10] = max(1 * 27, 2 * 18, 3 * 12, 4 * 9, 6 * 6) = 36
*/

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func integerBreak(n int) int {
	dp := make([]int, n+1)

	dp[1] = 0

	for dpIdx := 2; dpIdx <= n; dpIdx++ {
		for i := 1; i <= dpIdx/2; i++ {
			dp[dpIdx] = max(max(i, dp[i])*max(dpIdx-i, dp[dpIdx-i]), dp[dpIdx])
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(integerBreak(10))
}
