package main

import (
	"fmt"
)

//type Stack struct {
//	list *list.List
//}
//
//func NewStack() *Stack {
//	list := list.New()
//	return &Stack{list}
//}
//
//func (stack *Stack) Push(value interface{}) {
//	stack.list.PushBack(value)
//}
//
//func (stack *Stack) Pop() interface{} {
//	e := stack.list.Back()
//	if e != nil {
//		stack.list.Remove(e)
//		return e.Value
//	}
//	return nil
//}
//
//func (stack *Stack) Len() int {
//	return stack.list.Len()
//}
//
//func (stack *Stack) NotEmpty() bool {
//	return stack.list.Len() != 0
//}
//
//func combineSet(a, b map[int]bool) map[int]bool {
//	set := make(map[int]bool)
//	for k, v := range a {
//		set[k] = v
//	}
//
//	for k, v := range b {
//		set[k] = v
//	}
//
//	return set
//}
//
//func isBipartite(graph [][]int) bool {
//	setA := make(map[int]bool)
//	setB := make(map[int]bool)
//
//	stackA := NewStack()
//	stackB := NewStack()
//
//	for idx, item := range graph {
//		if len(item) == 0 {
//			continue
//		}
//		stackA.Push(idx)
//		setA[idx] = true
//
//	}
//
//	for idx := 0; stackA.NotEmpty() || stackB.NotEmpty(); idx++ {
//		var stack *Stack
//		if idx%2 == 0 {
//			stack = stackA
//		} else {
//			stack = stackB
//		}
//
//		for stack.NotEmpty() {
//			node := stack.Pop().(int)
//			for _, item := range graph[node] {
//				if idx%2 == 0 {
//					if setA[item] {
//						return false
//					}
//
//					_, exists := setB[item]
//					if !exists {
//						stackB.Push(item)
//					}
//					setB[item] = true
//				} else {
//					if setB[item] {
//						return false
//					}
//
//					_, exists := setA[item]
//					if !exists {
//						stackA.Push(item)
//					}
//					setA[item] = true
//
//				}
//			}
//		}
//	}
//	return true
//}

var (
	UNCOLORED, RED, GREEN = 0, 1, 2
	color                 []int
	valid                 bool
)

func isBipartite(graph [][]int) bool {
	valid = true
	length := len(graph)

	color = make([]int, length)

	for i := 0; i < length; i++ {
		if color[i] == UNCOLORED {
			dfs(i, RED, graph)
		}
	}
	return valid
}

func dfs(node, c int, graph [][]int) {
	color[node] = c
	cNew := RED
	if c == RED {
		cNew = GREEN
	}

	for _, neighbor := range graph[node] {
		if color[neighbor] == UNCOLORED {
			dfs(neighbor, cNew, graph)
			if !valid {
				return
			}
		} else if color[neighbor] != cNew {
			valid = false
			return
		}
	}
}

func main() {
	//fmt.Printf("result is %t\n", isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Printf("result is %t\n", isBipartite([][]int{{1}, {0, 3}, {3}, {1, 2}}))
}
