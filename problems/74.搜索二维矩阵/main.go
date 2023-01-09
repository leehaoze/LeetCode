package _4_搜索二维矩阵

// func searchMatrix(matrix [][]int, target int) bool {
// 	left, right := 0, len(matrix)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if matrix[mid][0] == target {
// 			return true
// 		} else if matrix[mid][0] > target {
// 			right = mid - 1
// 		} else {
// 			left = mid + 1
// 		}
// 	}
//
// 	idx := left - 1
// 	if idx < 0 {
// 		return false
// 	}
//
// 	left, right = 0, len(matrix[idx])-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if matrix[idx][mid] == target {
// 			return true
// 		} else if matrix[idx][mid] > target {
// 			right = mid - 1
// 		} else if matrix[idx][mid] < target {
// 			left = mid + 1
// 		}
// 	}
//
// 	return false
// }

func searchMatrix(matrix [][]int, target int) bool {
	row := len(matrix)
	column := len(matrix[0])
	left, right := 0, row*column-1
	for left <= right {
		mid := left + (right-left)/2
		i, j := mid/column, mid%column
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
