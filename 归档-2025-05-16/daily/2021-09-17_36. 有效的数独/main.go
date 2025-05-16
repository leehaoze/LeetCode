package main

import (
	"fmt"
	"strconv"
)

/*
	行列数字只能出现一次
*/

func isValidSudoku(board [][]byte) bool {
	boardRow := make([]map[byte]struct{}, 9)
	boardColumn := make([]map[byte]struct{}, 9)
	boardTable := make([]map[byte]struct{}, 9)

	for i := 0; i < 9; i++ {
		boardRow[i] = make(map[byte]struct{})
		boardColumn[i] = make(map[byte]struct{})
		boardTable[i] = make(map[byte]struct{})
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}
			v := board[i][j] - '1'
			if _, exists := boardRow[i][v]; exists {
				// fmt.Println("row", i, j, v)
				return false
			}

			if _, exists := boardColumn[j][v]; exists {
				// fmt.Println("column", i, j, v)
				return false
			}

			idx := i/3 * 3 + j/3
			if _, exists := boardTable[idx][v]; exists {
				// fmt.Println("table", idx, i, j, v)
				return false
			}

			boardRow[i][v] = struct{}{}
			boardColumn[j][v] = struct{}{}
			boardTable[idx][v] = struct{}{}
		}
	}

	return true
}

func main() {
	input := [][]string{
		{"5", "3", ".", ".", "7", ".", ".", ".", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", "6", "."},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."},
		{".", ".", ".", "4", "1", "9", ".", ".", "5"},
		{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	}
	//
	// input = [][]string{
	// 	{"8", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }
	//
	// input = [][]string{
	// 	{"5", "3", ".", ".", "7", ".", ".", ".", "."},
	// 	{"6", ".", ".", "1", "9", "5", ".", ".", "."},
	// 	{".", "9", "8", ".", ".", ".", ".", "6", "."},
	// 	{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
	// 	{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
	// 	{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
	// 	{".", "6", ".", ".", ".", ".", "2", "8", "."},
	// 	{".", ".", ".", "4", "1", "9", ".", ".", "5"},
	// 	{".", ".", ".", ".", "8", ".", ".", "7", "9"},
	// }

	convInput := make([][]byte, 9)
	for i := range input {
		convInput[i] = make([]byte, 9)
		for j := range input[i] {
			num, err := strconv.ParseInt(input[i][j], 10, 64)
			if err == nil && num != 0 {
				convInput[i][j] = byte(num + '1')
			} else {
				convInput[i][j] = byte('.')
			}
		}
	}

	fmt.Println(isValidSudoku(convInput))
}
