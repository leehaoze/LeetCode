package main

import "fmt"

// func canJump(nums []int) bool {
// 	// 深度优先搜索同时配合记忆
// 	nodes := make([]int, len(nums))
// 	nodesCount := 0

// 	nodes = append(nodes, 0)
// 	nodesCount++

// 	memory := make([]int, len(nums))

// 	for nodesCount > 0 {
// 		fmt.Println(fmt.Sprintf("当前待遍历数组：%#v， nodesCount: %v", nodes, nodesCount))

// 		// 取出最后压栈的节点
// 		idx := nodes[nodesCount-1]
// 		memory[idx] = 1
// 		val := nums[idx]
// 		nodesCount--

// 		fmt.Println(fmt.Sprintf("取出的节点: %v, 对应的数组值: %v", idx, val))

// 		if idx+val >= len(nums)-1 {
// 			return true
// 		}

// 		for i := 1; i <= val && idx+i < len(nums); i++ {
// 			if memory[idx+i] == 1 {
// 				continue
// 			}

// 			if nodesCount > len(nodes) {
// 				nodes = append(nodes, idx+i)
// 			} else {
// 				nodes[nodesCount] = idx + i
// 			}
// 			nodesCount++
// 		}
// 	}

// 	return false
// }

func canJump(nums []int) bool {
	maxDis := nums[0]

	for i := 1; i < len(nums) && i <= maxDis; i++ {
		if i+nums[i] > maxDis {
			maxDis = i + nums[i]
			// fmt.Println(maxDis)
		}
	}

	return maxDis >= len(nums)-1
}

func main() {
	fmt.Println(canJump([]int{0}))
}
