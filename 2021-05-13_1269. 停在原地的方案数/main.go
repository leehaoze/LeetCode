package main

import "fmt"

func numWays(steps int, arrLen int) int {
	minPos := steps
	if minPos > arrLen {
		minPos = arrLen
	}

	dp := make([][]int, steps+1)
	for i := range dp {
		dp[i] = make([]int, minPos)
	}

	dp[0][0] = 1
	for i := 1; i <= steps; i++ {
		for j := 0; j < minPos; j++ {
			temp := 0
			if i > 0 {
				temp = (dp[i-1][j] + temp) % 1000000007
			}

			if j > 0 {
				temp = (dp[i-1][j-1] + temp) % 1000000007
			}

			if j < minPos-1 {
				temp = (dp[i-1][j+1] + temp) % 1000000007
			}

			dp[i][j] = temp
		}
	}

	//fmt.Println(dp)
	return dp[steps][0]
}

func main() {
	fmt.Println(numWays(3, 2))
	fmt.Println(numWays(4, 2))
}
