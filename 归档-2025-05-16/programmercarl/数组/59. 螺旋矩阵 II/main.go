package main

import (
	"fmt"
)

var stepColumn = [4]int{1, 0, -1, 0}
var stepRow = [4]int{0, 1, 0, -1}

// generateMatrix 生成一个螺旋矩阵
func generateMatrix(n int) [][]int {
	// 碰到边界就换方向

	// 创建数组
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	// 填充
	var currentRow, currentColumn int
	var currentStep int
	for i := 1; i <= n*n; i++ {
		// 填充当前值
		matrix[currentRow][currentColumn] = i

		// 判断是否需要改变方向
		if changeDirection(matrix, currentRow, currentColumn, currentStep, n) {
			currentStep = (currentStep + 1) % 4
		}

		currentRow += stepRow[currentStep]
		currentColumn += stepColumn[currentStep]
	}

	return matrix
}

// changeDirection 判断是否需要更换排列的方向
func changeDirection(matrix [][]int, currentRow, currentColumn, currentStep, n int) bool {
	newRow, newColumn := currentRow+stepRow[currentStep], currentColumn+stepColumn[currentStep]

	// 越界需要换方向
	if newRow >= n || newColumn >= n || newRow < 0 || newColumn < 0 {
		return true
	}

	// 已经填充了值，需要换方向
	if matrix[newRow][newColumn] != 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(generateMatrix(3))
}
