package main

import (
	"fmt"
)

type Stack struct {
	data []int
	idx  int
}

func NewStack(length int) *Stack {
	return &Stack{
		data: make([]int, length),
		idx:  -1,
	}
}

func (s *Stack) Push(val int) {
	s.idx++
	s.data[s.idx] = val
}

func (s *Stack) Pop() {
	s.idx--
}

func (s *Stack) Empty() bool {
	return s.idx < 0
}

func (s *Stack) Top() int {
	return s.data[s.idx]
}

func validateStackSequences(pushed []int, popped []int) bool {
	stack := NewStack(len(pushed))
	pushIdx, popIdx := 0, 0
	for pushIdx < len(pushed) || (!stack.Empty() && popIdx < len(popped) && stack.Top() == popped[popIdx]) {
		if stack.Empty() {
			// fmt.Println(fmt.Sprintf("stack empty, push %v", pushed[pushIdx]))
			stack.Push(pushed[pushIdx])
			pushIdx++
		} else {
			if stack.Top() == popped[popIdx] {
				// fmt.Println(fmt.Sprintf("stack equal to popped, pop %v", stack.Top()))
				stack.Pop()
				popIdx++
			} else {
				// fmt.Println(fmt.Sprintf("stack not equal, push %v", pushed[pushIdx]))
				stack.Push(pushed[pushIdx])
				pushIdx++
			}
		}
	}

	return popIdx == len(popped)
}

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}))
}
