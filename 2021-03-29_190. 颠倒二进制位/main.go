package main

import (
	"fmt"
	"strconv"
)

func reverseBits(num uint32) uint32 {
	stack := fmt.Sprintf("%032b", num)

	fmt.Println(stack)

	stack = reverse(stack)

	res, _ := strconv.ParseUint(stack, 2, 64)
	return uint32(res)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverseBits(43261596))
}
