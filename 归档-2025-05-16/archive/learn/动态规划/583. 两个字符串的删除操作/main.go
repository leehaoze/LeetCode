package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lst(str1, str2 string) int {
	l1, l2 := len(str1), len(str2)
	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[l1][l2]
}

func minDistance(word1 string, word2 string) int {
	l := lst(word1, word2)
	return len(word1) - l + len(word2) - l
}
