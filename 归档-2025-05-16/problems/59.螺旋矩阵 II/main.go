package main

import (
	"fmt"
)

var RowStep = []int{0, 1, 0, -1}
var ColumnStep = []int{1, 0, -1, 0}

func generateMatrix(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}

	row, column := 0, 0
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	val := 1

	curStep := 0

	for {
		ret[row][column] = val
		val++

		nextRow := row + RowStep[curStep]
		nextColumn := column + ColumnStep[curStep]
		if nextRow >= n || nextRow < 0 || nextColumn >= n || nextColumn < 0 || ret[nextRow][nextColumn] != 0 {
			curStep = (curStep + 1) % 4
		}

		row += RowStep[curStep]
		column += ColumnStep[curStep]
		if ret[row][column] != 0 {
			break
		}
	}

	return ret
}

func main() {
	ret := generateMatrix(1)

	for i := 0; i < len(ret); i++ {
		for j := 0; j < len(ret[0]); j++ {
			fmt.Printf("%v ", ret[i][j])
		}
		fmt.Print("\n")
	}

}
