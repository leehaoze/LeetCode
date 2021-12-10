package main

import (
	"fmt"
	"strconv"
	"strings"
)

func findComplement(num int) int {
	b := fmt.Sprintf("%b", num)
	ret := strings.Builder{}
	for _, c := range b {
		if c == '0' {
			ret.WriteRune('1')
		} else {
			ret.WriteRune('0')
		}
	}
	s, _ := strconv.ParseInt(ret.String(), 2, 64)
	// fmt.Println(ret.String())
	return int(s)
}

func main() {
	// fmt.Println(fmt.Sprintf("%b", 1))
	// fmt.Println(fmt.Sprintf("%b, %v", findComplement(1), findComplement(1)))
	fmt.Println(fmt.Sprintf("%b", 20161211))
	fmt.Println(fmt.Sprintf("%b %v", findComplement(20161211), findComplement(20161211)))
}
