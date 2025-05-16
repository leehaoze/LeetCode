package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxScore(grid [][]int) int {
	// 初始化
	dp := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}

	// 求解
	ret := -100000 - 1
	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[i]) - 1; j >= 0; j-- {
			if i == len(grid)-1 && j == len(grid[i])-1 {
				dp[i][j] = 0
				continue
			}

			// 横向求解 不在最后一列就可以横向求解
			if j < len(grid[i])-1 {
				for walkColumn := j + 1; walkColumn < len(grid[i]); walkColumn++ {
					// 必须走一步
					if walkColumn == j+1 {
						dp[i][j] = max(grid[i][walkColumn]-grid[i][j], grid[i][walkColumn]-grid[i][j]+dp[i][walkColumn])
					} else {
						dp[i][j] = max(dp[i][j], max(grid[i][walkColumn]-grid[i][j], grid[i][walkColumn]-grid[i][j]+dp[i][walkColumn]))
					}
					ret = max(dp[i][j], ret)
				}
			}

			// 纵向求解，不在最后一行就可以纵向求解
			if i < len(grid)-1 {
				for walkRow := i + 1; walkRow < len(grid); walkRow++ {
					// 横向没走过，则纵向必须走一步
					if j == len(grid[i])-1 && walkRow == i+1 {
						dp[i][j] = max(grid[walkRow][j]-grid[i][j], grid[walkRow][j]-grid[i][j]+dp[walkRow][j])
					} else {
						dp[i][j] = max(dp[i][j], max(grid[walkRow][j]-grid[i][j], grid[walkRow][j]-grid[i][j]+dp[walkRow][j]))
					}

					ret = max(dp[i][j], ret)
				}
			}
		}
	}

	// for _, line := range dp {
	// 	fmt.Println(line)
	// }

	return ret
}

func main() {
	fmt.Println(maxScore([][]int{{4, 3, 2}, {3, 2, 1}}))
}
