package main

import (
	"fmt"
	"strings"
)

func findRepeatedDnaSequences(s string) []string {
	window := strings.Builder{}
	start, end := 0, 0
	m := make(map[string]int)

	for end < len(s) {
		window.WriteByte(s[end])
		if end-start == 10 {
			start++
			temp := window.String()[1:]
			window.Reset()
			window.WriteString(temp)
		}

		m[window.String()]++
		end++
	}

	ret := make([]string, 0)
	for s, v := range m {
		if v > 1 {
			ret = append(ret, s)
		}
	}

	return ret
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAAA"))
}
