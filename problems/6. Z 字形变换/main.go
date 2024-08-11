package main

import "fmt"

// func convert(s string, numRows int) string {
// 	if numRows == 1 {
// 		return s
// 	}

// 	// 操作数组
// 	opRow := [2]int{1, -1}
// 	opColumn := [2]int{0, 1}

// 	// 计算Z型数组的列长
// 	columns := len(s) / (numRows + numRows - 2) * (numRows - 1)
// 	if len(s)%(numRows+numRows-2) > 0 {
// 		left := len(s) % (numRows + numRows - 2)

// 		if left <= numRows {
// 			columns += 1
// 		} else {
// 			left -= numRows
// 			columns += left + 1
// 		}
// 	}

// 	// fmt.Println(columns)

// 	// 申请Z形数组
// 	zArray := make([][]rune, numRows)
// 	for idx := range zArray {
// 		zArray[idx] = make([]rune, columns)
// 	}

// 	// 填充数组，遇到边界 操作移动一位
// 	opIdx := 0
// 	zRow := 0
// 	zColumn := 0

// 	for _, c := range s {
// 		zArray[zRow][zColumn] = c

// 		if (zRow == numRows-1 || zRow == 0) && !(zRow == 0 && zColumn == 0) {
// 			opIdx = (opIdx + 1) % 2
// 		}

// 		// fmt.Println(fmt.Sprintf("zRow: %v, zColumn: %v, opIdx: %v", zRow, zColumn, opIdx))
// 		zRow = zRow + opRow[opIdx]
// 		zColumn = zColumn + opColumn[opIdx]

// 	}

// 	// 遍历Z形数组
// 	var ret string
// 	for _, val := range zArray {
// 		for _, c := range val {
// 			if c == ' ' || c == 0 {
// 				continue
// 			}

// 			ret += string(c)
// 		}
// 	}

// 	return ret
// }

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	zArray := make([][]rune, numRows)
	for i := range zArray {
		zArray[i] = make([]rune, 0)
	}

	zRow := 0
	flag := -1

	for _, c := range s {
		zArray[zRow] = append(zArray[zRow], c)
		if zRow == 0 || zRow == numRows-1 {
			flag = -flag
		}
		zRow += flag
	}

	var ret string
	for _, val := range zArray {
		for _, c := range val {
			if c != 0 {
				ret += string(c)
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 9))
	/*
		P
		A
		Y
		P         G
		A       N
		L     I
		I   R
		S I
		H
	*/
}
