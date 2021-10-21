package main

import "fmt"

func plusOne(digits []int) []int {
	plus := 1
	retReverse := make([]int, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + plus
		if tmp >= 10 {
			tmp = 0
			plus = 1
		} else {
			plus = 0
		}

		retReverse[len(digits) -1 - i] = tmp
	}

	length := len(digits)
	if plus == 1 {
		length++
	}
	ret := make([]int, length)
	if plus == 1 {
		ret[0] = 1
	}
	for i, v := range retReverse {
		ret[length - 1 - i] = v
	}
	return ret
}

func main() {
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{9,9,9}))
}
