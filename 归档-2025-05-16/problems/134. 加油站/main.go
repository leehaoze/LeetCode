package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	stations := len(gas)

	for startPos := 0; startPos < stations; {
		// 从当前点出发，看是否能走一圈

		// fmt.Println(fmt.Sprintf("当前出发点: %v", startPos))
		leftGas := 0
		curPos := startPos

		for {
			// 在当前加油站加油
			leftGas += gas[curPos]

			// fmt.Println(fmt.Sprintf("     当前所处站点: %v,剩余油量: %v,到达下一站点需要: %v", curPos, leftGas, cost[curPos]))

			// 能否到达下一个加油站
			if cost[curPos] > leftGas {
				// 不可以到达，寻找下一个解
				if curPos+1 <= startPos {
					// 转了一圈了 不行
					return -1
				}
				startPos = (curPos + 1)
				break
			}

			// 能到达下一个加油站
			// 更新油量
			leftGas -= cost[curPos]

			// 更新位置
			curPos = (curPos + 1) % len(gas)

			// 是否走了一圈
			if curPos == startPos {
				return startPos
			}
		}
	}

	return -1
}

func main() {
	// fmt.Println(canCompleteCircuit([]int{5, 1, 2, 3, 4}, []int{4, 4, 1, 5, 1}))
	fmt.Println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
}
