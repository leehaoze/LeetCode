package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	// init dp array
	dp := make([][]int, len(obstacleGrid))
	column := len(obstacleGrid[0])
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, column)
	}

	// init dp first row
	dp[0][0] = 1
	for i := 1; i < column && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	// init dp first column
	for i := 1; i < len(obstacleGrid) && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	// cal
	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < column; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[len(obstacleGrid)-1][column-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}))
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 1}, {0, 0}}))
}
