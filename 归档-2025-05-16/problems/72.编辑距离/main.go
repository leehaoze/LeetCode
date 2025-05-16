package main

import "fmt"

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成
*/

/*
A 经过1步变为 B有三种情况

三种操作对应三种情况

dp[i][j] = dp[i-1][j] + 1
		   dp[i][j-1] + 1
           dp[i-1][j-1] + 1

*/

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for idx := range dp {
		dp[idx] = make([]int, len(word2)+1)
	}

	// 初始化
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}

	for i := 1; i <= len(word2); i++ {
		dp[0][i] = i
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			// 替换的情况
			dist1 := dp[i-1][j-1] + 1
			if word1[i-1] == word2[j-1] {
				dist1 = dp[i-1][j-1]
			}

			// 插入的情况
			dist2 := dp[i-1][j] + 1
			dist3 := dp[i][j-1] + 1

			dp[i][j] = min(min(dist1, dist2), dist3)
		}
	}

	return dp[len(word1)][len(word2)]
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
