package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	first, second := 1, 2
	for i := 2; i < n; i++ {
		first, second = second, first + second
	}

	return second
}

func main() {
	fmt.Println(climbStairs(3))
}