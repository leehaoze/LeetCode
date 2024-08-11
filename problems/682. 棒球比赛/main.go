package main

import (
	"fmt"
	"strconv"
)

func calPoints(operations []string) int {
	scores := make([]int, 0)
	validLength := 0
	ret := 0

	for _, op := range operations {
		fmt.Println(scores, validLength)
		switch op {
		case "C":
			// 前一次得分无效
			ret -= scores[validLength-1]
			validLength--
			continue
		case "D":
			// 前一次得分两倍
			if validLength >= len(scores) {
				scores = append(scores, scores[validLength-1]*2)
			} else {
				scores[validLength] = scores[validLength-1] * 2
			}
		case "+":
			// 前两次的和
			if validLength >= len(scores) {
				scores = append(scores, scores[validLength-2]+scores[validLength-1])
			} else {
				scores[validLength] = scores[validLength-2] + scores[validLength-1]
			}
		default:
			// 转为整数
			val, _ := strconv.ParseInt(op, 10, 64)

			if validLength >= len(scores) {
				scores = append(scores, int(val))
			} else {
				scores[validLength] = int(val)
			}
		}

		ret += scores[validLength]
		validLength++
	}

	// fmt.Println(scores)
	// fmt.Println(ret)
	return ret
}

func main() {
	calPoints([]string{"5", "2", "C", "D", "+"})
}
