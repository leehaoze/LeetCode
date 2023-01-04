package main

type NumMatrix struct {
	preSum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	rows := len(matrix)
	columns := len(matrix[0])

	preSum := make([][]int, rows)
	for i := 0; i < rows; i++ {
		preSum[i] = make([]int, columns)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if i == 0 && j == 0 {
				preSum[i][j] = matrix[i][j]
			} else if i == 0 {
				// 第1行
				preSum[i][j] = preSum[i][j-1] + matrix[i][j]
			} else if j == 0 {
				preSum[i][j] = preSum[i-1][j] + matrix[i][j]
			} else {
				preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i][j]
			}
		}
	}

	return NumMatrix{preSum: preSum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ret := this.preSum[row2][col2]
	if row1 != 0 {
		ret -= this.preSum[row1-1][col2]
	}

	if col1 != 0 {
		ret -= this.preSum[row2][col1-1]
	}

	if row1 != 0 && col1 != 0 {
		ret += this.preSum[row1-1][col1-1]
	}

	return ret
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
