package main

import (
	"fmt"
)

/*

1,  2,  3,  4
5,  6,  7,  8
9,  10, 11, 12
13, 14, 15, 16

     ||
     v


13, 9,  5, 1
14, 10, 6, 2
15, 11, 7, 3
16, 12, 8, 4

[0, 0] => [0, n-1]
[0, 1] => [1, 3]


1,  2,  3,  4,  20
5,  6,  7,  8,  30
9,  10, 11, 12, 40
13, 14, 15, 16, 50
60, 70, 80, 90, 99

*/

func rotate(matrix [][]int) {
	n := len(matrix)

	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < len(matrix[i])-1-i; j++ {
			curX, curY := i, j
			for {
				newPosX, newPosY := n-1-curY, curX
				if newPosX == i && newPosY == j {
					break
				}
				matrix[newPosX][newPosY], matrix[curX][curY] = matrix[curX][curY], matrix[newPosX][newPosY]

				curX, curY = newPosX, newPosY
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(fmt.Sprintf("%d, ", matrix[i][j]))
		}
		fmt.Println("")
	}
}

func main() {
	generateMatrix := func(n int) [][]int {
		matrix := make([][]int, n)
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				matrix[i] = append(matrix[i], i*n+j+1)
			}
		}
		return matrix
	}

	matrix := generateMatrix(4)

	printMatrix(matrix)

	rotate(matrix)

	printMatrix(matrix)
}
