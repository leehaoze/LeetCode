package main

import "fmt"

func clumsy(N int) int {
	res := 0
	res, N = helper(N)
	res = res + N - 1
	for N >= 3 {
		var subRes int
		subRes, N = helper(N - 2)
		res = res - subRes
		res = res + N - 1
	}
	return res
}

func helper(N int) (int, int) {
	res := N
	no := 0
	for N-1 > 0 && no < 2 {
		res = myOp(res, N-1, no)
		N--
		no++
	}
	return res, N
}

func myOp(a, b, no int) int {
	idx := no % 4
	switch idx {
	case 0:
		return a * b
	case 1:
		return a / b
	case 2:
		return a + b
	default:
		return a - b
	}
}

func main() {
	fmt.Println(clumsy(10))
	fmt.Println(clumsy(4))
	fmt.Println(clumsy(5))
}
