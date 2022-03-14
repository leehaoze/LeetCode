package main

import "fmt"

func minFallingPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(dp); i++ {
		dp[0][i] = matrix[0][i]
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(dp); j++ {
			dp[i][j] = dp[i-1][j] + matrix[i][j]
			if j != 0 && (dp[i-1][j-1]+matrix[i][j]) < dp[i][j] {
				dp[i][j] = dp[i-1][j-1] + matrix[i][j]
			}

			if j != len(dp)-1 && (dp[i-1][j+1]+matrix[i][j]) < dp[i][j] {
				dp[i][j] = dp[i-1][j+1] + matrix[i][j]
			}
		}
	}

	ret := dp[len(matrix)-1][0]
	for i := 1; i < len(matrix[0]); i++ {
		if dp[len(matrix)-1][i] < ret {
			ret = dp[len(matrix)-1][i]
		}
	}
	return ret
}

func main() {
	fmt.Println(minFallingPathSum([][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}))
}
