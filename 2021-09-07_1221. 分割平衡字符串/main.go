package main

import "fmt"

func balancedStringSplit(s string) int {
	ret := 0
	top := 0
	stack := make([]uint8, len(s))
	for i := 0; i < len(s); i++ {
		if top == 0 {
			stack[top] = s[i]
			top++
		} else {
			if stack[top-1] != s[i] {
				top--
			} else {
				stack[top] = s[i]
				top++
			}

			if top == 0 {
				ret++
			}
		}
	}

	return ret
}

func main() {
	fmt.Println(balancedStringSplit("RLRRLLRLRL"))
	fmt.Println(balancedStringSplit("RLLLLRRRLR"))
	fmt.Println(balancedStringSplit("LLLLRRRR"))
	fmt.Println(balancedStringSplit("RLRRRLLRLL"))
}
