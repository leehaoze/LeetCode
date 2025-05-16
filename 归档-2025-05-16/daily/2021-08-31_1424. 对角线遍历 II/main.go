package main

import "fmt"

/*
	0,0 0,1 0,2 0,3
	1,0 1,1 1,2 1,3
	2,0 2,1 2,2 2,3
	3,0 3,1 3,2 3,3
*/

/*
	20 17 13 14
	12 6
	3  4

	0,0 0,1 0,2, 0,3
	1,0 1,1
	2,0 2,1
*/

/*
	14 12 19 16 9
	13 14 15 8  11
	11 13 1
*/

func findDiagonalOrder(nums [][]int) []int {
	ret := make([][]int, 100000*2)

	for row := range nums {
		for column := range nums[row] {
			if len(ret[row+column]) == 0 {
				ret[row+column] = make([]int, 0)
			}

			ret[row+column] = append(ret[row+column], nums[row][column])
		}
	}

	result := make([]int, 0)
	for i := range ret {
		if len(ret[i]) == 0 {
			continue
		}

		for j := len(ret[i]) - 1; j >= 0; j-- {
			result = append(result, ret[i][j])
		}

	}

	return result
}

func main() {
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5}, {6, 7}, {8}, {9, 10, 11}, {12, 13, 14, 15, 16}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4}, {5, 6, 7}, {8}, {9, 10, 11}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3, 4, 5, 6}}))
	fmt.Println(findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	fmt.Println(findDiagonalOrder([][]int{{20, 17, 13, 14}, {12, 6}, {3, 4}}))
	fmt.Println(findDiagonalOrder([][]int{{14, 12, 19, 16, 9}, {13, 14, 15, 8, 11}, {11, 13, 1}}))
}
