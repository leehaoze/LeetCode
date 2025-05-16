package main

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	dp := make([][]int, len(grid))
	for i := 0; i < len(dp) ; i++ {
		dp[i] = make([]int, len(grid[0]))
	}

	for rowNum := 0; rowNum < len(grid); rowNum++ {
		for columnNum := 0; columnNum < len(grid[rowNum]); columnNum++ {
			if rowNum == 0 && columnNum == 0 {
				dp[rowNum][columnNum] = grid[rowNum][columnNum]
			} else if rowNum == 0 {
				dp[rowNum][columnNum] = dp[rowNum][columnNum-1] + grid[rowNum][columnNum]
			} else if columnNum == 0 {
				dp[rowNum][columnNum] = dp[rowNum-1][columnNum] + grid[rowNum][columnNum]
			} else {
				dp[rowNum][columnNum] = Min(dp[rowNum-1][columnNum], dp[rowNum][columnNum-1]) + grid[rowNum][columnNum]
			}
		}
	}

	return dp[len(grid)-1][len(grid[0])-1]
}

func main() {
	ret := minPathSum([][]int{{1, 3, 1}, {1, 4, 1}, {4, 2, 1}})
	print(ret)
}
