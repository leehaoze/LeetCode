package main

import "fmt"

func secondGreaterElement(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	for i := 0; i < length; i++ {
		res[i] = -1
	}

	st1, st2 := []int{}, []int{}

	for idx, val := range nums {
		// 单调栈2有值，且栈顶小于当前v
		for len(st2) > 0 && nums[st2[len(st2)-1]] < val {
			res[st2[len(st2)-1]] = val
			// 弹出栈顶
			st2 = st2[:len(st2)-1]
		}

		// 操作单调栈1
		pos := len(st1) - 1
		for pos >= 0 && nums[st1[pos]] < val {
			pos--
		}

		st2 = append(st2, st1[pos+1:]...)
		st1 = append(st1[:pos+1], idx)
	}

	return res
}

func main() {
	ret := secondGreaterElement([]int{2, 4, 0, 9, 6})
	fmt.Println(ret)
}
