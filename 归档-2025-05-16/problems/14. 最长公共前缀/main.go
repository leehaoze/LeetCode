package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	var ret string
	var idx int	
	for {
		var curChar byte
		for _, str := range strs {
			if idx >= len(str) {
				return ret
			}

			if curChar == 0 {
				curChar = str[idx]
			} else if curChar != str[idx] {
				return ret
			}
		}

		ret += string(curChar)
		idx++
	}
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}