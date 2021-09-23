package main

import "fmt"

/**
128 64 32 16 8 4 2 1
3 0011
9 1001
27 0001 1011
71 0100 0111
213 1101 0101
*/

func isPowerOfThree(n int) bool {
	// fmt.Println("n is ", n)
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}

	// fmt.Println(n%3)
	if n == 0 || n%3 != 0 {
		return false
	}

	return isPowerOfThree(n / 3)
}

func main() {
	fmt.Println(isPowerOfThree(27))
	fmt.Println(isPowerOfThree(28))
	fmt.Println(isPowerOfThree(3))
	fmt.Println(isPowerOfThree(-3))
	fmt.Println(isPowerOfThree(-9))
	fmt.Println(isPowerOfThree(1))

}
