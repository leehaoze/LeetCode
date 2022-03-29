package main

// var memo [][]int

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	memo = make([][]int, len(text1))
// 	for i := 0; i < len(text1); i++ {
// 		memo[i] = make([]int, len(text2))
// 		for j := 0; j < len(text2); j++ {
// 			memo[i][j] = -1
// 		}
// 	}

// 	return dp(text1, 0, text2, 0)
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }

// func dp(s1 string, i int, s2 string, j int) int {
// 	if i == len(s1) || j == len(s2) {
// 		return 0
// 	}

// 	if memo[i][j] != -1 {
// 		return memo[i][j]
// 	}

// 	if s1[i] == s2[j] {
// 		memo[i][j] = 1 + dp(s1, i+1, s2, j+1)
// 	} else {
// 		memo[i][j] = max(dp(s1, i, s2, j+1), dp(s1, i+1, s2, j))
// 	}

// 	return memo[i][j]
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m][n]
}
