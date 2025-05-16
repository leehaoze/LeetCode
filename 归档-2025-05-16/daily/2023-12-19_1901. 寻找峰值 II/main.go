package main

import "fmt"

func findPeakGrid(mat [][]int) []int {
	getVal := func(row, column int) int {
		if row < 0 || row >= len(mat) {
			return -1
		}

		if column < 0 || column >= len(mat[row]) {
			return -1
		}

		return mat[row][column]
	}

	// 找到每一行的顶点
	for rowNum, line := range mat {
		// 找到这一行的峰值，然后看他的上下
		topIdx := findTop(line)
		if topIdx == -1 {
			// 这一行没有山峰
			// 其实理论上讲 不会没有山峰的情况
			continue
		}

		fmt.Println(mat[rowNum][topIdx])

		// 山峰的坐标是 (rowNum, topIdx)
		// 需要看上一行和下一行的值
		if getVal(rowNum-1, topIdx) < mat[rowNum][topIdx] && mat[rowNum][topIdx] > getVal(rowNum+1, topIdx) {
			return []int{rowNum, topIdx}
		}
	}

	return nil
}

// 找到数组中山峰元素的位置
func findTop(line []int) int {
	getVal := func(arr []int, idx int) int {
		if idx < 0 || idx >= len(arr) {
			return -1
		}

		return arr[idx]
	}

	// 左闭右开区间
	left, right := 0, len(line)

	// 搜索区间结束条件 = 区间长度为0 = [i,j) i == j
	for left < right {
		mid := left + (right-left)>>1
		// mid正好是山峰
		if getVal(line, mid-1) < line[mid] && line[mid] > getVal(line, mid+1) {
			return mid
		} else if getVal(line, mid-1) < line[mid] && line[mid] < getVal(line, mid+1) {
			// 上升区间
			left = mid + 1
		} else {
			// 下降区间 与 山谷的情况
			right = mid
		}
	}

	return -1
}

func main() {
	// fmt.Println(findPeakGrid([][]int{{1, 2, 3, 4, 5, 6, 7, 8}, {2, 3, 4, 5, 6, 7, 8, 9}, {3, 4, 5, 6, 7, 8, 9, 10}, {4, 5, 6, 7, 8, 9, 10, 11}}))
	fmt.Println(findPeakGrid([][]int{{7, 2, 3, 1, 2}, {6, 5, 4, 2, 1}}))
}
