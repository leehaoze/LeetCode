package main

import (
	"fmt"
	"strconv"
)

func getSum(a int, b int) int {
	binaryA := strconv.FormatInt(int64(a), 2)
	binaryB :=strconv.FormatInt(int64(b), 2)

	fmt.Println(binaryA)
	fmt.Println(binaryB)
	return 0
}


func main() {
	getSum(1000, -1000)
}
