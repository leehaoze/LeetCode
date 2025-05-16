package main

import "fmt"

/*
给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。

dp[n] = n个节点组成的互不相同的二叉搜索的数量

dp[i] = sum(max(dp[j] + dp[i-j-1], dp[j] * dp[i-1-j]), j=[0,i) )

dp[1] = 1
dp[2] = 2
dp[3] =
*/

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-1-j]
		}

	}

	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
