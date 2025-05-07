package main

// minDistance 返回两个字符串的最短编辑距离
// 编辑可以是以下三种操作中的某一个：
// - 在word1中添加一个字母
// - 在word2中添加一个字母
// - 修改word1中的一个字母
func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i <= len(word1); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	// 先初始化边界情况
	for i := 0; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for j := 0; j <= len(word2); j++ {
		dp[0][j] = j
	}

	// 计算
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			} else {
				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
			}
		}
	}

	return dp[len(word1)][len(word2)]
}
