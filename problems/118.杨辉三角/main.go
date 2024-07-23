package main

import "fmt"

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。

示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]


提示:

1 <= numRows <= 30
*/

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)
	for i := range ret {
		ret[i] = make([]int, i+1)
	}

	ret[0][0] = 1

	for i := 1; i < numRows; i++ {
		for j := 0; j < i+1; j++ {
			row := i - 1
			leftColumn := j - 1
			rightColumn := j

			var leftVal, rightVal int
			if row >= 0 && leftColumn >= 0 {
				leftVal = ret[row][leftColumn]
			}

			if row >= 0 && rightColumn < row+1 {
				rightVal = ret[row][rightColumn]
			}

			ret[i][j] = leftVal + rightVal
		}
	}

	return ret
}

func main() {
	fmt.Println(generate(5))
}
