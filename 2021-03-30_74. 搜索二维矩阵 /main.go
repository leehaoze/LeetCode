package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	column := len(matrix[0])

	// binary search
	return binarySearch(matrix, row, column, target)
}

func binarySearch(matrix [][]int, row, column, target int) bool {
	left := 0
	right := row*column - 1

	for left <= right {
		mid := (left + right) / 2
		midValue := matrix[mid/column][mid%column]
		if target == midValue {
			return true
		} else if target < midValue {
			right = mid - 1
		} else if target > midValue {
			left = mid + 1
		}
	}

	return false
}

func main() {
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 1))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 2))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3))
	fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))

}
