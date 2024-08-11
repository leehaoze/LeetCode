package main

import "fmt"

func strStr(haystack string, needle string) int {
	i, j := 0, 0
	ret := 0

	for i < len(haystack) && j < len(needle) {
		// fmt.Println(fmt.Sprintf("i: %v, j: %v, haystack[i]: %v, neelde[j]: %v", i, j, string(haystack[i]), string(needle[j])))
		if haystack[i] != needle[j] {
			j = 0
			i = ret + 1
			ret = i
			continue
		}

		i++
		j++
		if j >= len(needle) {
			return ret
		}
	}

	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}
