package main

import "fmt"

func execQuery(leftBound, rightBound, preSum []int, query []int) int {
	left, right := query[0], query[1]
	// leftBound, rightBound := 0, 0
	// for i := left; i <= right; i++ {
	// 	if s[i] == '|' {
	// 		leftBound = i
	// 		break
	// 	}
	// }
	//
	// for i := right; i >= left; i-- {
	// 	if s[i] == '|' {
	// 		rightBound = i
	// 		break
	// 	}
	// }
	l, r := -1, -1
	if rightBound[left] != -1 && rightBound[left] <= right {
		l = rightBound[left]
	}

	if leftBound[right] != -1 && leftBound[right] >= left {
		r = leftBound[right]
	}

	if l == r || l == -1 || r == -1 {
		return 0
	}

	// fmt.Println(fmt.Sprintf("left:%v, right:%v", l, r))
	return preSum[r] - preSum[l]
}

func platesBetweenCandles(s string, queries [][]int) []int {
	preSum := make([]int, len(s))
	preBound := -1
	for pos, c := range s {
		if c == '|' && preBound == -1 {
			preBound = pos
		} else if c == '|' && preBound != -1 {
			preSum[pos] = preSum[preBound] + pos - preBound - 1
			preBound = pos
		}
	}

	rightBound := make([]int, len(s))
	preBound = -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '|' {
			preBound = i
		}

		rightBound[i] = preBound
	}

	leftBound := make([]int, len(s))
	preBound = -1
	for i := 0; i < len(s); i++ {
		if s[i] == '|' {
			preBound = i
		}

		leftBound[i] = preBound
	}
	// for _, i := range preSum {
	// 	fmt.Print(i)
	// }
	// fmt.Println("\n")
	// for _, i := range leftBound {
	// 	fmt.Print(i)
	// }
	// fmt.Println("\n")
	// for _, i := range rightBound {
	// 	fmt.Print(i)
	// }
	// fmt.Println("\n")

	ret := make([]int, len(queries))
	for idx, query := range queries {
		ret[idx] = execQuery(leftBound, rightBound, preSum, query)
	}

	return ret
}

func main() {
	fmt.Println(platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}}))
	fmt.Println(platesBetweenCandles("***|**|*****|**||**|*", [][]int{{1, 17}, {4, 5}, {14, 17}, {5, 11}, {15, 16}}))
}
