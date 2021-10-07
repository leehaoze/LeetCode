package main

import "fmt"

func countSegments(s string) int {
	ret := 0
	word := false
	for idx, c := range s {
		if !word && c != ' ' {
			word = true
		}

		if word && c == ' ' {
			ret += 1
			word = false
		}

		if word && idx == len(s)-1 && c != ' ' {
			ret += 1
		}
	}

	return ret
}

func main() {
	fmt.Println(countSegments("Of all the gin joints in all the towns in all the world,   "))
}
