package main

import "fmt"

/*
dp[i]
*/

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]

	for i := 1; i < len(triangle); i++ {
		for j := 0; j <= i; j++ {
			if j-1 < 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
			} else if j >= i {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j]+triangle[i][j], dp[i-1][j-1]+triangle[i][j])
			}
		}
		//fmt.Println(dp)
	}

	var ret = 11111
	//fmt.Println(dp)
	for i := 0; i < len(triangle); i++ {
		ret = min(ret, dp[len(triangle)-1][i])
	}
	return ret
}

func main() {
	fmt.Println(minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}))
}
