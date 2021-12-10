package main

import (
	"fmt"
	"strings"
)

func countAndSay(n int) string {
	ret := "1"
	for i := 1; i < n; i++ {
		ret = trans(ret)
	}

	return ret
}

func trans(number string) string {
	descStr := strings.Builder{}
	preC := '0'
	count := 0
	for i, c := range number {
		if i == 0 {
			preC = c
			count += 1
			continue
		}

		if c == preC {
			count++
		} else {
			descStr.WriteRune(rune('0' + count))
			descStr.WriteRune(preC)

			preC = c
			count = 1
		}
	}

	descStr.WriteRune(rune('0' + count))
	descStr.WriteRune(preC)

	return descStr.String()
}


func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(countAndSay(i))
	}
}
