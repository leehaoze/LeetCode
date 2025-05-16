package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	var pre = 0
	var cur = 1
	var result = 0
	for i := 1; i < n; i++ {
		result = pre + cur
		pre = cur
		cur = result
	}
	return result
}

func main() {
	fmt.Println(fib(4))
}
