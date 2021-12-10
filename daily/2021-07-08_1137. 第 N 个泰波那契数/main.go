package main

import "fmt"

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	a, b, c := 0, 1, 1
	result := 0
	for i := 2; i < n; i++ {
		result = a + b + c
		a, b, c = b, c, result
	}

	return result
}

func main() {
	fmt.Println(tribonacci(25))
}
